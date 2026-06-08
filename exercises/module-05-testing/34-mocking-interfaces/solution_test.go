package mocking

import "errors"
import "testing"

type MockNotifier struct {
	Err error
	Called bool
}

func (m *MockNotifier) Send(to, message string) error {
	m.Called = true
	return m.Err
}

func TestWelcomeUser(t *testing.T) {
	mock := &MockNotifier{}
	svc := NewUserService(mock)
	err := svc.WelcomeUser("test@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !mock.Called {
		t.Error("notifier was not called")
	}
}

func TestWelcomeUserError(t *testing.T) {
	mock := &MockNotifier{Err: errors.New("send failed")}
	svc := NewUserService(mock)
	err := svc.WelcomeUser("test@example.com")
	if err == nil {
		t.Error("expected error when notifier fails")
	}
}
