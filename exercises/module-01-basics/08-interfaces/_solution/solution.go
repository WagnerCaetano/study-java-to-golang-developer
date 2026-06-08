package interfaces

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d *Dog) Speak() string { return d.Name + " says Woof!" }
func (c *Cat) Speak() string { return c.Name + " says Meow!" }

func MakeSpeak(s Speaker) string { return s.Speak() }
