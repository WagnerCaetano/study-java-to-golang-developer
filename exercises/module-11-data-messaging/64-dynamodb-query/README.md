# Exercise 64: dynamodb-query

## Problem
Implement an in-memory store that supports **Query by partition key** (the DynamoDB `Query` API shape). Wire `Put` and `QueryByPK` on `InMemoryStore`, then use the `QueryByPK` helper from a BFF service layer.

In production you would call `aws-sdk-go-v2` DynamoDB `Query` with a `KeyConditionExpression` on the PK — keep that behind the `ItemStore` interface.

## Learning Objectives
- Model DynamoDB Query (all items for a PK)
- Keep AWS SDK behind a small interface for tests
- Thread-safe in-memory fakes

## How to Test
```bash
go test -v
```
