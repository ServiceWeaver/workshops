# Hello, World!

In this part, you will install Service Weaver and write a simple "Hello, World!"
application. Begin by following [the installation instructions][installation] on
serviceweaver.dev to install Go version 1.20 or higher and install the `weaver`
command line tool. If done correctly, you should be able to run `weaver --help`:

```
$ weaver --help
USAGE

  weaver generate                 // weaver code generator
  weaver single    <command> ...  // for single process deployments
  weaver multi     <command> ...  // for multiprocess deployments
  ...
```

Next, create a directory called `primes`, `cd` into it, and initialize a go
module called `primes`:

```
$ mkdir primes/
$ cd primes/
$ go mod init primes
```

Review [the documentation on components][components]. Then, in a file called
`main.go`, write a simple Service Weaver application that prints "Hello, World!"
and terminates. When you call [`weaver.Run`][weaver_Run], pass it a function
called `serve`. Feel free to copy and paste code from the documentation.

<details>
<summary>Solution.</summary>

TODO(mwhittaker): Embed solution here.
</details>

Finally, build and run your application:

```
$ go mod tidy
$ weaver generate .
$ go run .
Hello, World!
```

[**Next Part :arrow_right:**](../02)

[installation]: https://serviceweaver.dev/docs.html#installation
[components]: https://serviceweaver.dev/docs.html#step-by-step-tutorial-components
[weaver_Run]: https://pkg.go.dev/github.com/ServiceWeaver/weaver#Run
