package pangram

func IsPangram(input string) bool {
	var letters map[rune]bool = map[rune]bool{}

	for _, r := range input {

		switch {
		case r >= 'A' && r <= 'Z':
			r += 32 // make lowercase
			fallthrough
		case r >= 'a' && r <= 'z':
			letters[r] = true
		default:
			// skip non a-z chars
			continue
		}
	}
	return len(letters) == 26
}
