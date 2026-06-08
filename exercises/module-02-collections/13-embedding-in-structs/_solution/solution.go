package embed

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	Doors int
}

func NewCar(make, model string, year, doors int) *Car {
	return &Car{
		Vehicle: Vehicle{Make: make, Model: model, Year: year},
		Doors:   doors,
	}
}
