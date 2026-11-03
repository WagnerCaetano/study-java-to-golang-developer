# Exercise 66: event-outbox-pattern

## Problem
Implement a transactional **outbox**: save events locally, then `PublishPending` to SNS (via a `Publisher` interface). Mark rows published only after a successful publish.

This keeps BFF write paths reliable without dual-writing to DynamoDB and SNS in one brittle call.

## Learning Objectives
- Outbox pattern for reliable messaging
- Interface-based SNS publishing (aws-sdk-go-v2 in production)
- Idempotent-friendly MarkPublished flow

## How to Test
```bash
go test -v
```
