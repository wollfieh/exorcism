package hamming

import "errors"

// function Distance computes the hamming distance (differences) between
// two equal length strands of DNA
func Distance(a, b string) (int, error) {
	var ham int

	if len(a) != len(b) {
		return 0, errors.New("not same length")
	}

	for i := range a {
		if a[i] != b[i] {
			ham++
		}
	}

	return ham, nil
}
