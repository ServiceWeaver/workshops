# Routing the Cache

In this part, you'll route requests to the `Cache` component so that repeated
requests for the same query are routed to the same replica. Review [the
documentation on routing][routing]. In `cache.go`, implement a routing struct
called `router` that routes `Get` and `Put` requests using the query as the
routing key.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Embed `weaver.WithRouter[router]` in your cache implementation to enable
routing.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Build and run your application using `weaver multi deploy`:

```
$ weaver generate .
$ go build .
$ weaver multi deploy config.toml
```

And again in a separate terminal, repeatedly curl the application.

```
$ curl localhost:9000/search?q=rock
["☘️","🚀","🪨"]
$ curl localhost:9000/search?q=rock
["☘️","🚀","🪨"]
$ curl localhost:9000/search?q=rock
["☘️","🚀","🪨"]
$ curl localhost:9000/search?q=rock
["☘️","🚀","🪨"]
```

The first request should be slow, but all subsequent requests should complete
nearly instantly. If you look at your application logs, you can confirm that
`Get` and `Put` requests for the same query are routed to the same `Cache`
replica. Here are the logs for the `/search?q=rock` requests above:

```
emojis.Searcher 1da63c0a searcher.go:53] Search query="rock"
emojis.Cache    e1ef982f cache.go:51   ] Get query="rock"
emojis.Cache    e1ef982f cache.go:58   ] Put query="rock"
emojis.Searcher 2dc51d83 searcher.go:53] Search query="rock"
emojis.Cache    e1ef982f cache.go:51   ] Get query="rock"
emojis.Searcher 1da63c0a searcher.go:53] Search query="rock"
emojis.Cache    e1ef982f cache.go:51   ] Get query="rock"
emojis.Searcher 2dc51d83 searcher.go:53] Search query="rock"
emojis.Cache    e1ef982f cache.go:51   ] Get query="rock"
```

Notice that every `Get` and `Put` is routed to replica `e1ef982f`. At this
point, feel free to remove the `time.Sleep(time.Second)` call from your code.

[**:arrow_left: Previous Part**](../07)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../09)

[routing]: https://serviceweaver.dev/docs.html#routing
