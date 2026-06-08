package sentinel

import "errors"

var (
	ErrNotFound      = errors.New("user not found")
	ErrAlreadyExists = errors.New("user already exists")
	ErrInvalidInput  = errors.New("invalid input")
)

func FindUser(id int) error {
	if id <= 0 {
		return ErrInvalidInput
	}
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

func CreateUser(id int) error {
	if id == 1 {
		return ErrAlreadyExists
	}
	return nil
}
