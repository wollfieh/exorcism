package series

func All(n int, s string) []string {
	// string len N has (N-M+1) substrings of length M
	count := len(s) - n + 1
	res := make([]string, count)

	for i := 0; i < count; i++ {
		res[i] = s[i : i+n]
	}

	return res
}

func UnsafeFirst(n int, s string) string {
	if n < len(s) {
		return s[0:n]
	}
	return s
}

func First(n int, s string) (first string, ok bool) {
	if n <= len(s) {
		return s[0:n], true
	}
	return s, false
}
