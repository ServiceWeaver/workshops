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

package main_test

import (
	"os/exec"
	"testing"
)

// TestInstructions tests that the instructions in README.md work as expected.
// Workshop participants can ignore this file.
func TestInstructions(t *testing.T) {
	if err := exec.Command("weaver", "generate", ".").Run(); err != nil {
		t.Fatalf("weaver generate .: %v", err)
	}

	out, err := exec.Command("go", "run", ".").Output()
	if err != nil {
		t.Fatalf("go run .: %v", err)
	}
	const want = "[🐖 🐗 🐷 🐽]\n"
	if string(out) != want {
		t.Fatalf("go run .: got %q, want %q", string(out), want)
	}
}
