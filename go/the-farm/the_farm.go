package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (err *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", err.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	// Duplicate fodder if scale malfunctions
	if err == ErrScaleMalfunction && amount > 0 {
		amount *= 2
	}

	// Return error if it is not a scale malfunction
	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}

	// Cannot have negative fodder
	if amount < 0 {
		return 0, errors.New("negative fodder")
	}

	// Division by zero error
	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	// Handle negative cows
	if cows < 0 {
		return 0, &SillyNephewError{cows}
	}

	return amount / float64(cows), nil
}
