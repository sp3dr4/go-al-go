package mazesolver

import "reflect"

type Point struct {
	x int
	y int
}

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func walk(maze *[]string, wall string, curr *Point, end *Point, seen *[][]bool, path *[]Point) bool {
	if curr.x < 0 || curr.x >= len((*maze)[0]) || curr.y < 0 || curr.y >= len(*maze) {
		return false
	}
	if string((*maze)[curr.y][curr.x]) == wall {
		return false
	}
	if reflect.DeepEqual(*curr, *end) {
		*path = append(*path, *end)
		return true
	}
	if (*seen)[curr.y][curr.x] {
		return false
	}

	(*seen)[curr.y][curr.x] = true
	*path = append(*path, *curr)

	for i := 0; i < len(dirs); i++ {
		newCurr := &Point{x: curr.x + dirs[i][0], y: curr.y + dirs[i][1]}
		res := walk(maze, wall, newCurr, end, seen, path)
		if res {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func Solve(maze *[]string, wall string, start *Point, end *Point) []Point {
	seen := make([][]bool, len(*maze))
	for i := range seen {
		seen[i] = make([]bool, len((*maze)[0]))
	}

	path := []Point{}

	walk(maze, wall, start, end, &seen, &path)

	return path
}
