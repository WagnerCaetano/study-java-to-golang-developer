package selectstmt

import (
	"testing"
	"time"
)

func TestFirstResultCh1(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "from ch1"
	ch2 <- "from ch2"
	result := FirstResult(ch1, ch2)
	if result != "from ch1" && result != "from ch2" {
		t.Errorf("FirstResult = %q, expected a channel value", result)
	}
}

func TestFirstResultTimeout(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "delayed"
	}()
	result := FirstResult(ch1, ch2)
	if result != "delayed" {
		t.Errorf("FirstResult = %q, want %q", result, "delayed")
	}
}
