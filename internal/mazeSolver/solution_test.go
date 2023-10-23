package mazesolver

import (
	"reflect"
	"testing"
)

var maze1 = []string{
	"xxxxxxxxxx x",
	"x        x x",
	"x        x x",
	"x xxxxxxxx x",
	"x          x",
	"x xxxxxxxxxx",
}

var maze1Expected = []Point{
	{x: 10, y: 0},
	{x: 10, y: 1},
	{x: 10, y: 2},
	{x: 10, y: 3},
	{x: 10, y: 4},
	{x: 9, y: 4},
	{x: 8, y: 4},
	{x: 7, y: 4},
	{x: 6, y: 4},
	{x: 5, y: 4},
	{x: 4, y: 4},
	{x: 3, y: 4},
	{x: 2, y: 4},
	{x: 1, y: 4},
	{x: 1, y: 5},
}

func TestMazeSolver(t *testing.T) {
	result := Solve(&maze1, "x", &Point{x: 10, y: 0}, &Point{x: 1, y: 5})
	if !reflect.DeepEqual(result, maze1Expected) {
		t.Error("output path does not match with expected path")
	}
}
