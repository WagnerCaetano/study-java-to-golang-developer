package interfaces

// Speaker is an interface for things that can speak
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// FIX ME: Implement Speak() for Dog so it returns "<Name> says Woof!"
func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// FIX ME: Implement Speak() for Cat so it returns "<Name> says Meow!"
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// MakeSpeak calls Speak on any Speaker
func MakeSpeak(s Speaker) string {
	return s.Speak()
}
