package pointers

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
