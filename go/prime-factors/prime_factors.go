package prime

func Factors(n int64) []int64 {
	res := make([]int64, 0)
	var pf int64 = 2

	for n > 1 {
		for n%pf == 0 {
			res = append(res, pf)
			n /= pf
		}
		pf++
	}

	return res
}
