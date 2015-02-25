// Copyright 2015 Seth Hoenig. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package offheap provides a way to grab and use a chunk of system memory.
package offheap

import (
	"launchpad.net/gommap"
)

// Memory represents a private anonymous mmap in-memory-only
// file that your program can use to read and write bytes
// without dealing with large heap issues.
type Memory []byte

// New creates an offheap Memory slice of the specified number of bytes.
func New(bytes int64) (Memory, error) {
	mapped, e := gommap.MapRegion(
		0,
		0,
		bytes,
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_ANONYMOUS|gommap.MAP_PRIVATE,
	)

	if e != nil {
		return nil, e
	}
	return Memory(mapped), nil
}

// Unmap will delete the offheap Memory slice. Any usage of the Memory
// afterwords will cause a panic.
func (m Memory) Unmap() error {
	return gommap.MMap(m).UnsafeUnmap()
}
