package offheap

import (
	"testing"

	"github.com/shoenig/test/must"
)

func Test_New(t *testing.T) {
	m, err := New(5 * 1024 * 1024)
	must.NoError(t, err)
	must.LenSlice(t, 5242880, m)
}

func Test_Unmap(t *testing.T) {
	m, err := New(5 * 1024 * 1025)
	must.NoError(t, err)
	must.EqCmp(t, 0, m[0])

	err = m.Unmap()
	must.NoError(t, err)
}
