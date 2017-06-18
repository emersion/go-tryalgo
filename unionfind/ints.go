package unionfind

type Ints struct {
	parents []int
	ranks []int
}

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
