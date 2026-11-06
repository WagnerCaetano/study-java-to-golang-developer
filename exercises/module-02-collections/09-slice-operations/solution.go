package slices

// FilterEven returns only even numbers from the input slice
// FIX ME: Currently returns an empty slice
func FilterEven(nums []int) []int {
	var result []int
	for _, v := range nums {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	// FIX: Iterate over nums and append even numbers to result
	return result
}
