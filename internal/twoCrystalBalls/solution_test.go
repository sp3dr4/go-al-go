package twocrystalballs

import (
	"math"
	"math/rand"
	"testing"
)

func createInput(n int) ([]bool, int) {
	idx := int(math.Floor(rand.Float64() * 10000))
	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}
	return data, idx
}

func TestTwoCrystalBalls(t *testing.T) {
	input, inputIdx := createInput(10000)

	tests := []struct {
		list     []bool
		expected int
	}{
		{input, inputIdx},
		{make([]bool, 821), -1},
	}

	for _, test := range tests {
		result := TwoCrystalBalls(test.list)
		if result != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result)
		}
	}
}
