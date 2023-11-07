package isbn

func IsValidISBN(isbn string) bool {
	var sum int
	var pos int = 10
	var c rune

	if len(isbn) > 13 || len(isbn) < 10 {
		return false
	}

	for _, c = range isbn {

		switch {
		case c == '-':
			continue

		case c >= '0' && c <= '9':
			sum += int(c-'0') * pos

		case pos == 1 && (c == 'x' || c == 'X'):
			// X only valid at this pos
			sum += 10 * pos

		default:
			// invalid character
			return false
		}

		pos--

	}
	// should be at max pos
	// sum is mod 11
	return pos == 0 && sum%11 == 0

}
