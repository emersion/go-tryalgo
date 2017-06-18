package bintree

type LessFunc func(a, b interface{}) bool

type AVL struct {
	less LessFunc
	root *node
}

type node struct {
	value interface{}
	right, left *node
}

func NewAVL(less LessFunc) *AVL {
	return &AVL{less: less}
}

func (avl *AVL) find(v interface{}) *node {
	n := avl.root
	for n != nil {
		if avl.less(v, n.value) {
			n = n.left
		} else if avl.less(n.value, v) {
			n = n.right
		} else {
			return n
		}
	}
	return nil
}

func (avl *AVL) Contains(v interface{}) bool {
	return avl.find(v) != nil
}

func (avl *AVL) Append(v interface{}) {
	newNode := &node{value: v}
	if avl.root == nil {
		avl.root = newNode
		return
	}

	n := avl.root
	for {
		if avl.less(v, n.value) {
			if n.left == nil {
				n.left = newNode
				break
			}
			n = n.left
		} else if avl.less(n.value, v) {
			if n.right == nil {
				n.right = newNode
				break
			}
			n = n.right
		} else {
			return
		}
	}

	avl.balance()
}

func (avl *AVL) balance() {
	// TODO
}
