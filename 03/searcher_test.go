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
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	type test struct {
		query string
		want  []string
	}

	for _, test := range []test{
		{"pig", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"PiG", []string{"🐖", "🐗", "🐷", "🐽"}},
		{"black cat", []string{"🐈\u200d⬛"}},
		{"foo bar baz", nil},
	} {
		for _, runner := range weavertest.AllRunners() {
			runner.Name = fmt.Sprintf("%s/%q", runner.Name, test.query)
			runner.Test(t, func(t *testing.T, searcher Searcher) {
				got, err := searcher.Search(context.Background(), test.query)
				if err != nil {
					t.Fatalf("Search: %v", err)
				}
				if diff := cmp.Diff(test.want, got); diff != "" {
					t.Fatalf("Search (-want,+got):\n%s", diff)
				}
			})
		}
	}
}
