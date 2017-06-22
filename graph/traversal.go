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

type iterativeDFSItem struct {
	node  int
	depth int
}

// IterativeDFS performs an iterative depth-first search on a graph.
//
// Iterative DFS behaves like BFS but uses DFS internally. It can be used to
// traverse very large graphs since it has a memory usage much lower than BFS.
// Though, IterativeDFS is ~2 times slower than DFS for a binary tree.
func IterativeDFS(g Graph, src int, f func(int)) {
	seen := make(map[int]int)
	seen[src] = 0
	f(src)
	maxDepth := 1
	for {
		stack := []iterativeDFSItem{{src, 0}}
		newNodes := 0
		for len(stack) > 0 {
			item := stack[len(stack)-1]
			node, depth := item.node, item.depth
			stack = stack[:len(stack)-1]
			if depth >= maxDepth {
				continue
			}
			for _, neighbour := range g.Neighbours(node) {
				v, ok := seen[neighbour]
				if !ok {
					f(neighbour)
					newNodes++
				}
				if v >= maxDepth {
					continue
				}
				seen[neighbour] = maxDepth
				stack = append(stack, iterativeDFSItem{neighbour, depth + 1})
			}
		}

		if newNodes == 0 {
			break
		}

		maxDepth++
	}
}
