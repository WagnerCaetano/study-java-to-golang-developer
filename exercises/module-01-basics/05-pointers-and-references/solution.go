package pointers

// Swap should swap the values of a and b using pointers
// FIX ME: The current implementation doesn't actually swap
func Swap(a, b *int) {
	// FIX: Implement proper swap using a temporary variable
	c := *a
	*a = *b
	*b = c
}
