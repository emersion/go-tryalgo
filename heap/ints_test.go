package heap_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/heap"
)

func TestInts(t *testing.T) {
	ints := heap.NewInts()

	if v, ok := ints.Min(); ok {
		t.Errorf("ints.Min() = %v, true after init; want _, false", v)
	}

	ints.Append(42)
	ints.Append(5)
	ints.Append(6)
	ints.Append(10)
	ints.Append(4)
	ints.Append(32)
	if v, ok := ints.Min(); !ok || v != 4 {
		t.Errorf("ints.Min() = %v, %v after Append; want 4, true", v, ok)
	}

	ints.RemoveMin()
	if v, ok := ints.Min(); !ok || v != 5 {
		t.Errorf("ints.Min() = %v, %v after RemoveMin; want 5, true", v, ok)
	}

	ints.RemoveMin()
	ints.RemoveMin()
	ints.RemoveMin()
	ints.RemoveMin()
	ints.RemoveMin()
	if v, ok := ints.Min(); ok {
		t.Errorf("ints.Min() = %v, true after removing all elements; want _, false", v)
	}
}
