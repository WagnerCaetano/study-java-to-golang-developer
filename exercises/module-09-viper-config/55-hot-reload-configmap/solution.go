package hotreload

import (
	"sync"

	"github.com/spf13/viper"
)

var _ = viper.New // remove when NewConfigWatcher is implemented

// AppConfig is the runtime configuration loaded from file.
type AppConfig struct {
	Message string `mapstructure:"message"`
}

// ConfigWatcher watches a config file and keeps the latest AppConfig.
type ConfigWatcher struct {
	mu      sync.RWMutex
	current AppConfig
	// FIX ME: hold viper instance and cancel watch if needed
}

// NewConfigWatcher loads path and invokes onChange whenever the file changes.
func NewConfigWatcher(path string, onChange func(AppConfig)) (*ConfigWatcher, error) {
	// FIX ME: use github.com/spf13/viper — ReadInConfig, Unmarshal, WatchConfig, OnConfigChange
	_ = path
	_ = onChange
	return &ConfigWatcher{}, nil
}

// Current returns the latest configuration.
func (w *ConfigWatcher) Current() AppConfig {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.current
}

// Close stops watching the config file.
func (w *ConfigWatcher) Close() error {
	return nil
}
