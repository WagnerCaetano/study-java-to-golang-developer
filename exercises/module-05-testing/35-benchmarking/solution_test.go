package benchmark

import "testing"

func TestConcat(t *testing.T) {
	result := Concat([]string{"a", "b", "c"})
	if result != "abc" {
		t.Errorf("Concat = %q, want %q", result, "abc")
	}
}

func BenchmarkConcat(b *testing.B) {
	parts := []string{"hello", " ", "world", "!"}
	for i := 0; i < b.N; i++ {
		Concat(parts)
	}
}
