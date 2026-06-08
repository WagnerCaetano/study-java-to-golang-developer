package tabletest

import "errors"

type Calculator struct{}

func (c Calculator) Add(a, b int) int { return a + b }
func (c Calculator) Subtract(a, b int) int { return a - b }
func (c Calculator) Multiply(a, b int) int { return a * b }
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
