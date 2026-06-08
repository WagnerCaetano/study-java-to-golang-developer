package customerrors

import "testing"

func TestHTTPError(t *testing.T) {
	err := &HTTPError{Code: 404, Message: "Not Found"}
	if err.Error() != "HTTP 404: Not Found" {
		t.Errorf("Error() = %q, want %q", err.Error(), "HTTP 404: Not Found")
	}
}

func TestHTTPError500(t *testing.T) {
	err := &HTTPError{Code: 500, Message: "Internal Server Error"}
	if err.Error() != "HTTP 500: Internal Server Error" {
		t.Errorf("Error() = %q", err.Error())
	}
}
