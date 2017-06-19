package trie

type Node struct {
	leaf bool
	children map[byte]*Node
}

func New() *Node {
	return new(Node)
}

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
