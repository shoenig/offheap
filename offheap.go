// Copyright 2015 Seth Hoenig. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package offheap provides a way to use system memory directly.
package offheap

import (
	"github.com/edsrzf/mmap-go"
)

// Memory represents a private anonymous mmap in-memory-only
// file that your program can use to read and write bytes.
// This allows your program to use very large amounts of memory
// without the Go runtime trying to garbage collect it.
type Memory []byte

// New creates an offheap Memory slice of the specified number of bytes.
func New(bytes int64) (Memory, error) {
	mapped, e := mmap.MapRegion(
		nil,
		int(bytes),
		mmap.COPY,
		mmap.ANON,
		0,
	)

	if e != nil {
		return nil, e
	}
	return Memory(mapped), nil
}

// Unmap will delete the offheap Memory slice. Any usage of the Memory
// afterwords will cause a panic.
func (m Memory) Unmap() error {
	p := mmap.MMap(m)
	return p.Unmap()
}
