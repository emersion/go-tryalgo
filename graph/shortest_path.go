package graph

import (
	"github.com/emersion/go-tryalgo/heap"
)

func reverse(l []int) {
	m := len(l)/2
	for i := 0; i < m; i++ {
		j := len(l)-i-1
		l[i], l[j] = l[j], l[i]
	}
}

// Dijkstra computes the shortest path from src to dst using the Dijkstra
// algorithm. If dst cannot be reached, nil is returned.
func Dijkstra(g ValuedGraph, src, dst int) []int {
	// TODO: use slices instead of maps
	dist := map[int]int{src: 0}
	prev := make(map[int]int)
	h := heap.NewInts([]int{src}, func (i, j int) bool {
		return dist[i] < dist[j]
	})

	for {
		node, ok := h.Pop()
		if !ok {
			return nil
		}
		if node == dst {
			break
		}

		nodeDist := dist[node]
		for _, neighbour := range g.Neighbours(node) {
			d, ok := dist[neighbour]
			var newDist int
			if ok {
				newDist = nodeDist + g.Distance(node, neighbour)
			}
			if !ok || d > newDist {
				dist[neighbour] = newDist
				prev[neighbour] = node
				h.Push(neighbour)
			}
		}
	}

	var l []int
	node := dst
	for node != src {
		l = append(l, node)
		node = prev[node]
	}
	l = append(l, src)
	reverse(l)
	return l
}
