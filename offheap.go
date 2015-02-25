package offheap

import (
	// todo open source version
	"indeed/gophers/3rdparty/p/launchpad.net/gommap"
)

type Memory gommap.MMap // use like []byte

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
