package strain

// I think the types could use some cleanup
type Ints []int
type Lists [][]int
type Strings []string

type containers interface{ int | []int | string }

func applyFilter[C containers](filter func(C) bool, sl []C) []C {
	var res []C

	for _, v := range sl {
		if filter(v) {
			res = append(res, v)
		}
	}

	return res
}

func (i Ints) Keep(filter func(int) bool) Ints {
	return applyFilter(filter, i)
}

func (i Ints) Discard(filter func(int) bool) Ints {
	// applying filter() is Keep, so
	// applying !filter() is Discard
	return applyFilter(func(n int) bool { return !filter(n) }, i)
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	return applyFilter(filter, l)
}

func (s Strings) Keep(filter func(string) bool) Strings {
	return applyFilter(filter, s)
}
