package sieve

// slop -- slice of primes
// index is the actual numbers; values are
// 0 initialized; all indexes with value 0 are prime
// -1 multiples of found primes are marked -1
type slop struct {
	sl  []int
	pos int
}

func newSlop(pos int, limit int) slop {
	return slop{pos: pos, sl: make([]int, limit+1)}
}

// append next prime to result
func (a *slop) append(num int) {
	a.sl[a.pos] = num
	a.pos++
}

// nextPrime returns next prime. At end of sive slice
// bool parameter is false
func (a *slop) nextPrime() (int, bool) {
	if a.pos >= len(a.sl) {
		return -11, false
	}
	for a.sl[a.pos] != 0 { // skip to next prime or end of slice
		a.pos++
		if a.pos >= len(a.sl) {
			return 0, false
		}
	}
	// prime is at a.pos; scratch all multiples
	result := a.pos

	for i, step := a.pos, a.pos; i < len(a.sl); i += step {
		a.sl[i] = -1
	}

	a.pos++
	return result, true

}

// slice of result
func (a slop) slice() []int {
	return a.sl[:a.pos]
}

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}

	primes := newSlop(2, limit)
	result := newSlop(0, limit)

	for {
		r, ok := primes.nextPrime()
		if !ok {
			break
		}
		result.append(r)
	}

	return result.slice()
}
