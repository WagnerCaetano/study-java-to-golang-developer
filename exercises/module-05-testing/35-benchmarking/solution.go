package benchmark

// Concat concatenates strings using + operator (slow for many strings)
// FIX ME: Optimize using strings.Builder
func Concat(parts []string) string {
	result := ""
	for _, p := range parts {
		result += p
	}
	return result
}
