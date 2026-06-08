package panicrecover

import "testing"

func TestSafeExecuteNoPanic(t *testing.T) {
	called := false
	err := SafeExecute(func() {
		called = true
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !called {
		t.Error("function was not called")
	}
}

func TestSafeExecuteWithPanic(t *testing.T) {
	err := SafeExecute(func() {
		panic("something went wrong")
	})
	if err == nil {
		t.Fatal("expected error from panic")
	}
	if err.Error() != "panic: something went wrong" {
		t.Errorf("error = %q, want to contain panic message", err.Error())
	}
}
