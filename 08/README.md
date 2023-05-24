# Routing the Cache

In this part, you'll route requests to the `Cache` component so that repeated
requests for the same integer `x` are routed to the same replica. Review [the
documentation on routing][routing]. In `cache.go`, implement a routing struct
called `router` that routes `Get` and `Put` requests using the integer being
factored as the routing key.

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

And again in a separate terminal, repeatedly curl the application for the prime
factorization of `777767777`:

```
$ curl localhost:12347/factor?x=777767777
[777767777]
$ curl localhost:12347/factor?x=777767777
[777767777]
$ curl localhost:12347/factor?x=777767777
[777767777]
$ curl localhost:12347/factor?x=777767777
[777767777]
```

The first request should be slow, but all subsequent requests should complete
nearly instantly. If you look at your application logs, you can confirm that
`Get` and `Put` requests for the same integer `x` are routed to the same `Cache`
replica. Here are the logs for the `/factor?x=777767777` requests above:

```
primes.Factorer 1da63c0a factorer.go:53] Factor x="777767777"
primes.Cache    e1ef982f cache.go:51   ] Get x="777767777"
primes.Cache    e1ef982f cache.go:58   ] Put factors="[777767777]" x="777767777"
primes.Factorer 2dc51d83 factorer.go:53] Factor x="777767777"
primes.Cache    e1ef982f cache.go:51   ] Get x="777767777"
primes.Factorer 1da63c0a factorer.go:53] Factor x="777767777"
primes.Cache    e1ef982f cache.go:51   ] Get x="777767777"
primes.Factorer 2dc51d83 factorer.go:53] Factor x="777767777"
primes.Cache    e1ef982f cache.go:51   ] Get x="777767777"
```

Notice that every `Get` and `Put` is routed to replica `e1ef982f`.

[**:arrow_left: Previous Part**](../07)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../09)

[routing]: https://serviceweaver.dev/docs.html#routing
