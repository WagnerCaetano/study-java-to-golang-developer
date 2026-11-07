package panicrecover

import "fmt"

// SafeExecute runs fn and converts any panic into an error
// FIX ME: Use defer + recover to catch panics
func SafeExecute(fn func()) (err error) {
	// FIX: Add defer with recover

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	fn()
	return nil
}
