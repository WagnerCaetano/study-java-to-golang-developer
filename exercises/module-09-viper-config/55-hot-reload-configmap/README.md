# Exercise 55: Hot Reload Config (Viper)

## Problem

In Kubernetes, a **ConfigMap** mounted as a file can change at runtime. [Viper](https://github.com/spf13/viper) can **watch** the file (via [fsnotify](https://github.com/fsnotify/fsnotify)) and fire callbacks on change — the same pattern works locally with YAML or JSON config files.

Package **`hotreload`**:

1. `AppConfig` with `Message string` (`mapstructure:"message"`).
2. `NewConfigWatcher(path string, onChange func(AppConfig)) (*ConfigWatcher, error)` — load the file, unmarshal into `AppConfig`, call `onChange` with the initial config, then watch for updates.
3. `Current() AppConfig` — latest config (thread-safe).
4. `Close()` — stop watching.

Supported formats: **YAML** and **JSON** (by file extension).

## Learning objectives

- Viper `ReadInConfig`, `Unmarshal`, `WatchConfig`, `OnConfigChange`.
- Hot reload without restarting pods (feature toggles, copy, limits).
- Thread-safe reads while the watcher updates in the background.

## How to test

```bash
go test -v
```
