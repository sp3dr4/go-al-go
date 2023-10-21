package bfsgraphmatrix

import (
	"slices"
	"testing"
)

var matrix WeightedAdjacencyMatrix = [][]int{
	{0, 3, 1, 0, 0, 0, 0},
	{0, 0, 0, 0, 1, 0, 0},
	{0, 0, 7, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0},
	{0, 1, 0, 5, 0, 2, 0},
	{0, 0, 18, 0, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 1},
}

func TestBfs(t *testing.T) {
	res := Bfs(matrix, 0, 6)
	expected := []int{0, 1, 4, 5, 6}
	if slices.Compare(res, expected) != 0 {
		t.Errorf("expected %v, got %v", expected, res)
	}

	res = Bfs(matrix, 6, 0)
	expected = []int{}
	if slices.Compare(res, expected) != 0 {
		t.Errorf("expected %v, got %v", expected, res)
	}
}
