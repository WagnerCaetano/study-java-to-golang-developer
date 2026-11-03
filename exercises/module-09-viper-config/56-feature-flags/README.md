# Exercise 56: Feature Flags (Viper)

## Problem

**Feature flags** let you ship code dark and enable it per environment. Load boolean flags from a Viper config map (YAML/JSON hot reload in production) and query them at runtime.

Package **`featureflags`**:

- `FeatureFlags` with `Enabled(name string) bool` — missing names are **false**.
- `NewFeatureFlagsFromMap(map[string]bool) *FeatureFlags` — initial flags.
- `SetFlags(map[string]bool)` — replace flags at runtime (simulates Viper `OnConfigChange` updating a `flags` key).

Optional: in your own app, bind Viper key `flags` to a `map[string]any` and convert booleans.

## Learning objectives

- Safe defaults (fail closed for unknown flags).
- Hot-updated toggles without redeploy.
- Same Viper YAML/JSON reload pattern as Exercise 55.

## How to test

```bash
go test -v
```
