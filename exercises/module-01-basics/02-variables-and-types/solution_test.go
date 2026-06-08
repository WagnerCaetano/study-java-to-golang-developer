package variables

import "testing"

func TestCreateUser(t *testing.T) {
	u := CreateUser("Alice", 30)
	if u.Name != "Alice" {
		t.Errorf("Name = %q, want %q", u.Name, "Alice")
	}
	if u.Age != 30 {
		t.Errorf("Age = %d, want %d", u.Age, 30)
	}
}

func TestCreateUserZeroValues(t *testing.T) {
	u := CreateUser("", 0)
	if u.Name != "" {
		t.Errorf("Empty name should be empty string, got %q", u.Name)
	}
	if u.Age != 0 {
		t.Errorf("Zero age should be 0, got %d", u.Age)
	}
}
