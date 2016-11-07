# Golang Learnings

This repo is for learning Go.

## Download
Download binaries from [here](https://golang.org/dl/).

## Compile
* golang-learning/
  * hello/
    * main.go
    * hello


* `cd golang-learning/hello/`
* `go build`

`go build` will create the executable `hello` named after the project (directory name).  To run, all you need to do is type: `./hello`

## Compile _and_ Run
* `cd src/hello/`
* `go run main.go`

Note: This does _not_ create an executable!

## Format Code
Automatically format source code with:
* `gofmt main.go`

OR, automatically format source code _and_ write to the original file source instead of writing to the console:
* `gofmt -w main.go`
