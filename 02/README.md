# A Search Component

In this part, you will implement a simple component that searches for emojis
that match a query. First, download or copy-and-paste [`emojis.go`](emojis.go)
into your `emojis/` directory. This file contains a `map[string][]string` called
`emojis` that maps every emoji to a list of labels. For example, the black cat
emoji 🐈‍⬛ has labels `"animal"`, `"animals"`, `"black"`, `"cat"`,
`"mammal"`, and `"nature"`. Refer to https://github.com/mwhittaker/emojis if you
want to understand how these labels were selected.

Next, review [the documentation on writing components][writing_components].
Then, in a file called `searcher.go`, write a component called `Searcher` with
the following method:

```go
Search(context context.Context, query string) ([]string, error)
```

The `Search` method receives a query like `"black cat"` and returns the emojis
that match the query. To implement the `Search` method, first lowercase the
query ([`strings.ToLower`](https://pkg.go.dev/strings#ToLower)) and then split
the query into words ([`strings.Fields`](https://pkg.go.dev/strings#Fields)).
Then iterate over the `emojis` map in `emojis.go` to find all matching emojis.
We say an emoji *matches* a query if every word in the query is a substring of
at least one of the emoji's labels. Return the matching emojis is sorted order
([`sort.Strings`](https://pkg.go.dev/sort#Strings)).

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Next, update your application to print out the emojis that match the query
`"rock"`:

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Finally, run your application.

```
$ weaver generate .
$ go run .
[☘️ 🚀 🪨]
```

The `"rock"` query matches the sham*rock* emoji ☘️, the *rock*et emoji 🚀, and
the *rock* emoji 🪨.

Note that you'll have to run `weaver generate` whenever you add a component,
remove a component, or change the interface of a component. If your application
ever fails to build with an error coming from a `weaver_gen.go` file, try
re-running `weaver generate`.

[**:arrow_left: Previous Part**](../01)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../03)

[fundamental_theorem]: https://en.wikipedia.org/wiki/Fundamental_theorem_of_arithmetic
[trial_division]: https://en.wikipedia.org/wiki/Trial_division
[writing_components]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-multiple-components
