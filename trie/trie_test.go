package trie_test

import (
	"bytes"
	"testing"

	"github.com/emersion/go-tryalgo/trie"
)

func TestNode_Contains_empty(t *testing.T) {
	n := trie.New()

	if n.Contains([]byte("serial")) {
		t.Errorf("n.Contains(serial) = true when empty, want false")
	}
}

func newPopulated() *trie.Node {
	n := trie.New()
	n.Append([]byte("serial"))
	n.Append([]byte("serial experiments"))
	n.Append([]byte("serial experiments lain"))
	n.Append([]byte("experiments lain"))
	n.Append([]byte("lain"))
	return n
}

func TestNode_Contains(t *testing.T) {
	n := newPopulated()

	if !n.Contains([]byte("serial")) {
		t.Errorf("n.Contains(serial) = false, want true")
	}
	if !n.Contains([]byte("serial experiments lain")) {
		t.Errorf("n.Contains(serial experiments lain) = false, want true")
	}
	if n.Contains([]byte("serial exp")) {
		t.Errorf("n.Contains(serial exp) = true, want false")
	}
}

func TestNode_Remove(t *testing.T) {
	n := newPopulated()

	n.Remove([]byte("serial"))
	if n.Contains([]byte("serial")) {
		t.Errorf("n.Contains(serial) = true, want false")
	}
	if !n.Contains([]byte("serial experiments lain")) {
		t.Errorf("n.Contains(serial experiments lain) = false, want true")
	}
}

func TestNode_Closest(t *testing.T) {
	n := newPopulated()
	dist := 1

	tests := []struct{
		word, want []byte
	}{
		{[]byte("serial"), []byte("serial")},
		{[]byte("seriaal"), []byte("serial")},
		{[]byte("seril"), []byte("serial")},
		{[]byte("seriel"), []byte("serial")},
		{[]byte("seriaaal"), nil},
	}

	for _, test := range tests {
		if got := n.Closest(test.word, dist); !bytes.Equal(got, test.want) {
			t.Errorf("n.Closest(%v) = %v, want %v", test.word, got, test.want)
		}
	}
}
