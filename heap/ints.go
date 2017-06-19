package heap

func childrenIndexes(i int) (int, int) {
	return 2*i+1, 2*i+2
}

func parentIndex(i int) int {
	return (i-1)/2
}

type Ints struct {
	tree []int
}

func NewInts() *Ints {
	return new(Ints)
}

func (ints *Ints) Len() int {
	return len(ints.tree)
}

func (ints *Ints) Min() (int, bool) {
	if len(ints.tree) == 0 {
		return 0, false
	}
	return ints.tree[0], true
}

func (ints *Ints) moveUp(v, i int) {
	if i == 0 {
		ints.tree[i] = v
		return
	}

	pi := parentIndex(i)
	pv := ints.tree[pi]
	if pv > v {
		ints.tree[i] = pv
		ints.moveUp(v, pi)
	} else {
		ints.tree[i] = v
	}
}

func (ints *Ints) moveDown(v, i int) {
	li, ri := childrenIndexes(i)
	if li >= len(ints.tree) {
		// No more children
		ints.tree[i] = v
		return
	}

	// Select the child with the lowest value
	mi := li
	if ri < len(ints.tree) && ints.tree[li] > ints.tree[ri] {
		mi = ri
	}

	mv := ints.tree[mi]
	if mv < v {
		ints.tree[i] = mv
		ints.moveDown(v, mi)
	} else {
		ints.tree[i] = v
	}
}

func (ints *Ints) Append(v int) {
	ints.tree = append(ints.tree, v)
	ints.moveUp(v, len(ints.tree)-1)
}

func (ints *Ints) RemoveMin() {
	n := len(ints.tree) - 1
	if n < 0 {
		return
	}

	mv := ints.tree[n]
	ints.tree = ints.tree[:n]

	if n > 0 {
		ints.moveDown(mv, 0)
	}
}
