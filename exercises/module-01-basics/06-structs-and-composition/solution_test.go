package structs

import "testing"

func TestFullName(t *testing.T) {
	p := Person{FirstName: "John", LastName: "Doe", Age: 30}
	result := p.FullName()
	expected := "John Doe"
	if result != expected {
		t.Errorf("FullName() = %q, want %q", result, expected)
	}
}

func TestFullNameEmpty(t *testing.T) {
	p := Person{}
	if p.FullName() != " " {
		t.Errorf("Empty person full name should be a single space")
	}
}
