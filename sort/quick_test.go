package sort_test

import (
	"reflect"
	"testing"

	"github.com/emersion/go-tryalgo/sort"
)

func TestQuick(t *testing.T) {
	a := newTestArray()
	sort.Quick(a)

	if !reflect.DeepEqual(a, testSorted) {
		t.Errorf("sort.Quick() = %v, want %v", a, testSorted)
	}
}
