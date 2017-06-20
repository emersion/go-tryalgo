package graph

// DFS performs a depth-first search on a graph.
func DFS(g Graph, src int, f func(int)) {
	seen := make(map[int]struct{})
	seen[src] = struct{}{}
	stack := []int{src}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		f(node)
		for _, neighbour := range g.Neighbours(node) {
			if _, ok := seen[neighbour]; ok {
				continue
			}
			seen[neighbour] = struct{}{}
			stack = append(stack, neighbour)
		}
	}
}

// DFS performs a breadth-first search on a graph.
func BFS(g Graph, src int, f func(int)) {
	seen := make(map[int]struct{})
	seen[src] = struct{}{}
	current := []int{src}
	var next []int
	for len(current) > 0 {
		for _, node := range current {
			f(node)
			for _, neighbour := range g.Neighbours(node) {
				if _, ok := seen[neighbour]; ok {
					continue
				}
				seen[neighbour] = struct{}{}
				next = append(next, neighbour)
			}
		}

		current = next
		next = nil
	}
}
