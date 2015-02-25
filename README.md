# offheap

Allocate system memory as if it were a []byte. 

Useful for allocating massive buffers off of Go's heap, which would otherwise consume a lot of resources.

### Install
     go get github.com/shoenig/offheap

### Example

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



