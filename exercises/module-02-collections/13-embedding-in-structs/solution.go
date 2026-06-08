package embed

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

// Car should embed Vehicle and add a Doors field
// FIX ME: Define Car struct with embedded Vehicle and Doors int field
type Car struct {
	// FIX: embed Vehicle and add Doors int
}

// NewCar creates a new Car with all fields set
// FIX ME: Initialize the embedded Vehicle and Doors
func NewCar(make, model string, year, doors int) *Car {
	return nil // FIX
}
