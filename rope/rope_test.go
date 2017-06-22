package rope_test

import (
	"bytes"
	"testing"

	"github.com/emersion/go-tryalgo/rope"
)

func ropeBytes(r *rope.Rope) []byte {
	var buf bytes.Buffer
	r.WriteTo(&buf)
	return buf.Bytes()
}

func TestRope_leaf(t *testing.T) {
	r := rope.New([]byte("abc"))

	if l := r.Len(); l != 3 {
		t.Errorf("r.Len() = %v, want %v", l, 3)
	}
	if b := r.Get(1); b != byte('b') {
		t.Errorf("r.Get(1) = %v, want %v", b, byte('b'))
	}
	if s := string(ropeBytes(r)); s != "abc" {
		t.Errorf("r.WriteTo(bytes.Buffer) = %q, want %q", s, "abc")
	}
	if s := string(ropeBytes(r.Slice(1, 3))); s != "bc" {
		t.Errorf("r.Slice(1, 2) = %q, want %q", s, "bc")
	}
}

func TestRope_tree(t *testing.T) {
	leaf1 := rope.New([]byte("abc"))
	leaf2 := rope.New([]byte("def"))
	leaf3 := rope.New([]byte("ghi"))

	r2 := rope.Merge(leaf2, leaf3)
	r := rope.Merge(leaf1, r2)

	if l := r.Len(); l != 9 {
		t.Errorf("r.Len() = %v, want %v", l, 9)
	}
	if b := r.Get(0); b != byte('a') {
		t.Errorf("r.Get(0) = %v, want %v", rune(b), 'a')
	}
	if b := r.Get(4); b != byte('e') {
		t.Errorf("r.Get(4) = %v, want %v", rune(b), 'e')
	}
	if b := r.Get(6); b != byte('g') {
		t.Errorf("r.Get(6) = %v, want %v", rune(b), 'g')
	}
	if s := string(ropeBytes(r)); s != "abcdefghi" {
		t.Errorf("r.WriteTo(bytes.Buffer) = %q, want %q", s, "abcdefghi")
	}
	if s := string(ropeBytes(r.Slice(1, 3))); s != "bc" {
		t.Errorf("r.Slice(1, 3) = %q, want %q", s, "bc")
	}
	if s := string(ropeBytes(r.Slice(1, 7))); s != "bcdefg" {
		t.Errorf("r.Slice(1, 7) = %q, want %q", s, "bcdefg")
	}
}
