package wrapping

import (
	"errors"
	"fmt"
)

var ErrConfigNotFound = errors.New("config not found")

func ReadConfig(path string) error {
	return fmt.Errorf("readConfig(%s): %w", path, ErrConfigNotFound)
}

func GetRootCause(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
