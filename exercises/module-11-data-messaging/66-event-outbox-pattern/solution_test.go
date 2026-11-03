package outbox

import (
	"context"
	"testing"
)

func TestOutboxPublishPending(t *testing.T) {
	ctx := context.Background()
	repo := NewInMemoryOutboxRepository()
	var pub RecordingPublisher
	svc := &OutboxService{Repo: repo, Pub: &pub}

	if err := svc.SaveEvent(ctx, OutboxEvent{ID: "1", Payload: "order-created"}); err != nil {
		t.Fatal(err)
	}
	if err := svc.PublishPending(ctx, "arn:aws:sns:1:2:events"); err != nil {
		t.Fatal(err)
	}
	if len(pub.Messages) != 1 {
		t.Fatalf("published = %d", len(pub.Messages))
	}
	if pub.Messages[0].Message != "order-created" {
		t.Fatalf("payload = %q", pub.Messages[0].Message)
	}
	pending, _ := repo.ListPending(ctx)
	if len(pending) != 0 {
		t.Fatalf("still pending: %+v", pending)
	}
}
