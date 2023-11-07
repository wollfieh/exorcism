package lsproduct

import (
	"errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span < 1 || span > len(digits) {
		return 0, errors.New("span not in range")
	}

	numberOfSeries := len(digits) - span + 1
	var largest int64

	for i := 0; i < numberOfSeries; i++ {
		prod, err := product(digits[i : i+span])
		if err != nil {
			return 0, err
		}
		if prod > largest {
			largest = prod
		}
	}
	return largest, nil
}

func product(series string) (int64, error) {
	var prod int64 = 1

	for _, n := range series {
		if n < '0' || n > '9' {
			return 0, errors.New("not a digit")
		}
		prod *= int64(n - '0')
	}
	return prod, nil
}
