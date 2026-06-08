package subtests

import "testing"

func TestStringTransform(t *testing.T) {
	t.Run("ToUpper", func(t *testing.T) {
		if ToUpper("hello") != "HELLO" {
			t.Error("ToUpper failed")
		}
	})
	t.Run("ToLower", func(t *testing.T) {
		if ToLower("HELLO") != "hello" {
			t.Error("ToLower failed")
		}
	})
	t.Run("Empty string", func(t *testing.T) {
		if ToUpper("") != "" || ToLower("") != "" {
			t.Error("Empty string handling failed")
		}
	})
}
