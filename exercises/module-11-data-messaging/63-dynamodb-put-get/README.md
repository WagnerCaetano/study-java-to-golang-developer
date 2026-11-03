# Exercise 63: DynamoDB Put / Get

## Problem

Key-value persistence on AWS uses **DynamoDB** [`PutItem`](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html) / [`GetItem`](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html) via [`aws-sdk-go-v2`](https://github.com/aws/aws-sdk-go-v2). Tests use an in-memory repository.

Package **`dynamoex`**:

```go
type Item struct {
    PK string
    SK string
    Data string
}

type ItemRepository interface {
    Put(ctx context.Context, item Item) error
    Get(ctx context.Context, pk, sk string) (Item, error)
}
```

Implement **`InMemoryRepository`** (thread-safe map keyed by `pk|sk`).

Stub **`DynamoDBRepository`** documents wiring to `dynamodb.Client` — students implement `Put`/`Get` using attribute maps or `attributevalue.MarshalMap` in production.

Return `ErrNotFound` when missing.

## Learning objectives

- Repository pattern over cloud SDKs.
- Composite keys (PK/SK) as in single-table design.
- Offline tests with fakes before hitting LocalStack or AWS.

## How to test

```bash
go test -v
```
