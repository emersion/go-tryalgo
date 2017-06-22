package heap

func childrenIndexes(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

// defaultIntsLess is the default less function for integers.
func defaultIntsLess(i, j int) bool {
	return i < j
}

// Ints is a heap that stores integers.
type Ints struct {
	tree []int
	less func(i, j int) bool
}

// NewInts creates a new integer heap. If not nil, l is the initial contents of
// the heap. less reports whether i should sort before j. If nil, less defaults
// to i < j.
func NewInts(l []int, less func(i, j int) bool) *Ints {
	if less == nil {
		less = defaultIntsLess
	}
	ints := &Ints{tree: l, less: less}

	if n := len(l); len(l) > 1 {
		// We have to reorder elements so that ints.tree becomes a heap
		// Move down all elements
		// Complexity: O(n)
		// TODO: do not move down the last level
		for i := n - 1; i >= 0; i-- {
			ints.moveDown(ints.tree[i], i)
		}
	}

	return ints
}

// Len returns the number of elements in the heap.
func (ints *Ints) Len() int {
	return len(ints.tree)
}

// Min returns the lowest element in the heap.
//
// Complexity: O(1)
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
	if ints.less(v, pv) {
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
	if ri < len(ints.tree) && ints.less(ints.tree[ri], ints.tree[li]) {
		mi = ri
	}

	mv := ints.tree[mi]
	if ints.less(mv, v) {
		ints.tree[i] = mv
		ints.moveDown(v, mi)
	} else {
		ints.tree[i] = v
	}
}

// Push adds a new element to the heap.
//
// Complexity: O(log(ints.Len()))
func (ints *Ints) Push(v int) {
	ints.tree = append(ints.tree, v)
	ints.moveUp(v, len(ints.tree)-1)
}

// Pop removes the lowest element from the heap.
//
// Complexity: O(log(ints.Len()))
func (ints *Ints) Pop() (int, bool) {
	v, ok := ints.Min()
	if !ok {
		return 0, false
	}

	n := len(ints.tree) - 1
	mv := ints.tree[n]
	ints.tree = ints.tree[:n]

	if n > 0 {
		ints.moveDown(mv, 0)
	}

	return v, true
}
