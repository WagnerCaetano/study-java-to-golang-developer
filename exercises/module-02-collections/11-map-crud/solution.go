package mapcrud

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}

// Set stores a key-value pair
// FIX ME: Implement this
func (s *Store) Set(key, value string) {
	s.data[key] = value
}

// Get retrieves a value by key. Returns ("", false) if not found.
// FIX ME: Implement this
func (s *Store) Get(key string) (string, bool) {
	result, ok := s.data[key]
	if ok {
		return result, true
	}
	return "", false // FIX
}

// Delete removes a key from the store
// FIX ME: Implement this
func (s *Store) Delete(key string) {
	delete(s.data, key)
}
