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
	"time"
)

// TestInstructions tests that the instructions in README.md work as expected.
// Workshop participants can ignore this file.
func TestInstructions(t *testing.T) {
	// Build the binary.
	if err := exec.Command("go", "build", ".").Run(); err != nil {
		t.Fatalf("go build .: %v", err)
	}

	// Start the server.
	server := exec.Command("weaver", "multi", "deploy", "config.toml")
	if err := server.Start(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := server.Process.Kill(); err != nil {
			t.Fatal(err)
		}
	}()

	// Give the server time to start.
	time.Sleep(time.Second)

	// Inspect the status.
	if err := exec.Command("weaver", "multi", "status").Run(); err != nil {
		t.Fatalf("weaver multi status: %v", err)
	}

	// Curl the server.
	for _, test := range []struct{ url, want string }{
		{"localhost:9000/search?q=sushi", `["🍣"]` + "\n"},
		{"localhost:9000/search?q=curry", `["🍛"]` + "\n"},
		{"localhost:9000/search?q=shrimp", `["🍤","🦐"]` + "\n"},
		{"localhost:9000/search?q=lobster", `["🦞"]` + "\n"},
	} {
		out, err := exec.Command("curl", test.url).Output()
		if err != nil {
			t.Fatalf("curl %s: %v", test.url, err)
		}
		if string(out) != test.want {
			t.Fatalf("curl %s: got %q, want %q", test.url, string(out), test.want)
		}
	}
}
