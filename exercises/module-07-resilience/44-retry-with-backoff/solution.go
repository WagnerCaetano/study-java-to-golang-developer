package retry

import "time"

// Retry executes fn with exponential backoff
// FIX ME: Implement retry logic with backoff
func Retry(maxRetries int, baseDelay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		// FIX: Calculate backoff delay and sleep
	}
	return err
}
