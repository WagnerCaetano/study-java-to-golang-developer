package customerrors

import "fmt"

type HTTPError struct {
	Code    int
	Message string
}

// HTTPError represents an HTTP error with status code and message
// FIX ME: Define HTTPError struct with Code int and Message string

func (h *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", h.Code, h.Message)
}

// Error implements the error interface
// FIX ME: Implement Error() method that returns "HTTP 404: Not Found" format
