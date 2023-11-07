package main

import "fmt"

type Numbers interface{ int | float64 }

func Keep[P Numbers](a []P, filter func(P) bool) []P {
	var res []P

	for _, v := range a {
		if filter(v) {
			res = append(res, v)
		}

	}
	return res
}
func Discard[P Numbers](a []P, filter func(P) bool) []P {
	return Keep(a, func(p P) bool { return !filter(p) })
}
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	fmt.Println("ints: ", a)
	fmt.Println("floats: ", b)

	fmt.Println("keep evens: ", Keep(a, func(i int) bool { return i%2 == 0 }))
	fmt.Println("keep >=3: ", Keep(b, func(i float64) bool { return i >= 3.0 }))

	fmt.Println("discard evens:", Discard(a, func(i int) bool { return i%2 == 0 }))
	fmt.Println("discard >=3 ", Discard(b, func(i float64) bool { return i >= 3.0 }))

}
