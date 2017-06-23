package sort_test

import (
	"reflect"
	"testing"

	"github.com/emersion/go-tryalgo/sort"
)

func TestInsertion(t *testing.T) {
	a := newTestArray()
	sort.Insertion(a)

	if !reflect.DeepEqual(a, testSorted) {
		t.Errorf("sort.Insertion() = %v, want %v", a, testSorted)
	}
}
