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
	for _, url := range []string{
		"localhost:9000/search?q=two",
		"localhost:9000/search?q=two",
		"localhost:9000/search?q=three",
		"localhost:9000/search?q=three",
		"localhost:9000/search?q=three",
		"localhost:9000/search?q=four",
		"localhost:9000/search?q=four",
		"localhost:9000/search?q=four",
		"localhost:9000/search?q=four",
	} {
		if err := exec.Command("curl", url).Run(); err != nil {
			t.Fatalf("curl %s: %v", url, err)
		}
	}

	// Fetch metrics.
	if err := exec.Command("weaver", "multi", "metrics", "cache").Run(); err != nil {
		t.Fatalf("weaver multi metrics cache: %v", err)
	}
}
