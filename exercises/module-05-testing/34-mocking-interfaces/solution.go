package mocking

import "fmt"

// Notifier sends notifications
type Notifier interface {
	Send(to, message string) error
}

// UserService uses a Notifier
type UserService struct {
	notifier Notifier
}

func NewUserService(n Notifier) *UserService {
	return &UserService{notifier: n}
}

// WelcomeUser sends a welcome notification
// FIX ME: The error from notifier.Send is not handled
func (s *UserService) WelcomeUser(email string) error {
	s.notifier.Send(email, "Welcome!")
	return nil // FIX: Return the error from Send
}
