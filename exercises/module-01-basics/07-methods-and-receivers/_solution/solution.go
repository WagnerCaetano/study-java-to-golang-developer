package methods

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) Value() int {
	return c.value
}

func (c *Counter) Reset() {
	c.value = 0
}
