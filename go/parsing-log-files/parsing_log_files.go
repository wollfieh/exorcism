package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)`)

	return re.FindString(text) != ""
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*0=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	re := regexp.MustCompile(`(?i)".*password.*"`)

	for _, line := range lines {
		res := re.FindAllString(line, -1)
		count += len(res)
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i := range lines {
		if x := re.FindStringSubmatch(lines[i]); len(x) > 0 {
			lines[i] = fmt.Sprintf("[USR] %s %s", x[1], lines[i])
		}
	}

	return lines
}
