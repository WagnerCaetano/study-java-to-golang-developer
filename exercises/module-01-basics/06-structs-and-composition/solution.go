package structs

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// FullName returns the person's full name
// FIX ME: Should return "FirstName LastName" format
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName // FIX: return proper full name
}
