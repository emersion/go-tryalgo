package trie_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/trie"
)

func TestNode(t *testing.T) {
	n := trie.New()

	if n.Contains([]byte("serial")) {
		t.Errorf("n.Contains(serial) = true after init, want false")
	}

	n.Append([]byte("serial"))
	n.Append([]byte("serial experiments"))
	n.Append([]byte("serial experiments lain"))
	n.Append([]byte("experiments lain"))
	n.Append([]byte("lain"))
	if !n.Contains([]byte("serial")) {
		t.Errorf("n.Contains(serial) = false, want true")
	}
	if !n.Contains([]byte("serial experiments lain")) {
		t.Errorf("n.Contains(serial experiments lain) = false, want true")
	}
	if n.Contains([]byte("serial exp")) {
		t.Errorf("n.Contains(serial exp) = true, want false")
	}

	n.Remove([]byte("serial"))
	if n.Contains([]byte("serial")) {
		t.Errorf("n.Contains(serial) = true, want false")
	}
	if !n.Contains([]byte("serial experiments lain")) {
		t.Errorf("n.Contains(serial experiments lain) = false, want true")
	}
}
