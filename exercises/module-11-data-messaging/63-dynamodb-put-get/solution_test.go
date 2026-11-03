package dynamoex

import (
	"context"
	"errors"
	"testing"
)

func TestInMemoryPutGet(t *testing.T) {
	ctx := context.Background()
	repo := NewInMemoryRepository()
	item := Item{PK: "USER", SK: "1", Data: "alice"}
	if err := repo.Put(ctx, item); err != nil {
		t.Fatal(err)
	}
	got, err := repo.Get(ctx, "USER", "1")
	if err != nil || got.Data != "alice" {
		t.Fatalf("got = %+v, err = %v", got, err)
	}
	_, err = repo.Get(ctx, "USER", "missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("err = %v", err)
	}
}
