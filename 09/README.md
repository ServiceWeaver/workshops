# Metrics

In this part, you'll add metrics to your application to measure the number of
cache hits and misses. Review [the documentation on metrics][metrics]. In
`searcher.go`, declare two counters called `"search_cache_hits"` and
`"search_cache_misses"` at package scope.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/912c215cecd611feadd2e23fcc80fe09f4af2045/09/searcher.go#L27-L37
</details>

Inside the `Search` method, increment the `"search_cache_hits"` counter whenever
there is a cache hit, and increment the `"search_cache_misses"` metric whenever
there is a cache miss.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/912c215cecd611feadd2e23fcc80fe09f4af2045/09/searcher.go#L54-L63
</details>

Build and run your application using `weaver multi deploy`:

```
$ go build .
$ weaver multi deploy config.toml
```

In a separate terminal, curl your application with various queries, making sure
to repeat some requests.

```
$ curl "localhost:9000/search?q=two"    # MISS 1
$ curl "localhost:9000/search?q=two"    # HIT  1
$ curl "localhost:9000/search?q=three"  # MISS 2
$ curl "localhost:9000/search?q=three"  # HIT  2
$ curl "localhost:9000/search?q=three"  # HIT  3
$ curl "localhost:9000/search?q=four"   # MISS 3
$ curl "localhost:9000/search?q=four"   # HIT  4
$ curl "localhost:9000/search?q=four"   # HIT  5
$ curl "localhost:9000/search?q=four"   # HIT  6
```

Run `weaver multi metrics` to see a snapshot of the metric values.

```
$ weaver multi metrics cache
╭────────────────────────────────────────────────────────────────────────╮
│ // Number of Search cache hits                                         │
│ search_cache_hits: COUNTER                                             │
├───────────────────┬────────────────────┬───────────────────────┬───────┤
│ serviceweaver_app │ serviceweaver_node │ serviceweaver_version │ Value │
├───────────────────┼────────────────────┼───────────────────────┼───────┤
│ workshops         │ 8e6a334e           │ dbfaaaa8              │ 4     │
│ workshops         │ a3cda9a8           │ dbfaaaa8              │ 0     │
│ workshops         │ a53facb6           │ dbfaaaa8              │ 0     │
│ workshops         │ b1e20fcf           │ dbfaaaa8              │ 0     │
│ workshops         │ f3a32208           │ dbfaaaa8              │ 2     │
│ workshops         │ ff25d22e           │ dbfaaaa8              │ 0     │
╰───────────────────┴────────────────────┴───────────────────────┴───────╯
╭────────────────────────────────────────────────────────────────────────╮
│ // Number of Search cache misses                                       │
│ search_cache_misses: COUNTER                                           │
├───────────────────┬────────────────────┬───────────────────────┬───────┤
│ serviceweaver_app │ serviceweaver_node │ serviceweaver_version │ Value │
├───────────────────┼────────────────────┼───────────────────────┼───────┤
│ workshops         │ 8e6a334e           │ dbfaaaa8              │ 1     │
│ workshops         │ a3cda9a8           │ dbfaaaa8              │ 0     │
│ workshops         │ a53facb6           │ dbfaaaa8              │ 0     │
│ workshops         │ b1e20fcf           │ dbfaaaa8              │ 0     │
│ workshops         │ f3a32208           │ dbfaaaa8              │ 2     │
│ workshops         │ ff25d22e           │ dbfaaaa8              │ 0     │
╰───────────────────┴────────────────────┴───────────────────────┴───────╯
```

`weaver multi metrics` shows the current value of every metric on all replicas.
Your application has six replicas&mdash;two replicas of `weaver.Main`, two
replicas of `Searcher`, and two replicas of `Cache`&mdash;which is why you see
six entries for each metric. Only two of these entries, the two replicas of
`Searcher`, should report non-zero values.

Refer to the documentation on [single process metrics][single_process_metrics]
and [multiprocess metrics][multiprocess_metrics] for instructions on how to view
metrics using Prometheus.

[**:arrow_left: Previous Part**](../08)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../10)

[metrics]: https://serviceweaver.dev/docs.html#metrics
[single_process_metrics]: https://serviceweaver.dev/docs.html#single-process-metrics
[multiprocess_metrics]: https://serviceweaver.dev/docs.html#multiprocess-metrics
