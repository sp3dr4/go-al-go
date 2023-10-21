package dijkstralist

import (
	"math"
	"slices"
)

type GraphEdge struct {
	to     int
	weight int
}

type WeightedAdjacencyList [][]GraphEdge

// would be better to use minHeap to speed it up
func hasUnvisited(seen *[]bool, dists *[]int) bool {
	for i, s := range *seen {
		if !s && (*dists)[i] < math.MaxInt {
			return true
		}
	}
	return false
}

func getLowestUnvisited(seen *[]bool, dists *[]int) int {
	idx, lowest := -1, math.MaxInt

	for i := 0; i < len(*seen); i++ {
		if (*seen)[i] {
			continue
		}
		if lowest > (*dists)[i] {
			lowest = (*dists)[i]
			idx = i
		}
	}

	return idx
}

func DijkstraList(arr *WeightedAdjacencyList, source int, sink int) []int {
	arrLen := len(*arr)
	seen := make([]bool, arrLen)
	prev := make([]int, arrLen)
	for i := range prev {
		prev[i] = -1
	}
	dists := make([]int, arrLen)
	for i := range dists {
		dists[i] = math.MaxInt
	}

	dists[source] = 0

	for hasUnvisited(&seen, &dists) {
		curr := getLowestUnvisited(&seen, &dists)
		seen[curr] = true

		adjs := (*arr)[curr]
		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.to] {
				continue
			}

			dist := dists[curr] + edge.weight
			if dist < dists[edge.to] {
				dists[edge.to] = dist
				prev[edge.to] = curr
			}
		}
	}

	curr := sink
	out := []int{}
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) == 0 {
		return []int{}
	}
	out = append(out, source)
	slices.Reverse(out)
	return out
}
