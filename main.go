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
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ServiceWeaver/weaver"
)

var (
	// The address on which the HTTP server serves locally. 12347 is prime ;)
	localAddress = flag.String("address", "localhost:12347", "Local server address")

	//go:embed index.html
	indexHtml string // index.html served on "/"
)

func main() {
	flag.Parse()
	if err := weaver.Run(context.Background(), serve); err != nil {
		panic(err)
	}
}

// app is the main component of our application.
type app struct {
	weaver.Implements[weaver.Main]
	factorer weaver.Ref[Factorer]
	chatgpt  weaver.Ref[ChatGPT]
}

// serve serves HTTP traffic for the main component.
func serve(ctx context.Context, app *app) error {
	opts := weaver.ListenerOptions{LocalAddress: *localAddress}
	lis, err := app.Listener("primes", opts)
	if err != nil {
		return err
	}
	app.Logger().Info("Primes listener available.", "addr", lis)

	http.HandleFunc("/", app.handleRoot)
	http.HandleFunc("/factor", app.handleFactor)
	http.HandleFunc("/chatgpt_factor", app.handleChatGPTFactor)
	return http.Serve(lis, nil)
}

// handleRoot handles HTTP requests to the / endpoint.
func (a *app) handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, indexHtml)
}

// handleFactor handles HTTP requests to the /factor?x=<number> endpoint.
func (a *app) handleFactor(w http.ResponseWriter, r *http.Request) {
	// Parse GET request query parameter x.
	s := r.URL.Query().Get("x")
	x, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Compute the prime factorization of x.
	factors, err := a.factorer.Get().Factor(r.Context(), x)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the JSON serialized factorization.
	bytes, err := json.Marshal(factors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, string(bytes))
}

// handleChatGPTFactor handles HTTP requests to the /chatgpt_factor?x=<number>
// endpoint.
func (a *app) handleChatGPTFactor(w http.ResponseWriter, r *http.Request) {
	// Parse GET request query parameter x.
	s := r.URL.Query().Get("x")
	x, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query ChatGPT via the ChatGPT component.
	prompt := fmt.Sprintf("The prime factors of %d are", x)
	response, err := a.chatgpt.Get().Complete(r.Context(), prompt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, response)
}
