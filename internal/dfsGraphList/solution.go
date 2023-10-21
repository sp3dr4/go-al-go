package dfsgraphlist

type GraphEdge struct {
	to int
}
type WeightedAdjacencyList [][]GraphEdge

func walk(graph *WeightedAdjacencyList, curr int, needle int, seen *[]bool, path *[]int) bool {
	if (*seen)[curr] {
		return false
	}
	(*seen)[curr] = true

	*path = append(*path, curr)
	if curr == needle {
		return true
	}

	list := (*graph)[curr]
	for i := 0; i < len(list); i++ {
		if walk(graph, list[i].to, needle, seen, path) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func Dfs(graph *WeightedAdjacencyList, source int, needle int) []int {
	seen := make([]bool, len(*graph))
	path := []int{}

	walk(graph, source, needle, &seen, &path)

	return path
}
