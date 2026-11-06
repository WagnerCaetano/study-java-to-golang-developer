package slicecap

// Merge combines two slices into a new one
// FIX ME: Use make to pre-allocate and append to fill
func Merge(a, b []int) []int {
	// FIX: Pre-allocate with make, then append elements
	result := make([]int, 0, len(a)+len(b))
	for _, value := range a {
		result = append(result, value)
	}
	for _, value := range b {
		result = append(result, value)
	}
	return result
}
