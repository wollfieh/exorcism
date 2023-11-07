package luhn

// Function Valid checks a Number (as string) if it's a valid
// according to the Luhn algorithm, e.g. for checking
// credit card numbers
func Valid(id string) bool {
	var total, strpos, count int

	for strpos = len(id) - 1; strpos >= 0; strpos-- {

		switch {
		case id[strpos] == ' ':
			continue

		case id[strpos] >= '0' && id[strpos] <= '9':
			if count%2 != 0 {
				total += double(id[strpos])
			} else {
				total += int(id[strpos] - '0')
			}
			count++

		default:
			return false
		}
	}

	return total%10 == 0 && count > 1
}

func double(b byte) int {
	var total int
	total = int(b-'0') * 2
	if total > 9 {
		total -= 9
	}

	return total
}
