# A Prime Factorization Component

In this part, you will implement a simple component that computes prime
factorizations. The [fundamental theorem of arithmetic][fundamental_theorem]
states that every integer can be uniquely represented as the product of its
prime factors. For example, the number `980` has prime factorization
`2 * 2 * 5 * 7 * 7`. The number `97` has prime factorization `97` because `97`
is prime.

Review [the documentation on writing components][components]. Then, in a file
called `factorer.go`, write a component called `Factorer` with a
`Factor(ctx context.Context, x int) ([]int, error)` method that returns the
prime factorization of the provided integer. For example, calling `Factor` on
`980` should return `[]int{2, 2, 5, 7, 7}`. Return an error if the provided
integer is less than 1. Feel free to implement a naive algorithm like [trial
division][trial_division]:

```go
var factors []int
for factor := 2; factor <= x; {
    if x%factor == 0 {
        factors = append(factors, factor)
        x /= factor
    } else {
        factor++
    }
}
```

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Next, update your application to print out the factors of `980` instead of
"Hello, World!".

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Finally, run your application.

```
$ weaver generate .
$ go run .
[2 2 5 7 7]
```

Note that you'll have to run `weaver generate` whenever you add a component,
remove a component, or change the interface of a component. If your application
ever fails to build with an error coming from a `weaver_gen.go` file, try
re-running `weaver generate`.

[**:arrow_left: Previous Part**](../01)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../03)

[components]: https://serviceweaver.dev/docs.html#components
[fundamental_theorem]: https://en.wikipedia.org/wiki/Fundamental_theorem_of_arithmetic
[trial_division]: https://en.wikipedia.org/wiki/Trial_division
[writing_components]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-multiple-components
