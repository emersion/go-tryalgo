package sort_test

import (
	"reflect"
	"testing"

	"github.com/emersion/go-tryalgo/sort"
)

func TestHeap(t *testing.T) {
	a := newTestArray()
	sort.Heap(a)

	if !reflect.DeepEqual(a, testSorted) {
		t.Errorf("sort.Heap() = %v, want %v", a, testSorted)
	}
}
