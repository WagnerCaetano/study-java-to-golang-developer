package pooling

import (
	"encoding/json"
	"testing"
)

func TestWriteJSON(t *testing.T) {
	pool := NewBufferPool()
	type payload struct {
		Name string `json:"name"`
		N    int    `json:"n"`
	}
	got, err := WriteJSON(pool, payload{Name: "test", N: 42})
	if err != nil {
		t.Fatalf("WriteJSON: %v", err)
	}
	var decoded payload
	if err := json.Unmarshal(got, &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.Name != "test" || decoded.N != 42 {
		t.Errorf("decoded = %+v", decoded)
	}
}

func TestPutResetsBuffer(t *testing.T) {
	pool := NewBufferPool()
	buf := pool.Get()
	buf.WriteString("leftover")
	pool.Put(buf)

	again := pool.Get()
	if again.Len() != 0 {
		t.Errorf("buffer Len after Get following Put = %d, want 0", again.Len())
	}
	pool.Put(again)
}
