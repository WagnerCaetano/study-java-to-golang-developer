package constants

import "testing"

func TestStatusValues(t *testing.T) {
	if StatusOK != 200 {
		t.Errorf("StatusOK = %d, want 200", StatusOK)
	}
	if StatusNotFound != 404 {
		t.Errorf("StatusNotFound = %d, want 404", StatusNotFound)
	}
	if StatusError != 500 {
		t.Errorf("StatusError = %d, want 500", StatusError)
	}
}

func TestStatusMessage(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{StatusOK, "OK"},
		{StatusNotFound, "Not Found"},
		{StatusError, "Internal Server Error"},
		{Status(999), "Unknown"},
	}
	for _, tt := range tests {
		result := StatusMessage(tt.status)
		if result != tt.expected {
			t.Errorf("StatusMessage(%d) = %q, want %q", tt.status, result, tt.expected)
		}
	}
}
