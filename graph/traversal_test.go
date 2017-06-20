package graph_test

import (
	"reflect"
	"testing"

	"github.com/emersion/go-tryalgo/graph"
)

func TestDFS(t *testing.T) {
	var l []int
	graph.DFS(testGraph, 0, func(i int) {
		l = append(l, i)
	})

	//want := []int{0, 1, 3, 5, 4, 2, 6} // forward
	want := []int{0, 3, 5, 2, 6, 1, 4} // backward
	if !reflect.DeepEqual(l, want) {
		t.Errorf("graph.DFS() = %v, want %v", l, want)
	}
}

func TestBFS(t *testing.T) {
	var l []int
	graph.BFS(testGraph, 0, func(i int) {
		l = append(l, i)
	})

	want := []int{0, 1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(l, want) {
		t.Errorf("graph.BFS() = %v, want %v", l, want)
	}
}
