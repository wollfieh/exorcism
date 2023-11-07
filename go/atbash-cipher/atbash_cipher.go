package atbash

import (
	"strings"
)

func Atbash(s string) string {
	var cnt int
	var runestring []rune

	for _, r := range strings.ToLower(s) {
		if cnt == 5 {
			runestring = append(runestring, ' ')
			cnt = 0
		}
		switch {
		case r >= 'a' && r <= 'z':
			runestring = append(runestring, swap(r))
			cnt++
		case r >= '0' && r <= '9':
			runestring = append(runestring, r)
			cnt++
		}
	}

	return strings.Trim(string(runestring), " ")
}

// swap 'encrypts' the rune according to the schema
func swap(r rune) rune {
	offset := r - 'a'
	return 'z' - offset
}
