# offheap

Allocate offheap memory in Go programs.

[![Go Report Card](https://goreportcard.com/badge/github.com/shoenig/offheap)](https://goreportcard.com/report/github.com/shoenig/offheap) [![Build Status](https://travis-ci.org/shoenig/offheap.svg?branch=master)](https://travis-ci.org/shoenig/offheap) [![GoDoc](https://godoc.org/github.com/shoenig/offheap?status.svg)](https://godoc.org/github.com/shoenig/offheap) [![License](https://img.shields.io/github/license/shoenig/offheap.svg?style=flat-square)](LICENSE)

Useful for allocating massive buffers off of Go's heap, which would otherwise incur 
expensive garbage collection busywork.

### Install
     go get github.com/shoenig/offheap

### Example

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


