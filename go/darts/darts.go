package darts

import (
	"math"
)

// function Score returns the score in a darts game
// x,y are the coordinates of the impact
func Score(x, y float64) int {
	r := math.Sqrt(x*x + y*y)

	switch {
	case r <= 1.0:
		return 10
	case r <= 5.0:
		return 5
	case r <= 10:
		return 1
	default:
		return 0
	}
}
