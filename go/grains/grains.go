package grains

import (
	"errors"
	"math"
)

// func Square calculates the number of rice grains on
// the Nth square of a chessboard (exponential growth)
// using math.Pow
// BenchmarkSquare-16       	 4869128	       250.6 ns/op
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("outside range")
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

// func SquareBit calculates the number of rice grains on
// the Nth square of a chessboard (exponential growth)
// using bit shifting
// BenchmarkSquareBit-16    	121615450	         9.762 ns/op
func SquareBit(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("outside range")
	}
	return 1 << (number - 1), nil
}

func Total() uint64 {
	// just return the result, it's fixed
	return 18446744073709551615

	// var sum uint64
	// for i := 1; i <= 64; i++ {
	// 	tmp, _ := Square(i)
	// 	sum += tmp
	// }
	// return sum

}
