package variables

type User struct {
	Name string
	Age  int
}

// CreateUser should return a User with the given name and age
// FIX ME: Initialize the User struct properly
func CreateUser(name string, age int) User {
	return User{} // FIX: set Name and Age fields
}
