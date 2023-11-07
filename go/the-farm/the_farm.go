package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	num int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.num)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	weight, err := weightFodder.FodderAmount()

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if err != nil {
		if err == ErrScaleMalfunction {
			weight *= 2
		} else {
			return 0, err
		}
	}

	if weight < 0 {
		return 0, errors.New("negative fodder")
	}
	if cows < 0 {
		return 0, &SillyNephewError{cows}
	}
	return weight / float64(cows), nil
}
