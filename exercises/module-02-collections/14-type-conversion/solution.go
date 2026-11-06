package conversion

type Celsius float64
type Fahrenheit float64

// CToF converts Celsius to Fahrenheit
// FIX ME: Formula is F = C * 9/5 + 32
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit((c * (9 / 5)) + 32)
}

// FToC converts Fahrenheit to Celsius
// FIX ME: Formula is C = (F - 32) * 5/9
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
