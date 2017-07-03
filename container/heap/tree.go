package heap

// LessFunc compares a and b and returns true if and only if a < b.
type LessFunc func(a, b interface{}) bool

type node struct {
	value       interface{}
	left, right *node
}

func merge(n1, n2 *node, less LessFunc) *node {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	if less(n1.value, n2.value) {
		n1.left, n1.right = merge(n1.right, n2, less), n1.left
		return n1
	} else {
		n2.left, n2.right = merge(n2.right, n1, less), n2.left
		return n2
	}
}

// Tree is a skew heap.
type Tree struct {
	root *node
	less LessFunc
}

// NewTree creates a new Tree.
func NewTree(less LessFunc) *Tree {
	return &Tree{less: less}
}

// Min returns the lowest element in the heap.
func (t *Tree) Min() (interface{}, bool) {
	if t.root == nil {
		return nil, false
	}
	return t.root.value, true
}

// Push adds a new element to the heap.
func (t *Tree) Push(v interface{}) {
	t.root = merge(t.root, &node{value: v}, t.less)
}

// Pop removes the lowest element from the heap.
func (t *Tree) Pop() (interface{}, bool) {
	v, ok := t.Min()
	if ok {
		t.root = merge(t.root.left, t.root.right, t.less)
	}
	return v, ok
}
