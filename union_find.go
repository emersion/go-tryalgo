package tryalgo

type UnionFind struct {
	Value interface{}
	parent *UnionFind
	rank int
}

func NewUnionFind(v interface{}) *UnionFind {
	return &UnionFind{Value: v}
}

func (uf *UnionFind) find() *UnionFind {
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

func (uf *UnionFind) Find() interface{} {
	return uf.find().Value
}

func (uf *UnionFind) Union(other *UnionFind) {
	root := uf.find()
	otherRoot := other.find()

	if root.rank == otherRoot.rank {
		root.rank++
	}

	newRoot, newChild := root, otherRoot
	if root.rank < otherRoot.rank {
		newRoot, newChild = otherRoot, root
	}
	newChild.parent = newRoot
}
