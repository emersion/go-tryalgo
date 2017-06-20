package heap_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/heap"
)

func TestInts(t *testing.T) {
	ints := heap.NewInts(nil)

	if v, ok := ints.Min(); ok {
		t.Errorf("ints.Min() = %v, true after init; want _, false", v)
	}

	ints.Push(42)
	ints.Push(5)
	ints.Push(6)
	ints.Push(10)
	ints.Push(4)
	ints.Push(32)
	if v, ok := ints.Min(); !ok || v != 4 {
		t.Errorf("ints.Min() = %v, %v after Push; want 4, true", v, ok)
	}

	ints.Pop()
	if v, ok := ints.Min(); !ok || v != 5 {
		t.Errorf("ints.Min() = %v, %v after Pop; want 5, true", v, ok)
	}

	ints.Pop()
	ints.Pop()
	ints.Pop()
	ints.Pop()
	ints.Pop()
	if v, ok := ints.Min(); ok {
		t.Errorf("ints.Min() = %v, true after removing all elements; want _, false", v)
	}
}
