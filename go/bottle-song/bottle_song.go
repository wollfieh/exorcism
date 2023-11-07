package bottlesong

import (
	"fmt"
	"strings"
)

var (
	number []string = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
)

func Recite(startBottles, takeDown int) []string {
	var res []string

	for i := takeDown; i > 0; i-- {
		res = append(res, genVerse(startBottles)...)
		if i > 1 {
			res = append(res, "")
		}
		startBottles--
	}

	return res
}

func genVerse(bottles int) []string {
	var res []string
	var botstr []string
	switch bottles {
	case 2:
		botstr = []string{"bottles", "bottle"}
	case 1:
		botstr = []string{"bottle", "bottles"}
	default:
		botstr = []string{"bottles", "bottles"}
	}

	res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(number[bottles]), botstr[0]))
	res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(number[bottles]), botstr[0]))
	res = append(res, "And if one green bottle should accidentally fall,")
	res = append(res, fmt.Sprintf("There'll be %s green %s hanging on the wall.", number[bottles-1], botstr[1]))

	return res
}
