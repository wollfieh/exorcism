package summultiples

import "fmt"

type set []int

func SumMultiples(limit int, divisors ...int) int {
	var res set

	for _, v := range divisors {
		res = res.Combine(NewSet(limit, v))
	}
	return res.Sum()
}

func NewSet(limit, divisor int) set {
	if divisor == 0 {
		return set{0}
	}
	a := set{}
	for i := divisor; i < limit; i = i + divisor {
		a = append(a, i)
	}
	return a
}

func (a set) Sum() int {
	var sum int
	for _, i := range a {
		sum += i
	}
	return sum
}
func (a set) Combine(b set) set {
	var apos, bpos int
	var res set

	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	fmt.Println("len ", len(a), len(b))

	for {
		switch {
		case apos == len(a):
			res = append(res, b[bpos:]...)
			return res
		case bpos == len(b):
			res = append(res, a[apos:]...)
			return res
		case a[apos] == b[bpos]:
			res = append(res, a[apos])
			apos++
			bpos++
		case a[apos] < b[bpos]:
			res = append(res, a[apos])
			apos++
		case a[apos] > b[bpos]:
			res = append(res, b[bpos])
			bpos++
		}
	}
}
