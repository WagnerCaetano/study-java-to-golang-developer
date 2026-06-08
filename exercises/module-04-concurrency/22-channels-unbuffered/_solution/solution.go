package channel

func Producer(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch <-chan int) int {
	total := 0
	for val := range ch {
		total += val
	}
	return total
}
