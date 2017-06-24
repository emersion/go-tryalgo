package sort_test

import (
	"reflect"
	"testing"

	"github.com/emersion/go-tryalgo/sort"
)

func TestMerge(t *testing.T) {
	a := newTestArray()
	a = sort.Merge(a)

	if !reflect.DeepEqual(a, testSorted) {
		t.Errorf("sort.Merge() = %v, want %v", a, testSorted)
	}
}
