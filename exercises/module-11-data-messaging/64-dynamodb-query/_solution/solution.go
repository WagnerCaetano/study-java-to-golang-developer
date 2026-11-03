package dynamoquery

import (
	"context"
	"sync"
)

type Item struct {
	PK   string
	SK   string
	Data string
}

type ItemStore interface {
	Put(ctx context.Context, item Item) error
	QueryByPK(ctx context.Context, pk string) ([]Item, error)
}

type InMemoryStore struct {
	mu    sync.RWMutex
	items map[string]Item
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{items: make(map[string]Item)}
}

func itemKey(pk, sk string) string { return pk + "|" + sk }

func (s *InMemoryStore) Put(ctx context.Context, item Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[itemKey(item.PK, item.SK)] = item
	return nil
}

func (s *InMemoryStore) QueryByPK(ctx context.Context, pk string) ([]Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []Item
	for _, item := range s.items {
		if item.PK == pk {
			out = append(out, item)
		}
	}
	return out, nil
}

func QueryByPK(ctx context.Context, store ItemStore, pk string) ([]Item, error) {
	return store.QueryByPK(ctx, pk)
}
