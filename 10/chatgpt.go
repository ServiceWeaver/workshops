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

	"github.com/ServiceWeaver/weaver"
	openai "github.com/sashabaranov/go-openai"
)

// ChatGPT is a frontend to OpenAI's ChatGPT API.
type ChatGPT interface {
	// Complete returns the ChatGPT completion of the provided prompt.
	Complete(ctx context.Context, prompt string) (string, error)
}

// chatgpt implements the ChatGPT component.
type chatgpt struct {
	weaver.Implements[ChatGPT]
	weaver.WithConfig[config]
}

// config configures the chatgpt component implementation.
type config struct {
	// OpenAI API key. You can generate an API key at
	// https://platform.openai.com/account/api-keys.
	APIKey string `toml:"api_key"`
}

func (gpt *chatgpt) Complete(ctx context.Context, prompt string) (string, error) {
	// Check for an API key.
	if gpt.Config().APIKey == "" {
		return "", fmt.Errorf("ChatGPT api_key not provided")
	}

	// Issue the ChatGPT request.
	client := openai.NewClient(gpt.Config().APIKey)
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: prompt},
		},
	}
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", fmt.Errorf("ChatGPT completion error: %w", err)
	}

	// Return the completion.
	return resp.Choices[0].Message.Content, nil
}
