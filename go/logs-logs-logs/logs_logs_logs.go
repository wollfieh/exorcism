package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	r_unes := map[rune]string{
		'\U00002757': "recommendation",
		'\U0001F50D': "search",
		'\U00002600': "weather",
	}
	for _, r := range log {
		_, found := r_unes[r]
		if found {
			return r_unes[r]
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	return strings.Replace(log, string(oldRune), string(newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	return len([]rune(log)) <= limit
}
