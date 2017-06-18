package unionfind_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/unionfind"
)

func TestInts(t *testing.T) {
	ints := unionfind.NewInts(4)

	intersect := func (i, j int) bool {
		return ints.Find(i) == ints.Find(j)
	}

	if intersect(0, 1) {
		t.Errorf("intersect(0, 1) = true after init, expected false")
	}

	ints.Union(0, 1)
	if !intersect(0, 1) {
		t.Errorf("intersect(0, 1) = false after Union(0, 1), expected true")
	}
	if intersect(0, 2) {
		t.Errorf("intersect(0, 2) = true after Union(0, 1), expected false")
	}
	if intersect(1, 2) {
		t.Errorf("intersect(1, 2) = true after Union(0, 1), expected false")
	}

	ints.Union(2, 3)
	ints.Union(1, 3)
	if !intersect(0, 2) {
		t.Errorf("intersect(0, 2) = false, expected true")
	}
	if !intersect(0, 3) {
		t.Errorf("intersect(0, 3) = false, expected true")
	}
	if !intersect(2, 1) {
		t.Errorf("intersect(2, 1) = false, expected true")
	}

	ints.Union(0, 1)
	if !intersect(0, 1) {
		t.Errorf("intersect(0, 1) = false after Union(0, 1), expected true")
	}
}
