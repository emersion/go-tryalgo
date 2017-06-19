package bintree

type LessFunc func(a, b interface{}) bool

type node struct {
	value interface{}
	h int
	left, right *node
}

func (n *node) height() int {
	if n == nil {
		return 0
	}
	return n.h
}

func (n *node) computeHeight() {
	hl, hr := n.left.height(), n.right.height()

	h := hl
	if hr > hl {
		h = hr
	}
	n.h = 1 + h
}

// Assumes n != nil && n.left != nil
func (n *node) rotateRight() *node {
	l := n.left
	l.right, n.left = n, l.right
	n.computeHeight()
	l.computeHeight()
	return l
}

// Assumes n != nil && n.right != nil
func (n *node) rotateLeft() *node {
	r := n.right
	n.right, r.left = r.left, n
	n.computeHeight()
	r.computeHeight()
	return r
}

func (n *node) balance() *node {
	if n == nil {
		return nil
	}

	l, r := n.left, n.right
	hl, hr := l.height(), r.height()
	if hl > hr + 1 {
		ll, lr := l.left, l.right
		if ll.height() < lr.height() {
			n.left = l.rotateLeft()
		}
		return n.rotateRight()
	} else if hr > hl + 1 {
		rl, rr := r.left, r.right
		if rr.height() < rl.height() {
			n.right = r.rotateRight()
		}
		return n.rotateLeft()
	} else {
		n.computeHeight()
		return n
	}
}

type AVL struct {
	less LessFunc
	root *node
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

	avl.root = avl.root.balance()
}
