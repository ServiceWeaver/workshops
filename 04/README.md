# An HTTP Server

In this part, you will implement an HTTP server as a frontend to your
application. Review [the documentation on listeners][listeners]. First, update
the `serve` function in `main.go` to get a listener called `"primes"`. Use a
[`weaver.ListenerOptions`][ListenerOptions] with `LocalAddress` set to
`"localhost:12347"` (`12347` is a prime :grin:). We recommend you print out the
listener to remind yourself the address it is listening on.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Next, inside of the `serve` function, implement an HTTP handler for the
`/factor` endpoint. The endpoint receives GET requests of the form
`/factor?x=<number>` and returns the JSON serialized prime factorization of `x`.
For example, `curl localhost:12347/factor?x=980` should return `[2,2,5,7,7]`. At
the end of the `serve` function, call [`http.Serve`][http.Serve] to serve HTTP
traffic using the handler you just implemented.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Finally, run your application. It should block serving traffic.

```
$ go run .
```

In a separate terminal, curl the `/factor` endpoint:

```
$ curl localhost:12347/factor?x=980
[2,2,5,7,7]
$ curl localhost:12347/factor?x=12347
[12347]
$ curl localhost:12347/factor?x=-1
non-positive x: -1
```

If you do not have `curl` installed on your machine, you can instead use a web
browser. If you visit `localhost:12347/factor?x=980`, for example, you should
see a page with `[2,2,5,7,7]` as its contents.

While your application is running, run `weaver single status` to see information
about the application.

```
$ weaver single status
╭─────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                         │
├────────┬──────────────────────────────────────┬─────┤
│ APP    │ DEPLOYMENT                           │ AGE │
├────────┼──────────────────────────────────────┼─────┤
│ primes │ 370e6a90-22ba-4b1f-8558-86d90b516a54 │ 9s  │
╰────────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────╮
│ COMPONENTS                                           │
├────────┬────────────┬─────────────────┬──────────────┤
│ APP    │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS │
├────────┼────────────┼─────────────────┼──────────────┤
│ primes │ 370e6a90   │ primes.Factorer │ 3493842      │
│ primes │ 370e6a90   │ main            │ 3493842      │
╰────────┴────────────┴─────────────────┴──────────────╯
╭──────────────────────────────────────────────────╮
│ LISTENERS                                        │
├────────┬────────────┬──────────┬─────────────────┤
│ APP    │ DEPLOYMENT │ LISTENER │ ADDRESS         │
├────────┼────────────┼──────────┼─────────────────┤
│ primes │ 370e6a90   │ primes   │ 127.0.0.1:12347 │
╰────────┴────────────┴──────────┴─────────────────╯
```

When you `go run .` a Service Weaver application, every component is co-located
in the same OS process. We can infer this from the output of `weaver single
status` because the two components, `main` and `primes.Factorer`, have the same
process id.

## (Optional) Web UI

We have written a web UI, [`index.html`](index.html), for the prime
factorization app. You can optionally serve `index.html` on the root endpoint
`/` of your HTTP server. First, import the [embed](embed) package into `main.go`.

```
import _ "embed"
```

Next, download or copy-and-paste the [`index.html`](index.html) file into your
`primes/` directory. Next, embed the contents of `index.html` into a
package-scoped `indexHtml` variable in `main.go`.

```
//go:embed index.html
indexHtml string
```

Finally, register the following HTTP handler.

```
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, indexHtml)
})
```

Re-run your application and open `localhost:12347` in a browser. You should see
the following web UI.

TODO(mwhittaker): Insert a video demo of the frontend.

[**:arrow_left: Previous Part**](../03)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../05)

[listeners]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-listeners
[ListenerOptions]: https://pkg.go.dev/github.com/ServiceWeaver/weaver#ListenerOptions
[http.Serve]: https://pkg.go.dev/net/http#Serve
[embed]: https://pkg.go.dev/embed
