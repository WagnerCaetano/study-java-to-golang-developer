package dynamoquery

import (
	"context"
	"testing"
)

func TestQueryByPK(t *testing.T) {
	ctx := context.Background()
	store := NewInMemoryStore()
	if err := store.Put(ctx, Item{PK: "ORDER", SK: "1", Data: "a"}); err != nil {
		t.Fatal(err)
	}
	if err := store.Put(ctx, Item{PK: "ORDER", SK: "2", Data: "b"}); err != nil {
		t.Fatal(err)
	}
	if err := store.Put(ctx, Item{PK: "USER", SK: "1", Data: "c"}); err != nil {
		t.Fatal(err)
	}

	items, err := QueryByPK(ctx, store, "ORDER")
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 {
		t.Fatalf("len = %d, want 2", len(items))
	}
}
