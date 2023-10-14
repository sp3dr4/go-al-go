package twocrystalballs

import "math"

func TwoCrystalBalls(breaks []bool) int {
	jump := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jump
	for ; i < len(breaks); i += jump {
		if breaks[i] {
			break
		}
	}
	i -= jump
	for j := 0; j < jump && i < len(breaks); i, j = i+1, j+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}
