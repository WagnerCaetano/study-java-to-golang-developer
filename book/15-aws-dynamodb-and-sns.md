# Chapter 15: AWS DynamoDB & SNS with SDK v2

> Many BFFs read denormalized views from DynamoDB and publish domain events to SNS. AWS SDK for Go v2 mirrors Java v2: modular clients, context-aware APIs, and the default credential chain.

---

## Java vs Go AWS SDK v2

| Concept | Java (`software.amazon.awssdk`) | Go (`github.com/aws/aws-sdk-go-v2`) |
|---------|--------------------------------|-------------------------------------|
| Client builder | `DynamoDbClient.builder().region(...).build()` | `dynamodb.NewFromConfig(cfg)` |
| Credentials | Default credentials provider chain | `config.LoadDefaultConfig` |
| Pagination | `PaginatedIterable` | `Paginator` types |
| Testing | Interface wrappers + Mockito | Small interfaces you define |
| Context | Less idiomatic in older SDK | Every API call: `context.Context` |

Spring Cloud AWS auto-wires clients. In Go, construct clients in `main` and inject interfaces into handlers — explicit and test-friendly.

---

## Setup

```bash
go get github.com/aws/aws-sdk-go-v2/config
go get github.com/aws/aws-sdk-go-v2/service/dynamodb
go get github.com/aws/aws-sdk-go-v2/service/sns
go get github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue
```

---

## Loading Config (Region & Credentials)

The **default credential chain** checks, in order: env vars (`AWS_ACCESS_KEY_ID`, …), shared config/credentials files, SSO, and **ECS/EKS/EC2 IAM roles** — same idea as Java’s `DefaultCredentialsProvider`.

```go
package awsx

import (
    "context"
    "os"

    "github.com/aws/aws-sdk-go-v2/aws"
    awscfg "github.com/aws/aws-sdk-go-v2/config"
)

func LoadAWSConfig(ctx context.Context) (aws.Config, error) {
    region := os.Getenv("AWS_REGION")
    if region == "" {
        region = "us-east-1"
    }
    return awscfg.LoadDefaultConfig(ctx,
        awscfg.WithRegion(region),
        // Optional: custom profile for local dev
        // awscfg.WithSharedConfigProfile("dev"),
    )
}
```

LocalStack / dev endpoints:

```go
func LoadLocalConfig(ctx context.Context) (aws.Config, error) {
    cfg, err := awscfg.LoadDefaultConfig(ctx, awscfg.WithRegion("us-east-1"))
    if err != nil {
        return cfg, err
    }
    cfg.BaseEndpoint = aws.String("http://localhost:4566")
    return cfg, nil
}
```

---

## Interface-Based Clients (Testability)

AWS Go clients are concrete structs with many methods. Define **narrow interfaces** for what your BFF uses:

```go
package store

import (
    "context"

    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoAPI interface {
    GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
    PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
    Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

type UserViewStore struct {
    table string
    db    DynamoAPI
}

func NewUserViewStore(table string, db DynamoAPI) *UserViewStore {
    return &UserViewStore{table: table, db: db}
}
```

`*dynamodb.Client` satisfies `DynamoAPI` automatically. Tests pass a fake or [localstack](https://localstack.cloud/).

**Java comparison:**

```java
public interface UserViewRepository {
    Optional<UserView> findById(String id);
}
// Production: DynamoDbClient injected
// Test: in-memory fake
```

---

## DynamoDB: GetItem

Table design (single-table style example):

| PK | SK | attributes |
|----|-----|------------|
| `USER#123` | `PROFILE` | name, email |
| `USER#123` | `ORDER#456` | total, status |

```go
import (
    "context"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserProfile struct {
    PK    string `dynamodbav:"PK"`
    SK    string `dynamodbav:"SK"`
    Name  string `dynamodbav:"name"`
    Email string `dynamodbav:"email"`
}

func (s *UserViewStore) GetProfile(ctx context.Context, userID string) (*UserProfile, error) {
    out, err := s.db.GetItem(ctx, &dynamodb.GetItemInput{
        TableName: aws.String(s.table),
        Key: map[string]types.AttributeValue{
            "PK": &types.AttributeValueMemberS{Value: "USER#" + userID},
            "SK": &types.AttributeValueMemberS{Value: "PROFILE"},
        },
    })
    if err != nil {
        return nil, fmt.Errorf("dynamodb get: %w", err)
    }
    if out.Item == nil {
        return nil, ErrNotFound
    }
    var p UserProfile
    if err := attributevalue.UnmarshalMap(out.Item, &p); err != nil {
        return nil, err
    }
    return &p, nil
}
```

Use `context.WithTimeout` in Fiber handlers — DynamoDB calls should not exceed upstream SLA.

---

## DynamoDB: PutItem

```go
func (s *UserViewStore) SaveProfile(ctx context.Context, userID string, name, email string) error {
    item, err := attributevalue.MarshalMap(UserProfile{
        PK:    "USER#" + userID,
        SK:    "PROFILE",
        Name:  name,
        Email: email,
    })
    if err != nil {
        return err
    }

    _, err = s.db.PutItem(ctx, &dynamodb.PutItemInput{
        TableName: aws.String(s.table),
        Item:      item,
        // ConditionExpression: aws.String("attribute_not_exists(PK)"), // create-only
    })
    return err
}
```

**Idempotency:** BFF writes to views are often upserts; use `ConditionExpression` when concurrent writers matter.

---

## DynamoDB: Query Basics

Query requires a partition key; use `KeyConditionExpression`:

```go
func (s *UserViewStore) ListUserOrders(ctx context.Context, userID string) ([]OrderSummary, error) {
    pk := "USER#" + userID
    out, err := s.db.Query(ctx, &dynamodb.QueryInput{
        TableName:              aws.String(s.table),
        KeyConditionExpression: aws.String("PK = :pk AND begins_with(SK, :prefix)"),
        ExpressionAttributeValues: map[string]types.AttributeValue{
            ":pk":     &types.AttributeValueMemberS{Value: pk},
            ":prefix": &types.AttributeValueMemberS{Value: "ORDER#"},
        },
    })
    if err != nil {
        return nil, err
    }

    var orders []OrderSummary
    if err := attributevalue.UnmarshalListOfMaps(out.Items, &orders); err != nil {
        return nil, err
    }
    return orders, nil
}
```

For large partitions, loop with `LastEvaluatedKey` or use `NewQueryPaginator`.

---

## SNS: Publish

```go
package events

import (
    "context"
    "encoding/json"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSAPI interface {
    Publish(ctx context.Context, params *sns.PublishInput, optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type Publisher struct {
    topicArn string
    sns      SNSAPI
}

func NewPublisher(topicArn string, sns SNSAPI) *Publisher {
    return &Publisher{topicArn: topicArn, sns: sns}
}

type UserUpdatedEvent struct {
    UserID string `json:"user_id"`
    Source string `json:"source"`
}

func (p *Publisher) UserUpdated(ctx context.Context, e UserUpdatedEvent) error {
    body, err := json.Marshal(e)
    if err != nil {
        return err
    }
    _, err = p.sns.Publish(ctx, &sns.PublishInput{
        TopicArn: aws.String(p.topicArn),
        Message:  aws.String(string(body)),
        MessageAttributes: map[string]types.MessageAttributeValue{
            "event_type": {
                DataType:    aws.String("String"),
                StringValue: aws.String("user.updated"),
            },
        },
    })
    return err
}
```

Add import `"github.com/aws/aws-sdk-go-v2/service/sns/types"` for `MessageAttributeValue`.

**Java comparison:** `SnsClient.publish(PublishRequest.builder()...build())` — same Publish API; Go returns errors, not exceptions.

Fan-out: SNS → SQS/Lambda subscribers; BFF only publishes after successful origin mutation.

---

## Wiring in `main`

```go
func main() {
    ctx := context.Background()
    cfg, err := awsx.LoadAWSConfig(ctx)
    if err != nil {
        log.Fatal(err)
    }

    ddb := dynamodb.NewFromConfig(cfg)
    snsClient := sns.NewFromConfig(cfg)

    store := store.NewUserViewStore(os.Getenv("DDB_TABLE"), ddb)
    pub := events.NewPublisher(os.Getenv("SNS_TOPIC_ARN"), snsClient)

    app := fiber.New()
    h := handlers.New(store, pub)
    app.Get("/users/:id", h.GetUser)
    app.Listen(":8080")
}
```

---

## BFF Usage Patterns

1. **Read path:** Query DynamoDB view → merge with live data from HTTP services → return aggregated JSON.
2. **Write path:** Call authoritative service → on success, publish SNS event for downstream projections (or dual-write carefully with idempotency keys).
3. **Never expose raw AWS errors** to clients — map to 502/503 with correlation IDs.

---

## Key Takeaways

1. **`config.LoadDefaultConfig`** — region + credential chain, equivalent to Java SDK v2 defaults.
2. **Define small interfaces** — test BFF logic without LocalStack when possible.
3. **GetItem / PutItem / Query** cover most BFF read models; learn partition key design first.
4. **attributevalue** marshaling — struct tags `dynamodbav` like Jackson annotations for DDB.
5. **SNS Publish** for async notifications — keep messages small, version your event schema.
6. **Context timeouts on every AWS call** — shared with Redis and HTTP downstream rules.

---

**Practice**: [DynamoDB Put/Get →](../exercises/module-11-data-messaging/63-dynamodb-put-get/) | [SNS Publish →](../exercises/module-11-data-messaging/65-sns-publish/) | [Outbox Pattern →](../exercises/module-11-data-messaging/66-event-outbox-pattern/)
