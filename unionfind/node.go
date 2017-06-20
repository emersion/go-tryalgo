package unionfind

// A Node is a set of values.
type Node struct {
	Value interface{}
	parent *Node
	rank int
}

// NewNode creates a new set of values.
func NewNode(v interface{}) *Node {
	return &Node{Value: v}
}

func (uf *Node) find() *Node {
	node := uf
	for node.parent != nil {
		node = node.parent
	}

	// Path compression
	if uf.parent != nil && uf.parent != node {
		uf.parent = node
	}

	return node
}

// Find returns the representative member of this node's value. Two values v1
// and v2 are in the same set if and only if Find(v1) == Find(v2).
//
// Complexity: O(α(n))
func (uf *Node) Find() interface{} {
	return uf.find().Value
}

// Union merges the sets where this node's value and other's value are.
//
// Complexity: O(α(n))
func (uf *Node) Union(other *Node) {
	root := uf.find()
	otherRoot := other.find()
	if root == otherRoot {
		return
	}

	if root.rank == otherRoot.rank {
		root.rank++
	}

	newRoot, newChild := root, otherRoot
	if root.rank < otherRoot.rank {
		newRoot, newChild = otherRoot, root
	}
	newChild.parent = newRoot
}
