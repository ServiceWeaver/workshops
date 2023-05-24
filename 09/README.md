# Metrics

In this part, you'll add metrics to your application to measure the number of
cache hits and misses.

Review [the documentation on metrics][metrics]. In `factorer.go`, declare two
counters called `"factor_cache_hits"` and `"factor_cache_misses"` at package
scope.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Inside the `Factor` method, increment the `"factor_cache_hits"` counter whenever
there is a cache hit, and increment the `"factor_cache_misses"` metric whenever
there is a cache miss.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Build and run your application using `weaver multi deploy`:

```
$ go build .
$ weaver multi deploy config.toml
```

In a separate terminal, curl your application with various integers, making sure
to repeat some requests.

```
$ curl localhost:12347/factor?x=2 # MISS 1
$ curl localhost:12347/factor?x=2 # HIT  1
$ curl localhost:12347/factor?x=3 # MISS 2
$ curl localhost:12347/factor?x=3 # HIT  2
$ curl localhost:12347/factor?x=3 # HIT  3
$ curl localhost:12347/factor?x=4 # MISS 3
$ curl localhost:12347/factor?x=4 # HIT  4
$ curl localhost:12347/factor?x=4 # HIT  5
$ curl localhost:12347/factor?x=4 # HIT  6
```

Run `weaver multi metrics` to see a snapshot of the metric values.

```
$ weaver multi metrics factor_cache
╭────────────────────────────────────────────────────────────────────────╮
│ // Number of Factor cache hits                                         │
│ factor_cache_hits: COUNTER                                             │
├───────────────────┬────────────────────┬───────────────────────┬───────┤
│ serviceweaver_app │ serviceweaver_node │ serviceweaver_version │ Value │
├───────────────────┼────────────────────┼───────────────────────┼───────┤
│ workshops         │ 03503d27           │ b6e10f77              │ 0     │
│ workshops         │ 469ede58           │ b6e10f77              │ 0     │
│ workshops         │ 57a50598           │ b6e10f77              │ 0     │
│ workshops         │ 7b4e1944           │ b6e10f77              │ 0     │
│ workshops         │ 893c7a35           │ b6e10f77              │ 3     │
│ workshops         │ 9fff72df           │ b6e10f77              │ 3     │
╰───────────────────┴────────────────────┴───────────────────────┴───────╯
╭────────────────────────────────────────────────────────────────────────╮
│ // Number of Factor cache misses                                       │
│ factor_cache_misses: COUNTER                                           │
├───────────────────┬────────────────────┬───────────────────────┬───────┤
│ serviceweaver_app │ serviceweaver_node │ serviceweaver_version │ Value │
├───────────────────┼────────────────────┼───────────────────────┼───────┤
│ workshops         │ 03503d27           │ b6e10f77              │ 0     │
│ workshops         │ 469ede58           │ b6e10f77              │ 0     │
│ workshops         │ 57a50598           │ b6e10f77              │ 0     │
│ workshops         │ 7b4e1944           │ b6e10f77              │ 0     │
│ workshops         │ 893c7a35           │ b6e10f77              │ 1     │
│ workshops         │ 9fff72df           │ b6e10f77              │ 2     │
╰───────────────────┴────────────────────┴───────────────────────┴───────╯
```

`weaver multi metrics` shows the current value of every metric on all replicas.
Your application has six replicas&mdash;two replicas of `weaver.Main`, two
replicas of `Factorer`, and two replicas of `Cache`&mdash;which is why you see
six entries for each metric. Only two of these entries, the two replicas of
`Factorer`, should report non-zero values.

**Question.** Based on the `curl` requests above and the output of `weaver multi
metrics`, how were the `curl` requests routed to replicas `893c7a35` and
`9fff72df`?

<details>
<summary>Solution.</summary>

Replica `893c7a35` has one cache miss and three cache hits, meaning that the
requests for the prime factorization of `4` were routed to `893c7a35`. Replica
`9fff72df` then received requests for the prime factorizations of `2` and `3`,
resulting in two cache misses and three cache hits.
</details>

Refer to the documentation on [single process metrics][single_process_metrics]
and [multiprocess metrics][multiprocess_metrics] for instructions on how to view
metrics using Prometheus.

[**:arrow_left: Previous Part**](../08)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../10)

[metrics]: https://serviceweaver.dev/docs.html#metrics
[single_process_metrics]: https://serviceweaver.dev/docs.html#single-process-metrics
[multiprocess_metrics]: https://serviceweaver.dev/docs.html#multiprocess-metrics
