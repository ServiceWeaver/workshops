# Logging

In this part, you'll add logging to your application. Review [the documentation
on logging][logging]. Add a `Debug("Factor", "x", x)` call to the beginning of
the `Factor` method in `factorer.go`.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Re-run your application.

```
go run .
```

In a separate terminal, curl the application a couple of times:

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

Your application should output logs that look something like this:

```
D0524 13:37:03.295982 primes.Factorer d4ad11e1 factorer.go:53] Factor x="1"
D0524 13:37:04.148701 primes.Factorer d4ad11e1 factorer.go:53] Factor x="2"
D0524 13:37:04.937959 primes.Factorer d4ad11e1 factorer.go:53] Factor x="3"
D0524 13:37:06.577844 primes.Factorer d4ad11e1 factorer.go:53] Factor x="4"
D0524 13:37:07.840446 primes.Factorer d4ad11e1 factorer.go:53] Factor x="5"
```

The first character of a log line indicates whether the log is a [D]ebug,
[I]nfo, or [E]rror log entry. Then comes the date in MMDD format, followed by
the time. Then comes the component name followed by a logical node id. If two
components are co-located in the same OS process, they are given the same node
id. Then comes the file and line where the log was produced, followed finally by
the contents of the log.

[**:arrow_left: Previous Part**](../04)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../06)

[logging]: https://serviceweaver.dev/docs.html#logging
