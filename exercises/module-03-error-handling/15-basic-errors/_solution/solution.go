package basicerrors

import "errors"

var ErrNegativeAge = errors.New("age cannot be negative")
var ErrTooOld = errors.New("age cannot exceed 150")

func ValidateAge(age int) error {
	if age < 0 {
		return ErrNegativeAge
	}
	if age > 150 {
		return ErrTooOld
	}
	return nil
}
