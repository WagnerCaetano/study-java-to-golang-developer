package snsex

import (
	"context"
	"testing"
)

func TestRecordingPublisher(t *testing.T) {
	ctx := context.Background()
	var pub RecordingPublisher
	if err := pub.Publish(ctx, "arn:aws:sns:1:2:orders", "created"); err != nil {
		t.Fatal(err)
	}
	if len(pub.Messages) != 1 || pub.Messages[0].Message != "created" {
		t.Fatalf("messages = %+v", pub.Messages)
	}
}
