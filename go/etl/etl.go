package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	res := make(map[string]int)

	for value, letters := range in {
		for _, l := range letters {
			res[strings.ToLower(l)] = value
		}
	}
	return res

}
