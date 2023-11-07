package acronym

// Abbreviate generates TLAs
func Abbreviate(s string) string {
	var res []rune = make([]rune, len(s))
	var i int
	var wordstart bool = true

	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z' && wordstart:
			r -= 32
			fallthrough
		case r >= 'A' && r <= 'Z' && wordstart:
			res[i] = r
			wordstart = false
			i++
		case r == ' ' || r == '-':
			wordstart = true
		}
	}
	return string(res[:i])
}
