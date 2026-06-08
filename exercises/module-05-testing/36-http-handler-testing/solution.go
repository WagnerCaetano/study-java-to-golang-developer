package httptest

import "net/http"

// HealthHandler returns a simple health check response
// FIX ME: Should return 200 with "OK" body
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// FIX: Write proper response
}
