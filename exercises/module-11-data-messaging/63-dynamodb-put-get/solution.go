package dynamoex

import (
	"context"
	"errors"
	"sync"
)

// ErrNotFound is returned when an item does not exist.
var ErrNotFound = errors.New("not found")

// Item is a single DynamoDB record (PK + SK).
type Item struct {
	PK   string
	SK   string
	Data string
}

// ItemRepository persists items.
type ItemRepository interface {
	Put(ctx context.Context, item Item) error
	Get(ctx context.Context, pk, sk string) (Item, error)
}

// InMemoryRepository is a thread-safe fake for tests.
type InMemoryRepository struct {
	mu    sync.RWMutex
	items map[string]Item
}

// NewInMemoryRepository creates an empty store.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{items: make(map[string]Item)}
}

func itemKey(pk, sk string) string { return pk + "|" + sk }

// Put stores item.
func (r *InMemoryRepository) Put(ctx context.Context, item Item) error {
	// FIX ME: store under itemKey(item.PK, item.SK)
	return errors.New("not implemented")
}

// Get loads item by keys.
func (r *InMemoryRepository) Get(ctx context.Context, pk, sk string) (Item, error) {
	// FIX ME: return ErrNotFound when missing
	return Item{}, ErrNotFound
}

// QueryByPK returns all items with the given partition key.
func (r *InMemoryRepository) QueryByPK(ctx context.Context, pk string) ([]Item, error) {
	// FIX ME: scan r.items for matching PK
	return nil, nil
}

// DynamoDBRepository wraps AWS SDK (implement in production).
type DynamoDBRepository struct {
	TableName string
}

func (r *DynamoDBRepository) Put(ctx context.Context, item Item) error {
	return errors.New("wire aws-sdk-go-v2 dynamodb.Client PutItem")
}

func (r *DynamoDBRepository) Get(ctx context.Context, pk, sk string) (Item, error) {
	return Item{}, errors.New("wire aws-sdk-go-v2 dynamodb.Client GetItem")
}
