package channel

// Producer generates numbers 1 to n and sends them on the channel
// FIX ME: Close the channel after sending all numbers
func Producer(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		// FIX: close the channel
	}()
	return ch
}

// Consumer reads all values from the channel and returns their sum
func Consumer(ch <-chan int) int {
	total := 0
	for val := range ch {
		total += val
	}
	return total
}
