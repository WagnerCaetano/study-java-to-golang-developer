package sentinel

import "errors"

// FIX ME: Define sentinel errors
var (
	ErrNotFound      = errors.New("user not found")      // FIX: "user not found"
	ErrAlreadyExists = errors.New("user already exists") // FIX: "user already exists"
	ErrInvalidInput  = errors.New("invalid input")       // FIX: "invalid input"
)

// FindUser simulates finding a user
func FindUser(id int) error {
	if id <= 0 {
		return ErrInvalidInput
	}
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

// CreateUser simulates creating a user
func CreateUser(id int) error {
	if id == 1 {
		return ErrAlreadyExists
	}
	return nil
}
