package unionfind

type Node struct {
	Value interface{}
	parent *Node
	rank int
}

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

func (uf *Node) Find() interface{} {
	return uf.find().Value
}

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
