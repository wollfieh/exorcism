package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if s.Length() == 0 {
		return initial
	}

	for i := 0; i < s.Length(); i++ {
		initial = fn(initial, s[i])
	}

	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if s.Length() == 0 {
		return initial
	}

	for i := s.Length() - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}

	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make([]int, s.Length())
	i := 0

	for _, v := range s {
		if fn(v) {
			res[i] = v
			i++
		}
	}
	return res[:i]
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	res := []int{}
	for _, v := range s {
		res = append(res, fn(v))
	}
	return res
}

func (s IntList) Reverse() IntList {
	for i := 0; i < s.Length()/2; i++ {
		s[i], s[s.Length()-1-i] = s[s.Length()-1-i], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}
	return s
}
