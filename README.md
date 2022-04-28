offheap
=======

Allocate offheap memory in Go programs

[![Go Report Card](https://goreportcard.com/badge/gophers.dev/pkgs/offheap)](https://goreportcard.com/report/gophers.dev/pkgs/offheap)
[![run-ci](https://github.com/shoenig/offheap/actions/workflows/ci.yml/badge.svg)](https://github.com/shoenig/offheap/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/gophers.dev/pkgs/offheap?status.svg)](https://godoc.org/gophers.dev/pkgs/offheap)
![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/offheap.svg)
![GitHub](https://img.shields.io/github/license/shoenig/offheap.svg)

# Project Overview

Package `offheap` provides a simple library for allocating offheap memory in
Go programs. Doing so is useful for allocating large `[]byte` that should not
be managed by the garbage collector.

# Getting Started

The `offheap` package can be installed by running

```bash
go get gophers.dev/pkgs/offheap
 ```

#### Example Usage

```go
package main

import (
    "fmt"

    "github/shoenig/offheap"
)

func main() {
    var m offheap.Memory
    var e error

    // allocate 5 MB
    if m, e = offheap.New(5 * 1024 * 1024); e != nil {
        panic(e)
    }
    fmt.Println("m[0]", m[0])

    m[0] = 42
    fmt.Println("m[0]", m[0])

    if e = m.Unmap(); e != nil {
        panic(e)
    }
    fmt.Println("finished")
}
```

# Contributing

Package `offheap` is pretty minimalist, and so it's basically feature complete.
Bug fixes and good ideas though are always welcome, please just file an issue.

# License

The `gophers.dev/pkgs/offheap` module is open source under the [BSD-3-Clause](LICENSE) license.


