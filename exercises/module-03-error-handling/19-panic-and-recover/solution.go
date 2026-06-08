package panicrecover

import "fmt"

// SafeExecute runs fn and converts any panic into an error
// FIX ME: Use defer + recover to catch panics
func SafeExecute(fn func()) error {
	// FIX: Add defer with recover
	fn()
	return nil
}
