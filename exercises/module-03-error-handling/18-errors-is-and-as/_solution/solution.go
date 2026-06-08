package errorsis

import "errors"

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string { return "not found" }

var ErrUnauthorized = errors.New("unauthorized")

func CheckNotFound(err error) bool {
	var nf *NotFoundError
	return errors.As(err, &nf)
}

func CheckUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}
