package heap_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/heap"
)

func TestTree(t *testing.T) {
	tree := heap.NewTree(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	if v, ok := tree.Min(); ok {
		t.Errorf("tree.Min() = %v, true after init; want _, false", v)
	}

	tree.Push(42)
	tree.Push(5)
	tree.Push(6)
	tree.Push(10)
	tree.Push(4)
	tree.Push(32)
	if v, ok := tree.Min(); !ok || v.(int) != 4 {
		t.Errorf("tree.Min() = %v, %v after Push; want 4, true", v, ok)
	}

	tree.Pop()
	if v, ok := tree.Min(); !ok || v.(int) != 5 {
		t.Errorf("tree.Min() = %v, %v after Pop; want 5, true", v, ok)
	}

	tree.Pop()
	tree.Pop()
	tree.Pop()
	tree.Pop()
	tree.Pop()
	if v, ok := tree.Min(); ok {
		t.Errorf("tree.Min() = %v, true after removing all elements; want _, false", v)
	}
}
