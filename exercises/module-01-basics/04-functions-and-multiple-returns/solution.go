package functions

import "errors"

// Divide divides a by b and returns the result
// FIX ME: Return an error when b is 0, otherwise return the result and nil error
func Divide(a, b float64) (res float64, err error) {
	// FIX: Add division by zero check
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
