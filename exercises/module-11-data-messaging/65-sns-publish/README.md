# Exercise 65: SNS Publish

## Problem

**Amazon SNS** fan-out messaging uses [`Publish`](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) in **aws-sdk-go-v2** (`github.com/aws/aws-sdk-go-v2/service/sns`). Tests record messages in memory.

Package **`snsex`**:

```go
type Publisher interface {
    Publish(ctx context.Context, topicARN, message string) error
}
```

Implement **`RecordingPublisher`** — append `{TopicARN, Message}` to `Messages` (thread-safe).

Stub **`SNSPublisher`** struct documenting real SDK wiring.

## Learning objectives

- Interface for cloud messaging (mock in unit tests).
- At-least-once delivery semantics (consumers must be idempotent).
- Topic ARN as address for events.

## How to test

```bash
go test -v
```
