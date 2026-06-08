package interfaces

import "testing"

func TestDogSpeak(t *testing.T) {
	d := Dog{Name: "Rex"}
	result := MakeSpeak(&d)
	expected := "Rex says Woof!"
	if result != expected {
		t.Errorf("Dog.Speak() = %q, want %q", result, expected)
	}
}

func TestCatSpeak(t *testing.T) {
	c := Cat{Name: "Whiskers"}
	result := MakeSpeak(&c)
	expected := "Whiskers says Meow!"
	if result != expected {
		t.Errorf("Cat.Speak() = %q, want %q", result, expected)
	}
}
