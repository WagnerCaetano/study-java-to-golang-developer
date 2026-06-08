package mocking

type Notifier interface {
	Send(to, message string) error
}

type UserService struct {
	notifier Notifier
}

func NewUserService(n Notifier) *UserService {
	return &UserService{notifier: n}
}

func (s *UserService) WelcomeUser(email string) error {
	return s.notifier.Send(email, "Welcome!")
}
