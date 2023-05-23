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

	"github.com/ServiceWeaver/weaver"
)

// Factorer is a prime factorization component.
type Factorer interface {
	// Factor returns the prime factors of the provided integer. If a factor
	// has multiplicity n, it is returned n times. For example, the prime
	// factors of 60 are 2, 2, 3, and 5.
	Factor(context.Context, int) ([]int, error)
}

// factorer is the implementation of the Factorer component.
type factorer struct {
	weaver.Implements[Factorer]
	cache weaver.Ref[Cache]
}

func (f *factorer) Factor(ctx context.Context, x int) ([]int, error) {
	f.Logger().Debug("Factor", "x", x)

	// Sanitize input.
	if x <= 0 {
		return nil, fmt.Errorf("non-positive x: %d", x)
	}
	if x == 1 {
		return []int{1}, nil
	}

	// Try to get the factorization from the cache, but continue if it's not
	// found or there is an error.
	if factors, err := f.cache.Get().Get(ctx, x); err != nil {
		f.Logger().Error("cache.Get", "x", x, "err", err)
	} else if len(factors) > 0 {
		return factors, nil
	}

	// Compute the prime factorization.
	var factors []int
	original := x
	for x >= 2 {
		for factor := 2; factor <= x; factor++ {
			if x%factor == 0 {
				factors = append(factors, factor)
				x = x / factor
				break
			}
		}
	}
	sort.Ints(factors)

	// Try to put the factorization in the cache, but continue if there is an
	// error.
	if err := f.cache.Get().Put(ctx, original, factors); err != nil {
		f.Logger().Error("cache.Put", "x", x, "factors", factors, "err", err)
	}

	// Return the result.
	return factors, nil
}
