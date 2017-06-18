package unionfind_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/unionfind"
)

func TestNode(t *testing.T) {
	one := unionfind.NewNode(1)
	two := unionfind.NewNode(2)
	three := unionfind.NewNode(3)
	four := unionfind.NewNode(4)

	intersect := func (a, b *unionfind.Node) bool {
		return a.Find().(int) == b.Find().(int)
	}

	if intersect(one, two) {
		t.Errorf("intersect(one, two) = true after init, expected false")
	}

	one.Union(two)
	t.Log(one, two)
	if !intersect(one, two) {
		t.Errorf("intersect(one, two) = false after one.Union(two), expected true")
	}
	if intersect(one, three) {
		t.Errorf("intersect(one, three) = true after one.Union(two), expected false")
	}
	if intersect(two, three) {
		t.Errorf("intersect(two, three) = true after one.Union(two), expected false")
	}

	three.Union(four)
	two.Union(four)
	if !intersect(one, three) {
		t.Errorf("intersect(one, three) = false, expected true")
	}
	if !intersect(one, four) {
		t.Errorf("intersect(one, four) = false, expected true")
	}
	if !intersect(three, two) {
		t.Errorf("intersect(three, two) = false, expected true")
	}
}
