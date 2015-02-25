package offheap

import (
	"testing"
)

func Test_New(t *testing.T) {
	m, e := New(5 * 1024 * 1024)
	if e != nil {
		t.Fatal("New failed", e)
	}
	if len(m) != 5242880 {
		t.Fatal("5 MB chunk wrong size")
	}
}

func Test_Unmap(t *testing.T) {
	m, e := New(5 * 1024 * 1025)
	if e != nil {
		t.Fatal("New failed", e)
	}
	if v1 := m[0]; v1 != 0 {
		t.Fatal("first bytes not zero")
	}
	if e := m.Unmap(); e != nil {
		t.Fatal("unmap failed", e)
	}
}
