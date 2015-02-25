package offheap

import (
	"launchpad.net/gommap"
)

// Memory represents a private anonymous mmap in-memory-only
// file that your program can use to read and write bytes
// without dealing with large heap issues.
type Memory []byte

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

func (m Memory) Unmap() error {
	return gommap.MMap(m).UnsafeUnmap()
}
