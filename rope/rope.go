// Package rope implements efficient []byte slicing and merging operations.
package rope

import (
	"io"
)

// Rope is a []byte representation.
type Rope struct {
	len int
	b []byte // for leafs
	left, right *Rope // for internal nodes
}

// New creates a new rope containing b.
func New(b []byte) *Rope {
	return &Rope{len: len(b), b: b}
}

// Merge creates a new rope which is the merge of left and right.
func Merge(left, right *Rope) *Rope {
	return &Rope{len: left.len + right.len, left: left, right: right}
}

func (r *Rope) isLeaf() bool {
	return r.left == nil && r.right == nil
}

// Len returns the rope length.
func (r *Rope) Len() int {
	return r.len
}

// Get returns the i-th byte in the rope. It panics i is out of range.
func (r *Rope) Get(i int) byte {
	if i >= r.len {
		panic("index out of range")
	}

	cur := r
	for !cur.isLeaf() {
		if i < cur.left.len {
			cur = cur.left
		} else {
			i -= cur.left.len
			cur = cur.right
		}
	}

	return cur.b[i]
}

// Slice returns a new rope which contains bytes from low to high.
func (r *Rope) Slice(low, high int) *Rope {
	if r.isLeaf() {
		return New(r.b[low:high])
	}

	if high <= r.left.len {
		return r.left.Slice(low, high)
	}
	if low > r.left.len {
		return r.right.Slice(low - r.left.len, high)
	}

	left := r.left.Slice(low, r.left.len)
	right := r.right.Slice(0, high - r.left.len)
	return Merge(left, right)
}

// WriteTo writes this rope's content to w.
func (r *Rope) WriteTo(w io.Writer) (n int64, err error) {
	if r.isLeaf() {
		n, err := w.Write(r.b)
		return int64(n), err
	}

	nl, err := r.left.WriteTo(w)
	if err != nil {
		return nl, err
	}
	nr, err := r.right.WriteTo(w)
	return nl + nr, err
}
