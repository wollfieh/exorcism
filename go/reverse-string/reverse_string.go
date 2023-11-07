package reverse

func Reverse(input string) string {
	var r []rune = []rune(input)

	for s, e := 0, len(r)-1; s < len(r)/2; s, e = s+1, e-1 {
		r[s], r[e] = r[e], r[s]
	}
	return string(r)
}
