package variables

type User struct {
	Name string
	Age  int
}

func CreateUser(name string, age int) User {
	return User{Name: name, Age: age}
}
