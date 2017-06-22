package bintree_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/bintree"
)

func TestAVL(t *testing.T) {
	avl := bintree.NewAVL(func(a, b interface{}) bool {
		return a.(int) < b.(int)
	})

	if avl.Contains(42) {
		t.Errorf("avl.Contains(42) = true after init, want false")
	}

	avl.Append(1)
	avl.Append(2)
	avl.Append(3)
	avl.Append(4)
	avl.Append(6)
	avl.Append(10)
	avl.Append(5)
	avl.Append(9)

	if avl.Contains(0) {
		t.Errorf("avl.Contains(0) = true, want false")
	}
	if !avl.Contains(1) {
		t.Errorf("avl.Contains(1) = false, want true")
	}
	if !avl.Contains(10) {
		t.Errorf("avl.Contains(10) = false, want true")
	}
	if !avl.Contains(6) {
		t.Errorf("avl.Contains(6) = false, want true")
	}
	if !avl.Contains(5) {
		t.Errorf("avl.Contains(5) = false, want true")
	}

	avl.Remove(9)
	if avl.Contains(9) {
		t.Errorf("avl.Contains(3) = true after avl.Remove(3), want false")
	}
}
