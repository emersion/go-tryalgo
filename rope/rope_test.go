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
		t.Errorf("r.Len() = %v, want 3", l)
	}
	if b := r.Get(1); b != byte('b') {
		t.Errorf("r.Get(1) = %v, want %v", b, byte('b'))
	}
	if s := string(ropeBytes(r.Slice(1, 3))); s != "bc" {
		t.Errorf("r.Slice(1, 2) = %q, want %q", s, "bc")
	}
	if s := string(ropeBytes(r)); s != "abc" {
		t.Errorf("r.WriteTo(bytes.Buffer) = %q, want %q", s, "abc")
	}
}
