package retry

import "time"

func Retry(maxRetries int, baseDelay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		if i < maxRetries-1 {
			delay := baseDelay * time.Duration(1<<uint(i))
			time.Sleep(delay)
		}
	}
	return err
}
