package outbox

import (
	"context"
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
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events[ev.ID] = ev
	return nil
}

func (r *InMemoryOutboxRepository) ListPending(ctx context.Context) ([]OutboxEvent, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var out []OutboxEvent
	for _, ev := range r.events {
		if !ev.Published {
			out = append(out, ev)
		}
	}
	return out, nil
}

func (r *InMemoryOutboxRepository) MarkPublished(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	ev, ok := r.events[id]
	if !ok {
		return nil
	}
	ev.Published = true
	r.events[id] = ev
	return nil
}

type OutboxService struct {
	Repo OutboxRepository
	Pub  Publisher
}

func (s *OutboxService) SaveEvent(ctx context.Context, ev OutboxEvent) error {
	return s.Repo.SaveEvent(ctx, ev)
}

func (s *OutboxService) PublishPending(ctx context.Context, topicARN string) error {
	pending, err := s.Repo.ListPending(ctx)
	if err != nil {
		return err
	}
	for _, ev := range pending {
		if err := s.Pub.Publish(ctx, topicARN, ev.Payload); err != nil {
			return err
		}
		if err := s.Repo.MarkPublished(ctx, ev.ID); err != nil {
			return err
		}
	}
	return nil
}
