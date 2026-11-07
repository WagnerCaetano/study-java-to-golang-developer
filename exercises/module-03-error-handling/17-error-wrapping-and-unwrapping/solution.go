package wrapping

import (
	"errors"
	"fmt"
)

var ErrConfigNotFound = errors.New("config not found")

// ReadConfig simulates reading a config file
// FIX ME: Wrap the error with context using %w
func ReadConfig(path string) error {
	return fmt.Errorf("Reason:%w", ErrConfigNotFound) // FIX: Wrap with "readConfig(<path>):"
}

// GetRootCause unwraps an error chain to find the root cause
func GetRootCause(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
