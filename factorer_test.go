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

func TestFactor(t *testing.T) {
	type test struct {
		x    int
		want []int
	}

	for _, test := range []test{
		{1, []int{1}},
		{2, []int{2}},
		{3, []int{3}},
		{12347, []int{12347}},
		{2 * 3, []int{2, 3}},
		{2 * 2 * 3, []int{2, 2, 3}},
		{2 * 2 * 2 * 3 * 3 * 7, []int{2, 2, 2, 3, 3, 7}},
	} {
		for _, runner := range weavertest.AllRunners() {
			name := fmt.Sprintf("%s/%d", runner.Name(), test.x)
			t.Run(name, func(t *testing.T) {
				runner.Run(t, func(factorer Factorer) {
					got, err := factorer.Factor(context.Background(), test.x)
					if err != nil {
						t.Fatalf("Factor: %v", err)
					}
					if diff := cmp.Diff(test.want, got); diff != "" {
						t.Fatalf("Factor (-want,+got):\n%s", diff)
					}
				})
			})
		}
	}
}

func TestFactorFailsOnNegativeNumber(t *testing.T) {
	for _, runner := range weavertest.AllRunners() {
		t.Run(runner.Name(), func(t *testing.T) {
			runner.Run(t, func(factorer Factorer) {
				_, err := factorer.Factor(context.Background(), -42)
				if err == nil {
					t.Fatal("Unexpected success.")
				}
			})
		})
	}
}

func TestFactorFailsOnZero(t *testing.T) {
	for _, runner := range weavertest.AllRunners() {
		t.Run(runner.Name(), func(t *testing.T) {
			runner.Run(t, func(factorer Factorer) {
				_, err := factorer.Factor(context.Background(), 0)
				if err == nil {
					t.Fatal("Unexpected success.")
				}
			})
		})
	}
}
