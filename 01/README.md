# Hello, World!

In this part, you will install Service Weaver and write a simple "Hello, World!"
application. Begin by following [the installation instructions][installation] on
serviceweaver.dev to install Go version 1.20 or higher and install the `weaver`
command-line tool. If done correctly, you should be able to run `weaver --help`:

```
$ weaver --help
USAGE

  weaver generate                 // weaver code generator
  weaver version                  // show weaver version
  weaver single    <command> ...  // for single process deployments
  weaver multi     <command> ...  // for multiprocess deployments
  ...
```

Next, create a directory called `emojis`, `cd` into it, and initialize a go
module called `emojis`:

```
$ mkdir emojis/
$ cd emojis/
$ go mod init emojis
```

Review [the documentation on components][components]. Then, in a file called
`main.go`, write a simple Service Weaver application that prints "Hello, World!"
and terminates. Feel free to copy and paste code from the documentation.

<details>
<summary>Solution.</summary>

https://github.com/ServiceWeaver/workshops/blob/e9c0573de0f20fca6a88106ad9f25fddf2f04233/01/main.go#L15-L39
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
