package outbox

import (
	"context"
	"errors"
	"sync"
)

type PublishedMessage struct {
	TopicARN string
	Message  string
}

type Publisher interface {
	Publish(ctx context.Context, topicARN, message string) error
}

type RecordingPublisher struct {
	mu       sync.Mutex
	Messages []PublishedMessage
}

func (p *RecordingPublisher) Publish(ctx context.Context, topicARN, message string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.Messages = append(p.Messages, PublishedMessage{TopicARN: topicARN, Message: message})
	return nil
}

type OutboxEvent struct {
	ID        string
	Payload   string
	Published bool
}

type OutboxRepository interface {
	SaveEvent(ctx context.Context, ev OutboxEvent) error
	ListPending(ctx context.Context) ([]OutboxEvent, error)
	MarkPublished(ctx context.Context, id string) error
}

type InMemoryOutboxRepository struct {
	mu     sync.Mutex
	events map[string]OutboxEvent
}

func NewInMemoryOutboxRepository() *InMemoryOutboxRepository {
	return &InMemoryOutboxRepository{events: make(map[string]OutboxEvent)}
}

func (r *InMemoryOutboxRepository) SaveEvent(ctx context.Context, ev OutboxEvent) error {
	_ = ctx
	return errors.New("not implemented")
}

func (r *InMemoryOutboxRepository) ListPending(ctx context.Context) ([]OutboxEvent, error) {
	_ = ctx
	return nil, nil
}

func (r *InMemoryOutboxRepository) MarkPublished(ctx context.Context, id string) error {
	_ = ctx
	_ = id
	return errors.New("not implemented")
}

type OutboxService struct {
	Repo OutboxRepository
	Pub  Publisher
}

func (s *OutboxService) SaveEvent(ctx context.Context, ev OutboxEvent) error {
	return errors.New("not implemented")
}

func (s *OutboxService) PublishPending(ctx context.Context, topicARN string) error {
	_ = topicARN
	return errors.New("not implemented")
}
