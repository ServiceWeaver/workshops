# Multiprocess Execution

In this part, you'll deploy your application across multiple processes. Review
[the documentation on multiprocess execution][multiprocess]. Create a config
file called `config.toml` with the following contents.

```toml
[serviceweaver]
binary = "./emojis"
```

This config file specifies the binary of your Service Weaver application. Next,
build and run your app using `weaver multi deploy`.

```
$ go build .
$ weaver multi deploy config.toml
```

In a separate terminal, run `weaver multi status` to see information about the
application.

```
$ weaver multi status
╭─────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                         │
├────────┬──────────────────────────────────────┬─────┤
│ APP    │ DEPLOYMENT                           │ AGE │
├────────┼──────────────────────────────────────┼─────┤
│ emojis │ 751d7710-f5e3-4428-8f4b-dfcb1ff64d69 │ 9s  │
╰────────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────────╮
│ COMPONENTS                                               │
├────────┬────────────┬─────────────────┬──────────────────┤
│ APP    │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS     │
├────────┼────────────┼─────────────────┼──────────────────┤
│ emojis │ 751d7710   │ weaver.Main     │ 3723743, 3723753 │
│ emojis │ 751d7710   │ emojis.Searcher │ 3723765, 3723777 │
╰────────┴────────────┴─────────────────┴──────────────────╯
╭──────────────────────────────────────────────────╮
│ LISTENERS                                        │
├────────┬────────────┬──────────┬─────────────────┤
│ APP    │ DEPLOYMENT │ LISTENER │ ADDRESS         │
├────────┼────────────┼──────────┼─────────────────┤
│ emojis │ 751d7710   │ emojis   │ 127.0.0.1:12347 │
╰────────┴────────────┴──────────┴─────────────────╯
```

When you deploy an application with `weaver multi deploy`, every component is
replicated twice, and every replica is run in a separate OS process. We can
infer this from the output of `weaver multi status` because the two components,
`weaver.Main` and `primes.Factorer`, each have two process ids, and all four
process ids are distinct.

Curl your application couple of times:

```
$ curl localhost:9000/search?query=sushi
["🍣"]
$ curl localhost:9000/search?query=curry
["🍛"]
$ curl localhost:9000/search?query=shrimp
["🍤", "🦐"]
$ curl localhost:9000/search?query=lobster
["🦞"]
```

The application should print out logs that look something like the following.
Note that requests are balanced across the two replicas of the `Searcher`
component (`d29c6296` and `23ebba75`).

```
D0525 09:40:32.466445 emojis.Searcher d29c6296 searcher.go:53] Search query="sushi"
D0525 09:40:33.317303 emojis.Searcher 23ebba75 searcher.go:53] Search query="curry"
D0525 09:40:35.433576 emojis.Searcher d29c6296 searcher.go:53] Search query="shrimp"
D0525 09:40:36.534745 emojis.Searcher d29c6296 searcher.go:53] Search query="lobster"
```

[**:arrow_left: Previous Part**](../05)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../07)

[multiprocess]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-multiprocess-execution
