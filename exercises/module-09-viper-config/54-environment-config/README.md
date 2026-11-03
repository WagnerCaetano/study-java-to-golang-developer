# Exercise 54: Environment Config

## Problem

Twelve-factor apps read configuration from the **environment**. Load service settings into a struct using either `os.Getenv` or [Viper](https://github.com/spf13/viper) with `AutomaticEnv()`.

Package **`envconfig`**:

```go
type Config struct {
    Port    int    // env: PORT
    AppName string // env: APP_NAME
}

func LoadFromEnv() (Config, error)
```

Rules:

- `PORT` must parse as a base-10 integer; invalid or missing values return an error.
- `APP_NAME` must be non-empty.

## Learning objectives

- Config structs vs. scattered `getenv` calls.
- Viper env binding (`SetEnvPrefix`, `AutomaticEnv`, `BindEnv`) as an alternative to manual parsing.
- Fail fast on bad config at startup (common in Kubernetes Deployments).

## How to test

```bash
go test -v
```
