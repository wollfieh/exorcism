package anagram

import (
	"sort"
	"strings"
)

// using slice of runes because a) tests are unicode
// b) want to sort, which cannot be done with strings
type ms []rune

func Detect(subject string, candidates []string) []string {
	res := []string{}
	for _, v := range candidates {
		if isAnagram(subject, v) {
			res = append(res, v)
		}
	}
	return res
}

// func isAnagram determins if two strings are
// anagrams of each other
func isAnagram(subject, anagram string) bool {
	mssubject := ms(strings.ToLower(subject))
	msanagram := ms(strings.ToLower(anagram))

	if len(mssubject) != len(msanagram) {
		return false
	}
	if mssubject.equals(msanagram) {
		return false
	}

	sort.Slice(msanagram, func(i, j int) bool { return msanagram[i] < msanagram[j] })
	sort.Slice(mssubject, func(i, j int) bool { return mssubject[i] < mssubject[j] })

	return mssubject.equals(msanagram)
}

func (m ms) String() string {
	return string(m)
}
func (m ms) equals(n ms) bool {
	return string(m) == string(n)
}
