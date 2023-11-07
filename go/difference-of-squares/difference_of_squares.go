package diffsquares

func SquareOfSum(n int) int {

	a := (1 + n) * n / 2
	return int(a * a)
}

func SumOfSquares(n int) int {
	// https://en.wikipedia.org/wiki/Square_pyramidal_number
	return int(n * (n + 1) * (2*n + 1) / 6)
}

func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)
}
