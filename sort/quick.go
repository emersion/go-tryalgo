package sort

func partition(a []int) ([]int, []int) {
	// TODO: choose a good pivot
	pivot := a[0]

	i, j := 0, len(a)-1
	for {
		for a[i] < pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}
		if i >= j {
			return a[:j], a[j+1:]
		}
		a[i], a[j] = a[j], a[i]
	}
}

// Quick is the quicksort. Complexity: O(n log n)
func Quick(a []int) {
	if len(a) <= 1 {
		return
	}

	a1, a2 := partition(a)
	Quick(a1)
	Quick(a2)
}
