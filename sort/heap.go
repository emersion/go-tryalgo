package sort

func childrenIndexes(i int) (li int, ri int) {
	li, ri = 2 * i + 1, 2 * i + 2
	return
}

func moveDown(a []int, i, v int) {
	li, ri := childrenIndexes(i)
	if li >= len(a) {
		// No more child, stop here
		a[i] = v
		return
	}

	// Take the greatest child
	ci := li
	if ri < len(a) && a[li] < a[ri] {
		ci = ri
	}

	// If v is greater than the greatest child, stop here
	if v >= a[ci] {
		a[i] = v
	} else {
		a[i] = a[ci]
		moveDown(a, ci, v)
	}
}

// Heap is the heap sort. Complexity: O(n log n)
func Heap(a []int) {
	// Build the heap from the end of the array
	for i := len(a) - 1; i >= 0; i-- {
		moveDown(a, i, a[i])
	}

	// Extract the heap's root, put it at the end of the array and reconstruct the
	// heap
	for i := len(a) - 1; i >= 1; i-- {
		a[0], a[i] = a[i], a[0]
		moveDown(a[:i], 0, a[0])
	}
}
