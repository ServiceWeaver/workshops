# Multiprocess Execution

In this part, you'll deploy your application across multiple processes. Review
[the documentation on multiprocess execution][multiprocess]. Create a config
file called `config.toml` with the following contents.

```toml
[serviceweaver]
binary = "./primes"
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
│ primes │ 751d7710-f5e3-4428-8f4b-dfcb1ff64d69 │ 9s  │
╰────────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────────╮
│ COMPONENTS                                               │
├────────┬────────────┬─────────────────┬──────────────────┤
│ APP    │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS     │
├────────┼────────────┼─────────────────┼──────────────────┤
│ primes │ 751d7710   │ weaver.Main     │ 3723743, 3723753 │
│ primes │ 751d7710   │ primes.Factorer │ 3723765, 3723777 │
╰────────┴────────────┴─────────────────┴──────────────────╯
╭──────────────────────────────────────────────────╮
│ LISTENERS                                        │
├────────┬────────────┬──────────┬─────────────────┤
│ APP    │ DEPLOYMENT │ LISTENER │ ADDRESS         │
├────────┼────────────┼──────────┼─────────────────┤
│ primes │ 751d7710   │ primes   │ 127.0.0.1:12347 │
╰────────┴────────────┴──────────┴─────────────────╯
```

When you deploy an application with `weaver multi deploy`, every component is
replicated twice, and every replica is run in a separate OS process. We can
infer this from the output of `weaver multi status` because the two components,
`weaver.Main` and `primes.Factorer`, each have two process ids, and all four
process ids are distinct.

Curl your application couple of times:

```
$ curl localhost:12347/factor?x=1
[1]
$ curl localhost:12347/factor?x=2
[2]
$ curl localhost:12347/factor?x=3
[3]
$ curl localhost:12347/factor?x=4
[2,2]
$ curl localhost:12347/factor?x=5
[5]
```

The application should print out logs that look something like the following.
Note that requests are balanced across the two replicas of the `Factorer`
component (`d29c6296` and `23ebba75`).

```
D0525 09:40:32.466445 primes.Factorer d29c6296 factorer.go:53] Factor x="1"
D0525 09:40:33.317303 primes.Factorer 23ebba75 factorer.go:53] Factor x="2"
D0525 09:40:35.433576 primes.Factorer d29c6296 factorer.go:53] Factor x="3"
D0525 09:40:36.534745 primes.Factorer d29c6296 factorer.go:53] Factor x="4"
D0525 09:40:38.562967 primes.Factorer 23ebba75 factorer.go:53] Factor x="5"
```

[**:arrow_left: Previous Part**](../05)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../07)

[multiprocess]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-multiprocess-execution
