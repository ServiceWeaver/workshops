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
	"sync"

	"github.com/ServiceWeaver/weaver"
)

// Cache caches the prime factorizations of integers.
type Cache interface {
	// Get returns the cached prime factorization of the provided integer. On
	// cache miss, Get returns nil, nil.
	Get(context.Context, int) ([]int, error)

	// Put stores a prime factorization in the cache.
	Put(context.Context, int, []int) error
}

// cache implements the Cache component.
type cache struct {
	weaver.Implements[Cache]
	weaver.WithRouter[router]

	mu             sync.Mutex
	factorizations map[int][]int
}

func (c *cache) Init(context.Context) error {
	c.factorizations = map[int][]int{}
	return nil
}

func (c *cache) Get(_ context.Context, x int) ([]int, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger().Debug("Get", "x", x)
	return c.factorizations[x], nil
}

func (c *cache) Put(_ context.Context, x int, factors []int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Logger().Debug("Put", "x", x, "factors", factors)
	c.factorizations[x] = factors
	return nil
}

// router routes requests to the Cache component. Both Get and Put use the
// integer x (i.e. the value whose prime factors are cached) as the routing
// key. Calls to these methods with the same value of x will tend to be routed
// to the same replica.
type router struct{}

func (router) Get(_ context.Context, x int) int {
	return x
}

func (router) Put(_ context.Context, x int, factors []int) int {
	return x
}
