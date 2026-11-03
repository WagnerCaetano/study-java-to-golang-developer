package hotreload

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Message string `mapstructure:"message"`
}

type ConfigWatcher struct {
	mu      sync.RWMutex
	current AppConfig
	v       *viper.Viper
}

func NewConfigWatcher(path string, onChange func(AppConfig)) (*ConfigWatcher, error) {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	w := &ConfigWatcher{v: v}
	if err := w.apply(onChange); err != nil {
		return nil, err
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		_ = w.apply(onChange)
	})
	return w, nil
}

func (w *ConfigWatcher) apply(onChange func(AppConfig)) error {
	var cfg AppConfig
	if err := w.v.Unmarshal(&cfg); err != nil {
		return err
	}
	w.mu.Lock()
	w.current = cfg
	w.mu.Unlock()
	if onChange != nil {
		onChange(cfg)
	}
	return nil
}

func (w *ConfigWatcher) Current() AppConfig {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.current
}

func (w *ConfigWatcher) Close() error {
	return nil
}
