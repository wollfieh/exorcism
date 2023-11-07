package prime

import "fmt"

// starting with some primes; calculating the rest
// (0 is not a prime)
var (
	primes []int = []int{0, 2, 3}
)

// Nth returns the nth prime number. An error must be returned
// if the nth prime number can't be calculated ('n' is equal or less than zero)
// BenchmarkNth-16         534151628                2.139 ns/op           0 B/op         0 allocs/op
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("no such prime")
	}
	for n >= len(primes) {
		primes = append(primes, nextPrime())
	}

	return primes[n], nil
}

// nextPrime calculates the next prime starting from the last
func nextPrime() int {
	// each prime  p > 3  for p>3 has the form
	// p = 6 k + 1  or
	// p = 6 k − 1
	// k being a natural number

	// get k from the biggest available prime
	k := 1 + primes[len(primes)-1]/6

	// next prime must be one of these
	for {
		if isPrime(k*6 - 1) {
			return k*6 - 1
		}
		if isPrime(k*6 + 1) {
			return k*6 + 1
		}
		// or try with next k
		k++
	}
}

// isPrime tests if a given it is prime
// prime :: only divisible by 1 and itself
func isPrime(p int) bool {
	for _, v := range primes[1:] {
		if p%v == 0 {
			return false
		}
	}

	return true
}
