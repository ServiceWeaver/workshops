# Routing the Cache

In this part, you'll route requests to the `Cache` component so that repeated
requests for the same query are routed to the same replica. Review [the
documentation on routing][routing]. In `cache.go`, implement a routing struct
called `router` that routes `Get` and `Put` requests using the query as the
routing key.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/e9c0573de0f20fca6a88106ad9f25fddf2f04233/08/cache.go#L63-L74
</details>

Embed `weaver.WithRouter[router]` in your cache implementation to enable
routing.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/e9c0573de0f20fca6a88106ad9f25fddf2f04233/08/cache.go#L34-L41
</details>

Build and run your application using `weaver multi deploy`:

```
$ weaver generate .
$ go build .
$ weaver multi deploy config.toml
```

And again in a separate terminal, repeatedly curl the application.

```
$ curl localhost:9000/search?q=pig
["🐖","🐗","🐷","🐽"]
$ curl localhost:9000/search?q=pig
["🐖","🐗","🐷","🐽"]
$ curl localhost:9000/search?q=pig
["🐖","🐗","🐷","🐽"]
$ curl localhost:9000/search?q=pig
["🐖","🐗","🐷","🐽"]
```

The first request should be slow, but all subsequent requests should complete
nearly instantly. If you look at your application logs, you can confirm that
`Get` and `Put` requests for the query `"pig"` are routed to the same `Cache`
replica. Here are the logs for the `/search?q=pig` requests above:

```
emojis.Searcher 1da63c0a searcher.go:53] Search query="pig"
emojis.Cache    e1ef982f cache.go:51   ] Get query="pig"
emojis.Cache    e1ef982f cache.go:58   ] Put query="pig"
emojis.Searcher 2dc51d83 searcher.go:53] Search query="pig"
emojis.Cache    e1ef982f cache.go:51   ] Get query="pig"
emojis.Searcher 1da63c0a searcher.go:53] Search query="pig"
emojis.Cache    e1ef982f cache.go:51   ] Get query="pig"
emojis.Searcher 2dc51d83 searcher.go:53] Search query="pig"
emojis.Cache    e1ef982f cache.go:51   ] Get query="pig"
```

Notice that every `Get` and `Put` is routed to replica `e1ef982f`. At this
point, feel free to remove the `time.Sleep(time.Second)` call from your code.

[**:arrow_left: Previous Part**](../07)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../09)

[routing]: https://serviceweaver.dev/docs.html#routing
