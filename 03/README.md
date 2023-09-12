# Unit Tests

In this part, you will write a unit test for the `Searcher` component. Review
[the documentation on testing][testing]. Then, in a file called
`searcher_test.go`, write a couple unit tests for the `Search` method.

- The query `"pig"` should return 🐖, 🐗, 🐷, and 🐽.
- The query `"PiG"` should return 🐖, 🐗, 🐷, and 🐽.
- The query `"black cat"` should return 🐈‍⬛.
- The query `"foo bar baz"` should return `[]string{}`.

Make sure to run every unit test with all the `Runner`s returned by
[`weavertest.AllRunners`][AllRunners].

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/912c215cecd611feadd2e23fcc80fe09f4af2045/03/searcher_test.go#L15-L51
</details>

Finally, run the unit tests with `go test`.

```
$ go test .
PASS
ok	emojis	0.491s
```

[**:arrow_left: Previous Part**](../02)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../04)

[testing]: https://serviceweaver.dev/docs.html#testing
[AllRunners]: https://pkg.go.dev/github.com/ServiceWeaver/weaver/weavertest#AllRunners
