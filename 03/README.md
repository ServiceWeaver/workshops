# Unit Tests

In this part, you will write a unit test for the `Factorer` component. Review
[the documentation on testing][testing]. Then, in a file called
`factorer_test.go`, write a couple unit tests for the `Factor` method. We
recommend you test `Factor` on `-1`, `0`, `1`, `2`, and `980`. Make sure to run
every unit test with all the `Runner`s returned by
[`weavertest.AllRunners`][AllRunners].

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Finally, run the unit tests with `go test`.

```
$ go test .
PASS
ok	primes	0.491s
```

[**:arrow_left: Previous Part**](../02)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../04)

[testing]: https://serviceweaver.dev/docs.html#testing
[AllRunners]: https://pkg.go.dev/github.com/ServiceWeaver/weaver@v0.12.0/weavertest#AllRunners
