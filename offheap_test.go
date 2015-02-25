package offheap

import (
	"testing"
	"time"
)

func Test_New(t *testing.T) {
	m, e := New(1024 * 1024 * 1024)

	println("len: ", len(m))

	if e != nil {
		t.Fatal("e not nil", e)
	}

	time.Sleep(10 * time.Minute)



	t.Log("m", m)
}
