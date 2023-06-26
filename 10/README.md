# ChatGPT

In this part, you will improve your emoji search engine using [OpenAI's GPT
API][gpt_api]. Thus far, your emoji search engine implements *lexical search*.
Given a query, it looks for all emojis with labels that match the literal
strings in the query. In this part, you will use ChatGPT to implement *semantic
search*. The search engine will return emojis that match the meaning of the
query.  Specifically, when you get a query, you will ask ChatGPT for emojis
related to the query. For the query "summer vibes", for example, ChatGPT returns
emojis like 🌞, 🌴, and 🍹. These emojis match the meaning of "summer vibes"
even though they aren't labeled "summer" or "vibes".

## OpenAI Account

First, visit https://platform.openai.com/signup and create an OpenAI account.
OpenAI's GPT API is not free to use, but as of this writing, new OpenAI accounts
automatically receive [$5 in free credit][pricing]. If you're reading this and
the promotion no longer exists, you'll have to set up [billing][billing] and
purchase credits.

Next, visit https://platform.openai.com/account/api-keys and click "Create new
secret key". When prompted, enter a name for the secret key (e.g., "emojis") and
click "Create secret key". Finally, OpenAI will generate and show you your
secret key. Copy this key and store it somewhere, as you won't be able to view
the secret key again once you click "Done". If you forget to copy the key,
simply delete it and create a new one.

⚠️ Note that this key is secret. Don't share it with anyone and **be careful not
to push it to a public repository**. ⚠️

## ChatGPT Component

Next, in a file called `chatgpt.go`, write a component called `ChatGPT` with the
following interface:

```go
// ChatGPT is a frontend to OpenAI's ChatGPT API.
type ChatGPT interface {
    // Complete returns the ChatGPT completion of the provided prompt.
    Complete(ctx context.Context, prompt string) (string, error)
}
```

The `Complete` method receives a prompt (e.g., `"1 + 1 = "`) and returns
ChatGPT's completion of the prompt (e.g., `"2"`). To implement the method, we
recommend using [sashabaranov/go-openai][go-openai], which is [OpenAI's
officially recommended Go library][openai_go]. In the implementation of the
`Complete` method, first create an OpenAI client by passing your secret key to
the [`openai.NewClient`][NewClient] function. Hard code the secret key for now;
later we'll provide the key via a config file. Next, call the client's
[`CreateChatCompletion`][CreateChatCompletion] method on the following
[`ChatCompletionRequest`][ChatCompletionRequest] to receive ChatGPT's completion
of the user provided prompt.

```go
req := openai.ChatCompletionRequest{
    Model: openai.GPT3Dot5Turbo,
    Messages: []openai.ChatCompletionMessage{
        {Role: openai.ChatMessageRoleUser, Content: prompt},
    },
}
```

Finally, given a response `resp` from `ChatGPT`, return the completion
`resp.Choices[0].Message.Content`.

<details>
<summary>Solution.</summary>

**NOTE** that this solution includes some configuration code (e.g.,
`weaver.WithConfig`) that won't make sense until the next section. Feel free to
ignore it for now.

https://github.com/ServiceWeaver/workshops/blob/c3e81d5c15ff9349b2c8d0f7da8a9f49607533e4/10/chatgpt.go#L25-L65
</details>

## Config

Now, you'll change the `ChatGPT` component to receive your OpenAI secret key via
a config file rather than having to hard code it in the source code. Review [the
documentation on configuring components][config]. Add a `config` struct to
`chatgpt.go` with an `APIKey` field. Use a `toml:"api_key"` annotation to
indicate that the field will be written `api_key` in the config file.

```golang
// config configures the chatgpt component implementation.
type config struct {
    // OpenAI API key. You can generate an API key at
    // https://platform.openai.com/account/api-keys.
    APIKey string `toml:"api_key"`
}
```

Then, embed `weaver.WithConfig[config]` in the struct that implements the
`ChatGPT` component.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/c3e81d5c15ff9349b2c8d0f7da8a9f49607533e4/10/chatgpt.go#L31-L42
</details>

Next, add your secret key to the `api_key` field in the `"emojis/ChatGPT"`
section of `config.toml`. Again, **be careful not to push your secret key to a
public repository**.

```toml
[serviceweaver]
binary = "./emojis"

["emojis/ChatGPT"]
api_key = "YOUR SECRET KEY GOES HERE"
```

Finally, update the implementation of the `Complete` method to read your secret
key from the config file, returning an error if the key is missing.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/c3e81d5c15ff9349b2c8d0f7da8a9f49607533e4/10/chatgpt.go#L45-L51
</details>

## Putting It All Together

In `searcher.go`, add a `SearchChatGPT(context.Context, query string) ([]string,
error)` method to the `Searcher` component. The `SearchChatGPT` method acts like
the `Search` method&mdash;it receives a query and returns a list of matching
emojis&mdash;but it returns the emojis suggested by ChatGPT.

To implement the `SearchChatGPT` method, form a ChatGPT prompt from the user's
query and pass it to the `ChatGPT.Complete` method. You can get creative with
the wording of your prompt, but we went with `"Give me a list of emojis that
related to the query $QUERY. Don't give an explanation."`.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/c3e81d5c15ff9349b2c8d0f7da8a9f49607533e4/10/searcher.go#L110-L117
</details>

Next, parse and return the emojis from ChatGPT's response. This is surprisingly
tricky, as some emojis are composed of multiple unicode code points. We
recommend using the https://github.com/rivo/uniseg library to segment ChatGPT's
response into a set of graphemes and checking to see if each grapheme is in the
`emojis` map. Here's our solution:

https://github.com/ServiceWeaver/workshops/blob/c3e81d5c15ff9349b2c8d0f7da8a9f49607533e4/10/searcher.go#L128-L138

In `main.go`, add a `/search_chatgpt?q=<query>` endpoint that calls the
`SearchChatGPT` method.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/c3e81d5c15ff9349b2c8d0f7da8a9f49607533e4/10/main.go#L57-L59
</details>

Finally, build and run your application:

```
$ weaver generate .
$ weaver multi deploy config.toml
```

You can curl the `/search_chatgpt` endpoint directly:

```console
$ curl localhost:9000/search_chatgpt?q=happy
["😀","😁","😃","😄","😊","😍","😎","🤗","😻","🌞","🎉","🎊","🎁","🎈","💐","👍","✨","🌟","💫","🌈"]
```

Though, we strongly recommend you use the web UI provided in [Part 4](../04).
The UI sends every query to the `/search` and `/search_chatgpt` endpoints in
parallel and merges the results together.

[emoji_search_demo.webm](https://github.com/ServiceWeaver/workshops/assets/3654277/cde50b36-7808-4c26-983d-54a37532e69a)

[**:arrow_left: Previous Part**](../09)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Home :house:**](..)

[ChatCompletionRequest]: https://pkg.go.dev/github.com/sashabaranov/go-openai#ChatCompletionRequest
[CreateChatCompletion]: https://pkg.go.dev/github.com/sashabaranov/go-openai#Client.CreateChatCompletion
[NewClient]: https://pkg.go.dev/github.com/sashabaranov/go-openai#NewClient
[billing]: https://platform.openai.com/account/billing/overview
[config]: https://serviceweaver.dev/docs.html#components-config
[go-openai]: https://github.com/sashabaranov/go-openai
[gpt_api]: https://platform.openai.com/docs/guides/gpt
[openai_go]: https://platform.openai.com/docs/libraries/go
[pricing]: https://openai.com/pricing
