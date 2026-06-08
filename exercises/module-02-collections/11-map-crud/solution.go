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
	// FIX: store in the data map
}

// Get retrieves a value by key. Returns ("", false) if not found.
// FIX ME: Implement this
func (s *Store) Get(key string) (string, bool) {
	return "", false // FIX
}

// Delete removes a key from the store
// FIX ME: Implement this
func (s *Store) Delete(key string) {
	// FIX: delete from data map
}
