package mapcrud

import "testing"

func TestStoreCRUD(t *testing.T) {
	s := NewStore()
	
	// Get non-existent
	_, ok := s.Get("missing")
	if ok {
		t.Error("expected missing key to return false")
	}
	
	// Set and Get
	s.Set("name", "Alice")
	val, ok := s.Get("name")
	if !ok || val != "Alice" {
		t.Errorf("Get(name) = %q, %v; want %q, true", val, ok, "Alice")
	}
	
	// Update
	s.Set("name", "Bob")
	val, _ = s.Get("name")
	if val != "Bob" {
		t.Errorf("After update: got %q, want %q", val, "Bob")
	}
	
	// Delete
	s.Delete("name")
	_, ok = s.Get("name")
	if ok {
		t.Error("expected key to be deleted")
	}
}
