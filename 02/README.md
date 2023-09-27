# A Search Component

In this part, you will implement a simple component that searches for emojis
that match a query. First, download or copy-and-paste [`emojis.go`](emojis.go)
into your `emojis/` directory. This file contains a `map[string][]string` called
`emojis` that maps every emoji to a list of labels. For example, the black cat
emoji 🐈‍⬛ has labels `"animal"`, `"animals"`, `"black"`, `"cat"`,
`"mammal"`, and `"nature"`.

Next, review [the documentation on writing components][writing_components].
Then, in a file called `searcher.go`, write a component called `Searcher` with
the following method:

```go
Search(ctx context.Context, query string) ([]string, error)
```

The `Search` method receives a query like `"black cat"` and returns the emojis
that match the query. To implement the `Search` method, first lowercase the
query ([`strings.ToLower`](https://pkg.go.dev/strings#ToLower)) and then split
the query into words ([`strings.Fields`](https://pkg.go.dev/strings#Fields)).
Then iterate over the `emojis` map in `emojis.go` to find all matching emojis.
We say an emoji *matches* a query if every word in the query is one of the
emoji's labels. Return the matching emojis is sorted order
([`sort.Strings`](https://pkg.go.dev/sort#Strings)).

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/4eca79ebc6bfe3ef1225c965ec729db70f175994/02/searcher.go#L15-L68
</details>

Next, update your application to print out the emojis that match the query
`"pig"`:

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/4eca79ebc6bfe3ef1225c965ec729db70f175994/02/main.go#L30-L44
</details>

Finally, run your application.

```
$ weaver generate .
$ go run .
[🐖 🐗 🐷 🐽]
```

The `"pig"` query matches the pig emoji 🐖, the boar emoji 🐗, the pig face
emoji 🐷, and the pig nose emoji 🐽.

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
