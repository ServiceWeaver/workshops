# A Cache Component

In this part, you'll implement a cache to speed up repeated requests for the
same prime factorization. Review [the documentation on component
semantics][component_semantics]. Then, in a file called `cache.go`, implement a
`Cache` component with the following interface:

```go
// Cache caches the prime factorizations of integers.
type Cache interface {
    // Get returns the cached prime factorization of the provided integer. On
    // cache miss, Get returns nil, nil.
    Get(context.Context, int) ([]int, error)

    // Put stores a prime factorization in the cache.
    Put(context.Context, int, []int) error
}
```

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Next, update your `Factorer` component to use the cache. When the `Factor`
method is called on integer `x`, it should first see if the prime factorization
of `x` is in the cache. If it is, `Factor` should return the prime factorization
immediately. Otherwise, `Factor` should compute the prime factorization and
store it in the cache.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Run your application using `go run .` (don't use `weaver multi deploy`).

```
$ weaver generate .
$ go run .
```

In a separate terminal, curl the application for the prime factorization of
`777767777`.

```
$ curl localhost:12347/factor?x=777767777
[777767777]
```

Your application should return `[777767777]` but because `777767777` is prime,
it might take a couple of seconds. Then, re-run the same `curl` command. This
time, the request should return nearly instantly, as the prime factorization of
`777767777` should now be in the cache.

```
$ curl localhost:12347/factor?x=777767777
[777767777]
```

Now, run your application using `weaver multi deploy`:

```
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

Surprisingly, every request is slow! Add some logs to the `Get` and `Put`
methods and see if you can figure out why this surprising behavior is happening.

<details>
<summary>Solution.</summary>

Here are the logs produced by our application after curling it four times for
the prime factorization of `777767777`:

```
primes.Factorer a3b14619 factorer.go:53] Factor x="777767777"
primes.Cache    7545ca78 cache.go:51   ] Get x="777767777"
primes.Cache    ef518f6b cache.go:58   ] Put factors="[777767777]" x="777767777"
primes.Factorer a3b14619 factorer.go:53] Factor x="777767777"
primes.Cache    7545ca78 cache.go:51   ] Get x="777767777"
primes.Cache    ef518f6b cache.go:58   ] Put factors="[777767777]" x="777767777"
primes.Factorer 7ed120c4 factorer.go:53] Factor x="777767777"
primes.Cache    7545ca78 cache.go:51   ] Get x="777767777"
primes.Cache    ef518f6b cache.go:58   ] Put factors="[777767777]" x="777767777"
primes.Factorer a3b14619 factorer.go:53] Factor x="777767777"
primes.Cache    7545ca78 cache.go:51   ] Get x="777767777"
primes.Cache    ef518f6b cache.go:58   ] Put factors="[777767777]" x="777767777"
```

You'll notice that there are two replicas of the `Cache` component: `7545ca78`
and `ef518f6b`. Method calls to the `Cache` component are being routed round
robin across these two replicas.

When `Factor` is first called on `777767777`, it calls `Get` to see if the prime
factorization of `777767777` is in the cache. This `Get` is routed to replica
`7545ca78`. The prime factorization is not in replica `7545ca78`'s cache, so
`Factor` computes the prime factorization and stores it in the cache by calling
`Put`. This `Put` is routed to the other replica, `ef518f6b`, where the prime
factorization is stored in the cache.

When we call `Factor` on `777767777` again, it again calls `Get` to see if the
prime factorization of `777767777` is in the cache, and this `Get` is again
routed to replica `7545ca78`. The prime factorization has been cached at replica
`ef518f6b` but not at replica `7545ca78`, so `Factor` computes the prime
factorization and stores it in the cache by calling `Put`. This `Put` is again
routed to the replica `ef518f6b` where it is redundantly cached.

This repeats for every call to `Factor`. Because `Get`s and `Put`s are routed to
two different replicas, every request results in a cache miss.
</details>

In the next part, we'll see how to prevent this surprising behavior from
happening.

[**:arrow_left: Previous Part**](../06)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../08)

[component_semantics]: https://serviceweaver.dev/docs.html#components-semantics
