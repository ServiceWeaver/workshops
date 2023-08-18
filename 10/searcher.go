// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver"
	"github.com/ServiceWeaver/weaver/metrics"
	"github.com/rivo/uniseg"
	"golang.org/x/exp/slices"
)

// Two counters that count cache hits and misses in the Search method.
var (
	cacheHits = metrics.NewCounter(
		"search_cache_hits",
		"Number of Search cache hits",
	)
	cacheMisses = metrics.NewCounter(
		"search_cache_misses",
		"Number of Search cache misses",
	)
)

// Searcher is an emoji search engine component.
type Searcher interface {
	// Search returns the set of emojis that match the provided query.
	Search(context.Context, string) ([]string, error)

	// SearchChatGPT returns the set of emojis that ChatGPT thinks match the
	// provided query.
	SearchChatGPT(context.Context, string) ([]string, error)
}

// searcher is the implementation of the Searcher component.
type searcher struct {
	weaver.Implements[Searcher]
	cache   weaver.Ref[Cache]
	chatgpt weaver.Ref[ChatGPT]
}

func (s *searcher) Search(ctx context.Context, query string) ([]string, error) {
	s.Logger(ctx).Debug("Search", "query", query)

	// Try to get the emojis from the cache, but continue if it's not found or
	// there is an error.
	if emojis, err := s.cache.Get().Get(ctx, query); err != nil {
		s.Logger(ctx).Error("cache.Get", "query", query, "err", err)
	} else if emojis != nil {
		cacheHits.Inc()
		return emojis, nil
	} else {
		cacheMisses.Inc()
	}

	// Perform the search. First, we lowercase the query and split it into
	// words. For example, the query "Black cat" is tokenized to the words
	// "black" and "cat". Then, we say an emoji matches a query if every word
	// in the query is one of the emoji's labels.
	//
	// For example, the cat emoji has labels ["animal", "cat"]. It does NOT
	// match the "Black cat" query because "black" is not a label. The black
	// cat emoji, on the other hand, has labels ["animal", "black", "cat"] and
	// therefore does match the query "Black cat".
	//
	// We iterate over all emojis and return the ones that match the query.
	words := strings.Fields(strings.ToLower(query))
	results := []string{}
	for emoji, labels := range emojis {
		if matches(labels, words) {
			results = append(results, emoji)
		}
	}
	sort.Strings(results)

	// Try to cache the results, but continue if there is an error.
	if err := s.cache.Get().Put(ctx, query, results); err != nil {
		s.Logger(ctx).Error("cache.Put", "query", query, "err", err)
	}

	return results, nil
}

// matches returns whether words is a subset of labels.
func matches(labels, words []string) bool {
	for _, word := range words {
		if !slices.Contains(labels, word) {
			return false
		}
	}
	return true
}

func (s *searcher) SearchChatGPT(ctx context.Context, query string) ([]string, error) {
	// Issue a completion request to ChatGPT.
	prompt := fmt.Sprintf("Give me a list of emojis that related to the query %q. Don't give an explanation.", query)
	completion, err := s.chatgpt.Get().Complete(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("chatgpt: %w", err)
	}
	s.Logger(ctx).Debug("ChatGPT completion", "query", query, "completion", completion)

	// Parse the emojis from ChatGPT's response. This is surprisingly tricky
	// since some emoji sequences contain many codepoints, and not every
	// sequence of emoji codepoints is a valid emoji sequence. Every emoji and
	// emoji sequence is a single grapheme cluster [1], so to extract the
	// emojis, we split the comletion into its graphemes and check each one
	// against our database of emojis. We also make sure to remove duplicates,
	// as ChatGPT tends to reply with the same emoji multiple times.
	//
	// [1]: https://unicode.org/reports/tr29/
	var results []string
	seen := map[string]bool{}
	graphemes := uniseg.NewGraphemes(completion)
	for graphemes.Next() {
		emoji := graphemes.Str()
		if _, ok := emojis[emoji]; ok && !seen[emoji] {
			results = append(results, emoji)
		}
		seen[emoji] = true
	}
	return results, nil
}
