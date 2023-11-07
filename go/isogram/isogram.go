package isogram

import "strings"

// Function IsIsoram tests if given string is an isogram, meaning
// each letter except '-' and ' ' only appearing once
// this is case insensitive
func IsIsogram(word string) bool {
	m := make(map[rune]bool)

	for _, r := range strings.ToLower(word) {
		switch r {
		case '-', ' ':
			continue
		default:
			exists := m[r]
			if exists {
				return false
			}
		}
		m[r] = true
	}

	return true
}
