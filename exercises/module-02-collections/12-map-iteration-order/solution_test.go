package mapiter

import "testing"

func TestSortedKeys(t *testing.T) {
	m := map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
	}
	result := SortedKeys(m)
	expected := []string{"alice", "bob", "charlie"}
	if len(result) != len(expected) {
		t.Fatalf("len = %d, want %d", len(result), len(expected))
	}
	for i, k := range result {
		if k != expected[i] {
			t.Errorf("key[%d] = %q, want %q", i, k, expected[i])
		}
	}
}
