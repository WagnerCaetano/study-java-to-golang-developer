package errorsis

import "errors"

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return "not found"
}

var ErrUnauthorized = errors.New("unauthorized")

// CheckNotFound should return true if the error wraps a NotFoundError
// FIX ME: Use errors.As properly
func CheckNotFound(err error) bool {
	var notFoundError *NotFoundError
	return errors.As(err, &notFoundError)
}

// CheckUnauthorized should return true if the error is ErrUnauthorized
// FIX ME: Use errors.Is properly
func CheckUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}
