package featureflags

import "sync"

// FeatureFlags holds named toggles.
type FeatureFlags struct {
	mu    sync.RWMutex
	flags map[string]bool
}

// NewFeatureFlagsFromMap creates flags from a static map.
func NewFeatureFlagsFromMap(m map[string]bool) *FeatureFlags {
	// FIX ME: copy m into flags map
	return &FeatureFlags{}
}

// Enabled reports whether name is on (unknown names → false).
func (f *FeatureFlags) Enabled(name string) bool {
	// FIX ME: read under RLock
	return false
}

// SetFlags replaces all flags (hot reload).
func (f *FeatureFlags) SetFlags(m map[string]bool) {
	// FIX ME: copy m under Lock
}
