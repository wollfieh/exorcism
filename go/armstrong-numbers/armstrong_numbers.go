package armstrong

import "math"

func IsNumber(n int) bool {
	return n == armstrong2(n)
}

// BenchmarkIsNumber-16             1692136               657.9 ns/op             0 B/op          0 allocs/op
func armstrong2(n int) int {
	var sum float64
	var l int = 1 + int(math.Log10(float64(n)))

	for n > 0 {
		sum += (math.Pow(float64(n%10), float64(l)))
		n /= 10
	}
	return int(sum)
}
