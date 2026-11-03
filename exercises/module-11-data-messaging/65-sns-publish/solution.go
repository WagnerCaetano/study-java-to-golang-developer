package snsex

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
	return errors.New("not implemented")
}

type SNSPublisher struct {
	TopicARN string
}

func (p *SNSPublisher) Publish(ctx context.Context, topicARN, message string) error {
	return errors.New("wire sns.Client Publish")
}
