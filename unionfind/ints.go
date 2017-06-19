package unionfind

// Ints represents disjoint sets of integers.
//
// Complexity in memory: O(n), n being the capacity
type Ints struct {
	parents []int
	ranks []int
}

// NewInts creates a new disjoint sets container containing one class for each
// integer from 0 to n-1.
func NewInts(n int) *Ints {
	parents := make([]int, n, n)
	for i := range parents {
		parents[i] = i
	}

	return &Ints{
		parents: parents,
		ranks: make([]int, n, n),
	}
}

// Find returns the representative member of i. Two integers i and j are in the
// same set if and only if Find(i) == Find(j).
//
// Complexity: O(α(n))
func (ints *Ints) Find(i int) int {
	j := i
	for {
		p := ints.parents[j]
		if p == j {
			// Path compression
			ints.parents[i] = p
			return j
		}
		j = p
	}
}

// Union merges the sets where i and j are.
//
// Complexity: O(α(n))
func (ints *Ints) Union(i, j int) {
	pi := ints.Find(i)
	pj := ints.Find(j)
	if pi == pj {
		return
	}

	if ints.ranks[pi] > ints.ranks[pj] {
		ints.parents[pj] = pi
	} else {
		if ints.ranks[pi] == ints.ranks[pj] {
			ints.ranks[pj]++
		}
		ints.parents[pi] = pj
	}
}
