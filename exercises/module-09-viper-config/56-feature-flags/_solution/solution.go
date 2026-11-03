package featureflags

import "sync"

type FeatureFlags struct {
	mu    sync.RWMutex
	flags map[string]bool
}

func NewFeatureFlagsFromMap(m map[string]bool) *FeatureFlags {
	cp := make(map[string]bool, len(m))
	for k, v := range m {
		cp[k] = v
	}
	return &FeatureFlags{flags: cp}
}

func (f *FeatureFlags) Enabled(name string) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.flags[name]
}

func (f *FeatureFlags) SetFlags(m map[string]bool) {
	cp := make(map[string]bool, len(m))
	for k, v := range m {
		cp[k] = v
	}
	f.mu.Lock()
	f.flags = cp
	f.mu.Unlock()
}
