package sort

func split(a []int) ([]int, []int) {
	i := len(a) / 2
	return a[:i], a[i:]
}

func merge(a1, a2 []int) []int {
	a := make([]int, len(a1)+len(a2))
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] <= a2[j] {
			a[i+j] = a1[i]
			i++
		} else {
			a[i+j] = a2[j]
			j++
		}
	}
	for i < len(a1) {
		a[i+j] = a1[i]
		i++
	}
	for j < len(a2) {
		a[i+j] = a2[j]
		j++
	}
	return a
}

// Merge is the merge sort. Complexity: O(n log n)
func Merge(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	a1, a2 := split(a)
	a1, a2 = Merge(a1), Merge(a2)
	return merge(a1, a2)
}
