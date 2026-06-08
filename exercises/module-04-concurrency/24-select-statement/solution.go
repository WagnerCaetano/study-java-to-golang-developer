package selectstmt

// FirstResult returns the first value received from either ch1 or ch2
// FIX ME: Use select to wait on both channels
func FirstResult(ch1, ch2 <-chan string) string {
	// FIX: Use select to read from whichever channel is ready first
	return ""
}
