# Logging

In this part, you'll add logging to your application. Review [the documentation
on logging][logging]. Add a `Debug("Search", "query", query)` call to the
beginning of the `Search` method in `searcher.go`.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/5b26ed2f334b061315b49320cf9ee04fc0e009e3/05/searcher.go#L38
</details>

Re-run your application.

```
$ SERVICEWEAVER_CONFIG=config.toml go run .
```

In a separate terminal, curl the application a couple of times:

```
$ curl localhost:9000/search?q=sushi
["🍣"]
$ curl localhost:9000/search?q=curry
["🍛"]
$ curl localhost:9000/search?q=shrimp
["🍤","🦐"]
$ curl localhost:9000/search?q=lobster
["🦞"]
```

Your application should output logs that look something like this:

```
D0524 13:37:03.295982 emojis.Searcher d4ad11e1 searcher.go:53] Search query="sushi"
D0524 13:37:04.148701 emojis.Searcher d4ad11e1 searcher.go:53] Search query="curry"
D0524 13:37:04.937959 emojis.Searcher d4ad11e1 searcher.go:53] Search query="shrimp"
D0524 13:37:06.577844 emojis.Searcher d4ad11e1 searcher.go:53] Search query="lobster"
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
