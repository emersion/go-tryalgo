// Package trie implements a radix/prefix tree.
package trie

// A Node contains byte words.
type Node struct {
	leaf bool
	children map[byte]*Node
}

// New creates a new node.
func New() *Node {
	return new(Node)
}

// Contains checks if a word is in this node.
//
// Complexity: O(len(word))
func (n *Node) Contains(word []byte) bool {
	cur := n
	for _, b := range word {
		cur = cur.children[b]
		if cur == nil {
			return false
		}
	}
	return cur.leaf
}

// Append adds a word to this node.
//
// Complexity: O(len(word))
func (n *Node) Append(word []byte) {
	cur := n
	for _, b := range word {
		next := cur.children[b]
		if next == nil {
			if cur.children == nil {
				cur.children = make(map[byte]*Node)
			}
			next = new(Node)
			cur.children[b] = next
		}
		cur = next
	}
	cur.leaf = true
}
