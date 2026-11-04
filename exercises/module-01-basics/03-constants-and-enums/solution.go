package constants

// Status represents an HTTP status code
type Status int

// FIX ME: Define StatusOK = 200, StatusNotFound = 404, StatusError = 500
// using explicit values (not iota, since these aren't sequential)
const (
	StatusOK       Status = 200 // FIX: should be 200
	StatusNotFound Status = 404 // FIX: should be 404
	StatusError    Status = 500 // FIX: should be 500
)

// StatusMessage returns a text message for the given status
func StatusMessage(s Status) string {
	messages := map[Status]string{
		StatusOK:       "OK",
		StatusNotFound: "Not Found",
		StatusError:    "Internal Server Error",
	}
	if msg, ok := messages[s]; ok {
		return msg
	}
	return "Unknown"
}
