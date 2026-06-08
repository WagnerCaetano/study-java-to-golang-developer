package methods

// Counter keeps track of a count
type Counter struct {
	value int
}

// Increment increases the counter by 1
// FIX ME: This uses a value receiver — it won't modify the original!
func (c Counter) Increment() {
	c.value++
}

// Value returns the current counter value
func (c Counter) Value() int {
	return c.value
}

// Reset sets the counter back to 0
// FIX ME: This also uses a value receiver — it won't modify the original!
func (c Counter) Reset() {
	c.value = 0
}
