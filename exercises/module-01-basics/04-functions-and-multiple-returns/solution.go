package functions

import "errors"

// Divide divides a by b and returns the result
// FIX ME: Return an error when b is 0, otherwise return the result and nil error
func Divide(a, b float64) (float64, error) {
	// FIX: Add division by zero check
	return a / b, nil
}
