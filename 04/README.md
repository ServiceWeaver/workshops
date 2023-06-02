# An HTTP Server

In this part, you will implement an HTTP server as a frontend to your
application. Review [the documentation on listeners][listeners]. First, update
the `Main` method in `main.go` to get a listener called `"emojis"`. Use a
[`weaver.ListenerOptions`][ListenerOptions] with `LocalAddress` set to
`"localhost:9000"`. We recommend you print out the listener to remind yourself
the address it is listening on.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/e9c0573de0f20fca6a88106ad9f25fddf2f04233/04/main.go#L44-L49
</details>

Next, inside of the `Main` method, implement an HTTP handler for the `/search`
endpoint. The endpoint receives GET requests of the form `/search?q=<query>` and
returns the JSON serialized list of emojis that match the query
([`json.Marshal`](https://pkg.go.dev/encoding/json#Marshal)). For example, `curl
localhost:9000/search?q=pig` should return `["🐖","🐷","🐽"]`. At the end of the
`Main` method, call [`http.Serve`][http.Serve] to serve HTTP traffic using the
handler you just implemented.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/e9c0573de0f20fca6a88106ad9f25fddf2f04233/04/main.go#L58-L74
</details>

Finally, run your application. It should block serving traffic.

```
$ go run .
```

In a separate terminal, curl the `/search` endpoint:

```
$ curl localhost:9000/search?q=pig
["🐖","🐷","🐽"]
$ curl localhost:9000/search?q=cow
["🐄","🐮"]
$ curl localhost:9000/search?baby%bird
["🐤","🐥"]
```

If you do not have `curl` installed on your machine or if your terminal does not
render emojis well, you can instead use a web browser. If you visit
`localhost:9000/search?q=pig`, for example, you should see a page with
`["🐖","🐷","🐽"]` as its contents.

While your application is running, run `weaver single status` to see information
about the application.

```
$ weaver single status
╭─────────────────────────────────────────────────────╮
│ DEPLOYMENTS                                         │
├────────┬──────────────────────────────────────┬─────┤
│ APP    │ DEPLOYMENT                           │ AGE │
├────────┼──────────────────────────────────────┼─────┤
│ emojis │ 370e6a90-22ba-4b1f-8558-86d90b516a54 │ 9s  │
╰────────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────╮
│ COMPONENTS                                           │
├────────┬────────────┬─────────────────┬──────────────┤
│ APP    │ DEPLOYMENT │ COMPONENT       │ REPLICA PIDS │
├────────┼────────────┼─────────────────┼──────────────┤
│ emojis │ 370e6a90   │ emojis.Searcher │ 3493842      │
│ emojis │ 370e6a90   │ main            │ 3493842      │
╰────────┴────────────┴─────────────────┴──────────────╯
╭──────────────────────────────────────────────────╮
│ LISTENERS                                        │
├────────┬────────────┬──────────┬─────────────────┤
│ APP    │ DEPLOYMENT │ LISTENER │ ADDRESS         │
├────────┼────────────┼──────────┼─────────────────┤
│ emojis │ 370e6a90   │ emojis   │ 127.0.0.1:12347 │
╰────────┴────────────┴──────────┴─────────────────╯
```

When you `go run .` a Service Weaver application, every component is co-located
in the same OS process. We can infer this from the output of `weaver single
status` because the two components, `main` and `emojis.Searcher`, have the same
process id.

## (Optional) Web UI

We have written a web UI, [`index.html`](index.html), for your app. You can
optionally serve `index.html` on the root endpoint `/` of your HTTP server.
First, import the [embed](embed) package into `main.go`.

```
import _ "embed"
```

Next, download or copy-and-paste the [`index.html`](index.html) file into your
`emojis/` directory. Next, embed the contents of `index.html` into a
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

Re-run your application and open `localhost:9000` in a browser. You should see
the following web UI.

[emoji_search_demo.webm](https://github.com/ServiceWeaver/workshops/assets/3654277/8ced2cb0-18c2-41fc-b14f-cc4f602deb38)

[**:arrow_left: Previous Part**](../03)
&nbsp;&nbsp;&nbsp;:black_circle:&nbsp;&nbsp;&nbsp;
[**Next Part :arrow_right:**](../05)

[listeners]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-listeners
[ListenerOptions]: https://pkg.go.dev/github.com/ServiceWeaver/weaver#ListenerOptions
[http.Serve]: https://pkg.go.dev/net/http#Serve
[embed]: https://pkg.go.dev/embed
