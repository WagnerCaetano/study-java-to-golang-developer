package dynamoquery

import (
	"context"
	"sync"
)

// Item is a DynamoDB-style record.
type Item struct {
	PK   string
	SK   string
	Data string
}

// ItemStore supports put and query-by-partition-key (Query API shape).
type ItemStore interface {
	Put(ctx context.Context, item Item) error
	QueryByPK(ctx context.Context, pk string) ([]Item, error)
}

// InMemoryStore is a thread-safe fake for tests (maps to DynamoDB Query in production).
type InMemoryStore struct {
	mu    sync.RWMutex
	items map[string]Item
}

// NewInMemoryStore creates an empty store.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{items: make(map[string]Item)}
}

func itemKey(pk, sk string) string { return pk + "|" + sk }

// Put stores an item.
// FIX ME: lock and store under itemKey
func (s *InMemoryStore) Put(ctx context.Context, item Item) error {
	_ = ctx
	return nil // FIX: actually store the item
}

// QueryByPK returns all items with the given partition key.
// FIX ME: scan items where item.PK == pk
func (s *InMemoryStore) QueryByPK(ctx context.Context, pk string) ([]Item, error) {
	_ = ctx
	_ = pk
	return nil, nil
}

// QueryByPK helper delegates to any ItemStore (BFF service layer).
func QueryByPK(ctx context.Context, store ItemStore, pk string) ([]Item, error) {
	return store.QueryByPK(ctx, pk)
}
