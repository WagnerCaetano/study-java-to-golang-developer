package mapiter

import "sort"

// SortedKeys returns the keys of a map[string]int in alphabetical order
// FIX ME: Extract keys and sort them
func SortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// FIX: Extract keys, sort them, return
	return keys
}
