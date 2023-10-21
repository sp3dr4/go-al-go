package bfsgraphmatrix

import (
	"slices"
)

type WeightedAdjacencyMatrix [][]int

func Bfs(graph WeightedAdjacencyMatrix, source int, needle int) []int {

	seen := make([]bool, len(graph))
	seen[source] = true

	prev := make([]int, len(graph))
	for i := range prev {
		prev[i] = -1
	}

	q := []int{source}

	for len(q) > 0 {
		var curr int
		curr, q = q[0], q[1:]
		if curr == needle {
			break
		}

		adjs := graph[curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 || seen[i] {
				continue
			}
			seen[i] = true
			prev[i] = curr
			q = append(q, i)
		}
	}

	curr := needle
	out := []int{}
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) > 0 {
		out = append(out, source)
		slices.Reverse(out)
		return out
	}
	return []int{}
}
