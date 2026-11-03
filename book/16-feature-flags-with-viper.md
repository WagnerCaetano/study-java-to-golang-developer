# Chapter 16: Feature Flags with Viper

> Ship code dark, enable behavior from config, flip flags without redeploying. Viper plus hot reload gives you a lightweight alternative to Spring Cloud Config or Unleash — with Go-idiomatic thread safety.

---

## Feature Flags in Java vs Go

| Approach | Java ecosystem | Go + Viper |
|----------|----------------|------------|
| Central config server | Spring Cloud Config | Files, ConfigMaps, or S3 synced to disk |
| Client library | Unleash, LaunchDarkly SDK | Viper keys + custom `Flags` type |
| Hot reload | `@RefreshScope`, `/actuator/refresh` | `WatchConfig` + atomic snapshot |
| Request gating | `@ConditionalOnProperty`, filters | Fiber middleware |
| Safe defaults | `matchIfMissing = false` | Explicit defaults in code before merge |

**Rule:** Defaults in **code** should be the safe production-off path; YAML can enable in staging. Never default risky features to `true`.

---

## Config Shape

### `config/features.yaml`

```yaml
features:
  new_checkout_flow: false
  experimental_search: true
  maintenance_mode: false

rollouts:
  new_checkout_flow: 25   # optional percentage 0-100
```

### `config/features.json` (override)

```json
{
  "features": {
    "experimental_search": false
  },
  "rollouts": {
    "new_checkout_flow": 50
  }
}
```

---

## Typed Flag Snapshot

```go
package flags

type Snapshot struct {
    Features map[string]bool `mapstructure:"features"`
    Rollouts map[string]int  `mapstructure:"rollouts"` // 0-100
}

func (s *Snapshot) Enabled(name string) bool {
    if s == nil {
        return false
    }
    return s.Features[name]
}

func (s *Snapshot) RolloutPercent(name string) int {
    if s == nil || s.Rollouts == nil {
        return 0
    }
    p := s.Rollouts[name]
    if p < 0 {
        return 0
    }
    if p > 100 {
        return 100
    }
    return p
}
```

---

## Thread-Safe Store (atomic.Value)

Readers on the Fiber hot path must not take locks for long. Swap immutable snapshots:

```go
package flags

import (
    "hash/fnv"
    "log/slog"
    "sync/atomic"

    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

type Store struct {
    v    *viper.Viper
    snap atomic.Pointer[Snapshot]
}

func NewStore(configPath string) (*Store, error) {
    v := viper.New()
    v.SetConfigFile(configPath)
    if err := v.ReadInConfig(); err != nil {
        return nil, err
    }

    s := &Store{v: v}
    if err := s.reload(); err != nil {
        return nil, err
    }

    v.WatchConfig()
    v.OnConfigChange(func(e fsnotify.Event) {
        if err := s.reload(); err != nil {
            slog.Error("feature flags reload failed", "error", err, "file", e.Name)
        } else {
            slog.Info("feature flags reloaded")
        }
    })
    return s, nil
}

func (s *Store) reload() error {
    // Safe defaults merged before unmarshal
    s.v.SetDefault("features.new_checkout_flow", false)
    s.v.SetDefault("features.experimental_search", false)
    s.v.SetDefault("features.maintenance_mode", false)

    var snap Snapshot
    if err := s.v.Unmarshal(&snap); err != nil {
        return err
    }
    if snap.Features == nil {
        snap.Features = map[string]bool{}
    }
    s.snap.Store(&snap)
    return nil
}

func (s *Store) Snapshot() *Snapshot {
    return s.snap.Load()
}
```

### Alternative: `sync.RWMutex`

If reload merges complex state you mutate incrementally, guard with `RWMutex`:

```go
import "sync"

type MutexStore struct {
    mu   sync.RWMutex
    snap Snapshot
}

func (m *MutexStore) Enabled(name string) bool {
    m.mu.RLock()
    defer m.mu.RUnlock()
    return m.snap.Enabled(name)
}
```

Prefer `atomic.Pointer` when reload **replaces** the whole struct — no partial reads.

---

## Boolean Flags vs Percentage Rollouts

Simple rollout: hash a stable user id into 0–99 and compare to configured percent:

```go
func (s *Store) EnabledForUser(flag, userID string) bool {
    snap := s.Snapshot()
    if !snap.Enabled(flag) && snap.RolloutPercent(flag) == 0 {
        return false
    }
    if snap.Enabled(flag) && snap.RolloutPercent(flag) == 0 {
        return true // fully on via features map
    }
    pct := snap.RolloutPercent(flag)
    if pct >= 100 {
        return true
    }
    if userID == "" {
        return false
    }
    h := fnv32(userID + ":" + flag)
    return int(h%100) < pct
}

func fnv32(s string) uint32 {
    h := fnv.New32a()
    _, _ = h.Write([]byte(s))
    return h.Sum32()
}
```

**Caveats:**

- Percentage rollouts are **sticky per user** only if `userID` is stable.
- Changing from 25% → 50% reshuffles membership — acceptable for experiments, not for billing logic.
- For enterprise-grade targeting, integrate Unleash/LaunchDarkly; Viper flags suit ops toggles and gradual internal rollouts.

**Java / Unleash style:** server evaluates strategies (`userWithId`, `gradualRollout`); here you encode minimal rules in-process.

---

## Fiber Middleware Gating

### Global maintenance mode

```go
func MaintenanceMiddleware(flags *flags.Store) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if flags.Snapshot().Enabled("maintenance_mode") {
            return c.Status(503).JSON(fiber.Map{
                "error": "maintenance",
            })
        }
        return c.Next()
    }
}
```

### Per-route feature

```go
func RequireFeature(flags *flags.Store, name string) fiber.Handler {
    return func(c *fiber.Ctx) error {
        if !flags.Snapshot().Enabled(name) {
            return c.Status(404).JSON(fiber.Map{"error": "not found"})
        }
        return c.Next()
    }
}

// Registration
app.Post("/checkout/v2",
    RequireFeature(flagStore, "new_checkout_flow"),
    checkoutV2Handler,
)
```

Return **404** when hiding unfinished APIs (security through obscurity is not enough — still authZ). Return **403** when the user lacks entitlement.

### Handler-level check with rollout

```go
func searchHandler(flags *flags.Store) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userID := c.Get("X-User-ID")
        if flags.EnabledForUser("experimental_search", userID) {
            return experimentalSearch(c)
        }
        return legacySearch(c)
    }
}
```

---

## Env Overrides for Emergencies

Combine with Chapter 12 patterns:

```go
v.SetEnvPrefix("BFF")
v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
v.AutomaticEnv()
// BFF_FEATURES_MAINTENANCE_MODE=true kills traffic quickly
```

Document emergency vars in runbooks — operators may not edit YAML during incidents.

---

## Testing

```go
func TestEnabledForUser_rollout(t *testing.T) {
    snap := &Snapshot{
        Features: map[string]bool{"new_checkout_flow": false},
        Rollouts: map[string]int{"new_checkout_flow": 50},
    }
    var store Store
    store.snap.Store(snap)

    enabled := 0
    for i := 0; i < 1000; i++ {
        id := fmt.Sprintf("user-%d", i)
        if store.EnabledForUser("new_checkout_flow", id) {
            enabled++
        }
    }
    // Roughly half — statistical test with tolerance
    assert.InDelta(t, 500, enabled, 120)
}
```

Table-test middleware with `app.Test()` from Fiber — inject a `Store` preloaded with a snapshot.

---

## Operational Guidance

1. **Log flag evaluations at debug** — not every request at info (cardinality).
2. **Metric counters** per flag (`bff_feature_invocation_total{flag="..."}`) — see Chapter 13.
3. **Reload failures keep last good snapshot** — never zero out flags on parse error.
4. **Version config** — git-track `features.yaml`; CI validates schema.

---

## Key Takeaways

1. **Safe defaults in code** — off unless config explicitly enables.
2. **Hot reload via Viper + fsnotify** — swap snapshots with `atomic.Pointer`.
3. **Middleware for coarse gates** — maintenance mode, route-level toggles.
4. **Handler checks for nuanced rules** — percentage rollouts need stable user keys.
5. **Not a full FF platform** — graduate to Unleash/LD when you need audit trails and UI.
6. **Same patterns as Spring Cloud Config** — without bean refresh; you control read paths.

---

**Practice**: [Module 09 — Feature Flags →](../exercises/module-09-viper-config/56-feature-flags/) | [Viper Hot Reload →](../exercises/module-09-viper-config/55-hot-reload-configmap/) | [Module 06 — Fiber Middleware →](../exercises/module-06-bff-fiber/39-fiber-middleware/)
