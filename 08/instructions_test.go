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
	if err := exec.Command("weaver", "generate", ".").Run(); err != nil {
		t.Fatalf("weaver generate .: %v", err)
	}

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

	// Curl the server.
	const url = "localhost:9000/search?q=pig"
	const want = `["🐖","🐗","🐷","🐽"]` + "\n"
	for i := 0; i < 4; i++ {
		out, err := exec.Command("curl", url).Output()
		if err != nil {
			t.Fatalf("curl %s: %v", url, err)
		}
		if string(out) != want {
			t.Fatalf("curl %s: got %q, want %q", url, string(out), want)
		}
	}
}
