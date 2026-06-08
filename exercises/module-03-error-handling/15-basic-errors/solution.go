package basicerrors

import "errors"

// ValidateAge checks if age is valid (0-150)
// FIX ME: Return error for negative ages and ages > 150
func ValidateAge(age int) error {
	// FIX: Add validation
	return nil
}

var ErrNegativeAge = errors.New("age cannot be negative")
var ErrTooOld = errors.New("age cannot exceed 150")
