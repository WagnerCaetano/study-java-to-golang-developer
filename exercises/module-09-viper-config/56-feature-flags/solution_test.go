package featureflags

import "testing"

func TestFeatureFlagsEnabled(t *testing.T) {
	ff := NewFeatureFlagsFromMap(map[string]bool{
		"new-checkout": true,
		"beta-banner":  false,
	})
	if !ff.Enabled("new-checkout") {
		t.Error("expected new-checkout enabled")
	}
	if ff.Enabled("beta-banner") {
		t.Error("expected beta-banner disabled")
	}
	if ff.Enabled("unknown") {
		t.Error("unknown flag should be false")
	}
}

func TestFeatureFlagsHotUpdate(t *testing.T) {
	ff := NewFeatureFlagsFromMap(map[string]bool{"dark-mode": false})
	ff.SetFlags(map[string]bool{"dark-mode": true})
	if !ff.Enabled("dark-mode") {
		t.Error("expected dark-mode after SetFlags")
	}
}
