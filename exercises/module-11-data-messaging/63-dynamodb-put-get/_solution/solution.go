package dynamoex

import (
	"context"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Item struct {
	PK   string
	SK   string
	Data string
}

type ItemRepository interface {
	Put(ctx context.Context, item Item) error
	Get(ctx context.Context, pk, sk string) (Item, error)
}

type InMemoryRepository struct {
	mu    sync.RWMutex
	items map[string]Item
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{items: make(map[string]Item)}
}

func itemKey(pk, sk string) string { return pk + "|" + sk }

func (r *InMemoryRepository) Put(ctx context.Context, item Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[itemKey(item.PK, item.SK)] = item
	return nil
}

func (r *InMemoryRepository) Get(ctx context.Context, pk, sk string) (Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[itemKey(pk, sk)]
	if !ok {
		return Item{}, ErrNotFound
	}
	return item, nil
}

func (r *InMemoryRepository) QueryByPK(ctx context.Context, pk string) ([]Item, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var out []Item
	for _, item := range r.items {
		if item.PK == pk {
			out = append(out, item)
		}
	}
	return out, nil
}

type DynamoDBRepository struct {
	TableName string
}

func (r *DynamoDBRepository) Put(ctx context.Context, item Item) error {
	return errors.New("wire aws-sdk-go-v2 dynamodb.Client PutItem")
}

func (r *DynamoDBRepository) Get(ctx context.Context, pk, sk string) (Item, error) {
	return Item{}, errors.New("wire aws-sdk-go-v2 dynamodb.Client GetItem")
}
