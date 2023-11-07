// Package bob is about https://exercism.org/tracks/go/exercises/bob
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
)

// Hey delivers different responses to remarks
func Hey(remark string) string {

	if isSilence(remark) {
		return "Fine. Be that way!"
	}

	if isNonLetter(remark) && isQuestion(remark) {
		return "Sure."
	}

	if isNonLetter(remark) {
		return "Whatever."
	}

	if isQuestion(remark) && isShouting(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isQuestion(remark) {
		return "Sure."
	}

	if isShouting(remark) {
		return "Whoa, chill out!"
	}

	// default
	return "Whatever."
}

func isQuestion(remark string) bool {
	re := regexp.MustCompile(`.*\?\s*$`)
	return re.Match([]byte(remark))
}

func isShouting(remark string) bool {
	re := regexp.MustCompile(`^[[:digit:][:upper:]\s[:punct:]]+$`)
	return re.Match([]byte(remark))
}

func isSilence(remark string) bool {
	re := regexp.MustCompile(`^\s*$`)
	return re.Match([]byte(remark))
}

func isNonLetter(remark string) bool {
	re := regexp.MustCompile(`^[[:digit:][:punct:]\s]+$`)
	return re.Match([]byte(remark))
}
