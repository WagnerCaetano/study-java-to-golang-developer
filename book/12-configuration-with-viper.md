# Chapter 12: Configuration with Viper

> Externalize settings from code. Viper is the de facto configuration library for Go services — YAML, JSON, env vars, and live reload in one package.

---

## Why Viper?

| Concern | Java (Spring Boot) | Go (Viper) |
|---------|-------------------|------------|
| File config | `application.yml` + `@ConfigurationProperties` | YAML/JSON/TOML via Viper |
| Env overrides | `SPRING_APPLICATION_JSON`, relaxed binding | `AutomaticEnv` + `SetEnvKeyReplacer` |
| Profiles / layers | `application-dev.yml`, `application-prod.yml` | Merge multiple config files |
| Hot reload | Spring Cloud Config + `@RefreshScope` | `WatchConfig` + `OnConfigChange` (fsnotify) |
| Typed binding | `@ConfigurationProperties` POJOs | `Unmarshal` into structs |

In Java, Spring binds properties at startup (and optionally refreshes beans). In Go, you typically load config once at startup, optionally watch the file, and swap an in-memory snapshot — no container magic required.

---

## Setup

```bash
go get github.com/spf13/viper
```

Viper uses [fsnotify](https://github.com/fsnotify/fsnotify) internally when you call `WatchConfig()`.

---

## Example Files

### `config/config.yaml` (defaults and structure)

```yaml
server:
  host: "0.0.0.0"
  port: 8080
  read_timeout: 10s
  write_timeout: 10s

database:
  host: localhost
  port: 5432
  name: bff
  max_open_conns: 25

logging:
  level: info
  format: json

features:
  new_checkout: false
```

### `config/config.json` (environment-specific overlay)

```json
{
  "server": {
    "port": 3000
  },
  "database": {
    "host": "postgres.internal"
  },
  "logging": {
    "level": "debug"
  }
}
```

Later files in a merge chain **override** keys from earlier files. Env vars override file values when `AutomaticEnv()` is enabled.

---

## Typed Structs

Map keys to struct tags with `mapstructure` (Viper’s unmarshaller uses the same tags as [mapstructure](https://github.com/go-viper/mapstructure)):

```go
package config

import "time"

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Logging  LoggingConfig  `mapstructure:"logging"`
    Features map[string]bool `mapstructure:"features"`
}

type ServerConfig struct {
    Host         string        `mapstructure:"host"`
    Port         int           `mapstructure:"port"`
    ReadTimeout  time.Duration `mapstructure:"read_timeout"`
    WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type DatabaseConfig struct {
    Host         string `mapstructure:"host"`
    Port         int    `mapstructure:"port"`
    Name         string `mapstructure:"name"`
    MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type LoggingConfig struct {
    Level  string `mapstructure:"level"`
    Format string `mapstructure:"format"`
}
```

**Java comparison:**

```java
@ConfigurationProperties(prefix = "server")
public record ServerProperties(
    String host,
    int port,
    Duration readTimeout,
    Duration writeTimeout
) {}
```

Go has no annotation processor — explicit struct tags replace `@ConfigurationProperties` binding rules.

---

## Loading YAML and JSON

```go
package config

import (
    "fmt"
    "strings"

    "github.com/spf13/viper"
)

func Load(configDir string, env string) (*Config, error) {
    v := viper.New()

    // Base defaults (optional — can also use SetDefault per key)
    v.SetDefault("server.host", "0.0.0.0")
    v.SetDefault("server.port", 8080)

    v.SetConfigType("yaml")
    v.AddConfigPath(configDir)

    // 1) Shared defaults
    v.SetConfigName("config")
    if err := v.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("read base config: %w", err)
    }

    // 2) Merge environment overlay (JSON example)
    if env != "" {
        overlay := viper.New()
        overlay.SetConfigType("json")
        overlay.AddConfigPath(configDir)
        overlay.SetConfigName("config." + env)
        if err := overlay.ReadInConfig(); err == nil {
            if err := v.MergeConfigMap(overlay.AllSettings()); err != nil {
                return nil, fmt.Errorf("merge env overlay: %w", err)
            }
        }
        // Missing overlay file is OK for local dev — skip silently or log
    }

    // 3) Environment variables override files
    v.SetEnvPrefix("BFF")
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    v.AutomaticEnv()

    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, fmt.Errorf("unmarshal config: %w", err)
    }
    return &cfg, nil
}
```

With `BFF_SERVER_PORT=9090`, Viper maps `server.port` from the env var (prefix + replacer turns `server.port` → `BFF_SERVER_PORT`).

**Merge order (last wins):** defaults → `config.yaml` → `config.{env}.json` → environment variables.

---

## Environment Overrides Only (12-Factor Style)

For containers, you often skip files and bind env directly:

```go
func LoadFromEnv() (*Config, error) {
    v := viper.New()
    v.SetEnvPrefix("BFF")
    v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
    v.AutomaticEnv()

    // Tell Viper which keys exist (env-only apps still need defaults)
    v.SetDefault("server.port", 8080)

    var cfg Config
    if err := v.Unmarshal(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
```

This mirrors Spring’s externalized configuration without shipping YAML inside the image.

---

## Hot Reload with fsnotify

Viper’s `WatchConfig()` registers a filesystem watcher. On change, Viper re-reads the file and fires `OnConfigChange`.

```go
package config

import (
    "log/slog"
    "sync/atomic"

    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

type HotConfig struct {
    v   *viper.Viper
    ptr atomic.Pointer[Config]
}

func NewHot(configDir string) (*HotConfig, error) {
    v := viper.New()
    v.SetConfigName("config")
    v.SetConfigType("yaml")
    v.AddConfigPath(configDir)

    if err := v.ReadInConfig(); err != nil {
        return nil, err
    }

    hc := &HotConfig{v: v}
    if err := hc.reload(); err != nil {
        return nil, err
    }

    v.WatchConfig()
    v.OnConfigChange(func(_ fsnotify.Event) {
        if err := hc.reload(); err != nil {
            slog.Error("config reload failed", "error", err)
            return
        }
        slog.Info("config reloaded")
    })

    return hc, nil
}

func (hc *HotConfig) reload() error {
    var cfg Config
    if err := hc.v.Unmarshal(&cfg); err != nil {
        return err
    }
    hc.ptr.Store(&cfg)
    return nil
}

func (hc *HotConfig) Current() *Config {
    return hc.ptr.Load()
}
```

**Java comparison:** Spring Cloud Config pushes or polls remote config; `@RefreshScope` recreates proxied beans. In Go, **you** decide what reloads: log level, feature flags, HTTP client timeouts — not every struct needs to be hot-swappable. Prefer atomic pointer swaps for read-mostly config (see Chapter 16).

**Caveats:**

- Editors that save via “atomic rename” usually work; some tools break write-then-rename — test your deployment path.
- Reload does not retroactively change already-open DB pools; document which fields require restart.
- Debounce rapid saves in production if operators edit files by hand.

---

## Wiring Config into a Fiber BFF

```go
func main() {
    cfg, err := config.Load("./config", os.Getenv("APP_ENV"))
    if err != nil {
        log.Fatal(err)
    }

    app := fiber.New(fiber.Config{
        ReadTimeout:  cfg.Server.ReadTimeout,
        WriteTimeout: cfg.Server.WriteTimeout,
    })

    addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
    log.Fatal(app.Listen(addr))
}
```

Pass `*config.Config` (or interfaces) into handlers — avoid global `viper.Get*` in request paths; it hides dependencies and complicates tests.

---

## Testing

```go
func TestLoad_mergesEnvOverlay(t *testing.T) {
    dir := t.TempDir()
    writeFile(t, dir, "config.yaml", "server:\n  port: 8080\n")
    writeFile(t, dir, "config.test.json", `{"server":{"port":3001}}`)

    cfg, err := config.Load(dir, "test")
    require.NoError(t, err)
    assert.Equal(t, 3001, cfg.Server.Port)
}
```

Use `t.Setenv("BFF_SERVER_PORT", "9999")` to assert env wins over files.

---

## Key Takeaways

1. **Viper centralizes** file formats, defaults, env binding, and optional file watch in one API.
2. **Merge layers explicitly** — base YAML + env JSON + env vars; document precedence for operators.
3. **`Unmarshal` into structs** — typed config like Spring `@ConfigurationProperties`, with `mapstructure` tags.
4. **Hot reload is opt-in** — use `WatchConfig` + atomic snapshot swaps; not every setting should reload live.
5. **Inject config, don’t scatter globals** — constructor injection keeps Fiber handlers testable.
6. **Validate after load** — Viper unmarshals shapes, not business rules; add a `Validate()` method on `Config`.

---

**Practice**: [Module 09 — Environment Config →](../exercises/module-09-viper-config/54-environment-config/) | [Viper Hot Reload →](../exercises/module-09-viper-config/55-hot-reload-configmap/) | [Module 07 — Timeouts & Health →](../exercises/module-07-resilience/)
