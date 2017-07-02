// Package trie implements a radix/prefix tree.
package trie

// A Node contains byte words.
type Node struct {
	leaf     bool
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

func (n *Node) closest(word []byte, dist int, closest []byte) []byte {
	if len(word) == 0 {
		if n.leaf {
			return closest
		} else {
			return nil
		}
	}

	// Perfect match
	if child, ok := n.children[word[0]]; ok {
		if closest := child.closest(word[1:], dist, append(closest, word[0])); closest != nil {
			return closest
		}
	}

	// Cannot mutate word anymore
	if dist == 0 {
		return nil
	}

	for b, child := range n.children {
		// Insert
		if closest := child.closest(word, dist-1, append(closest, b)); closest != nil {
			return closest
		}

		// Change
		if closest := child.closest(word[1:], dist-1, append(closest, b)); closest != nil {
			return closest
		}
	}

	// Remove
	return n.closest(word[1:], dist-1, closest)
}

// Closest returns the closest word stored in this node with a Levenstein
// distance of at most dist.
func (n *Node) Closest(word []byte, dist int) []byte {
	return n.closest(word, dist, nil)
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

// Remove removes a word from this node.
//
// Complexity: O(len(word))
func (n *Node) Remove(word []byte) {
	cur := n
	for _, b := range word {
		cur = cur.children[b]
		if cur == nil {
			return
		}
	}
	cur.leaf = false
	// TODO: cleanup empty branches
}
