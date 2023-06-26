# A Cache Component

In this part, you'll implement a cache to speed up repeated requests for the
same query. Review [the documentation on component
semantics][component_semantics]. Then, in a file called `cache.go`, implement a
`Cache` component with the following interface:

```go
// Cache caches query results.
type Cache interface {
    // Get returns the cached emojis produced by the provided query. On cache
    // miss, Get returns nil, nil.
    Get(context.Context, string) ([]string, error)

    // Put stores a query and its corresponding emojis in the cache.
    Put(context.Context, string, []string) error
}
```

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/62322cde0019ad7c3c02804590f342291aebccf2/07/cache.go#L15-L60
</details>

Next, update your `Searcher` component to use the cache. When the `Search`
method is called on a `query`, it should first see if the matching emojis for
`query` are in the cache by calling `Get`. If they are, `Search` should return
the emojis immediately. Otherwise, `Search` should find the emojis and store
them in the cache by calling `Put`.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/62322cde0019ad7c3c02804590f342291aebccf2/07/searcher.go#L32-L36
https://github.com/ServiceWeaver/workshops/blob/62322cde0019ad7c3c02804590f342291aebccf2/07/searcher.go#L41-L47
https://github.com/ServiceWeaver/workshops/blob/62322cde0019ad7c3c02804590f342291aebccf2/07/searcher.go#L69-L72
</details>

Because our basic search algorithm is already quite fast, it's hard to notice
the speedup from caching. To simulate the slowdown of implementing a more
advanced search algorithm, temporarily add `time.Sleep(time.Second)` to the
`Search` method right after checking the cache and right before performing the
search.

Now, run your application using `go run .` (don't use `weaver multi deploy`).

```
$ weaver generate .
$ SERVICEWEAVER_CONFIG=config.toml go run .
```

In a separate terminal, curl the application with query `"pig"`.

```
$ curl localhost:9000/search?q=pig
["🐖","🐗","🐷","🐽"]
```

Your application should return `["🐖","🐗","🐷","🐽"]` after a one second delay.
Then, re-run the same curl command. This time, the request should return nearly
instantly, as the results of query `"pig"` are now in the cache.

```
$ curl localhost:9000/search?q=pig
["🐖","🐗","🐷","🐽"]
```

Now, run your application using `weaver multi deploy`:

```
$ go build .
$ weaver multi deploy config.toml
```

And again in a separate terminal, repeatedly curl the application with query
`"pig"`:

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

Surprisingly, every request is slow! Add some logs to the `Get` and `Put`
methods of the `Cache` component and see if you can figure out why this
surprising behavior is happening. Keep reading for an explanation.

Here are the logs produced by our application after curling it four times with
query `"pig"`:

```
emojis.Searcher a3b14619 searcher.go:53] Search query="pig"
emojis.Cache    7545ca78 cache.go:51   ] Get query="pig"
emojis.Cache    ef518f6b cache.go:58   ] Put query="pig"
emojis.Searcher a3b14619 searcher.go:53] Search query="pig"
emojis.Cache    7545ca78 cache.go:51   ] Get query="pig"
emojis.Cache    ef518f6b cache.go:58   ] Put query="pig"
emojis.Searcher 7ed120c4 searcher.go:53] Search query="pig"
emojis.Cache    7545ca78 cache.go:51   ] Get query="pig"
emojis.Cache    ef518f6b cache.go:58   ] Put query="pig"
emojis.Searcher a3b14619 searcher.go:53] Search query="pig"
emojis.Cache    7545ca78 cache.go:51   ] Get query="pig"
emojis.Cache    ef518f6b cache.go:58   ] Put query="pig"
```

You'll notice that there are two replicas of the `Cache` component: `7545ca78`
and `ef518f6b`. Method calls to the `Cache` component are being routed round
robin across these two replicas.

When `Search` is first called on `"pig"`, it calls `Get` to see if the matching
emojis are in the cache. This `Get` is routed to replica `7545ca78`. The emojis
are not in replica `7545ca78`'s cache, so `Search` computes the matching emojis
and stores them in the cache by calling `Put`. This `Put` is routed to the other
replica, `ef518f6b`, where the emojis are stored in the cache.

When we call `Search` on `"pig"` again, it again calls `Get` to see if the
matching emojis are in the cache, and this `Get` is again routed to replica
`7545ca78`. The emojis have been cached at replica `ef518f6b` but not at replica
`7545ca78`, so `Search` computes the emojis and stores them in the cache by
calling `Put`. This `Put` is again routed to the replica `ef518f6b` where it is
redundantly cached.

This repeats for every call to `Search`. Because `Get`s and `Put`s are routed to
two different replicas, every request results in a cache miss.

In the next part, we'll see how to prevent this surprising behavior from
happening.

[**:arrow_left: Previous Part**](../06)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../08)

[component_semantics]: https://serviceweaver.dev/docs.html#components-semantics
