package envconfig

import (
	"os"
	"testing"
)

func TestLoadFromEnvSuccess(t *testing.T) {
	t.Setenv("PORT", "8080")
	t.Setenv("APP_NAME", "bff")

	cfg, err := LoadFromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != 8080 || cfg.AppName != "bff" {
		t.Fatalf("cfg = %+v", cfg)
	}
}

func TestLoadFromEnvMissingAppName(t *testing.T) {
	t.Setenv("PORT", "8080")
	_ = os.Unsetenv("APP_NAME")

	if _, err := LoadFromEnv(); err == nil {
		t.Fatal("expected error for missing APP_NAME")
	}
}

func TestLoadFromEnvBadPort(t *testing.T) {
	t.Setenv("PORT", "not-a-number")
	t.Setenv("APP_NAME", "bff")

	if _, err := LoadFromEnv(); err == nil {
		t.Fatal("expected error for invalid PORT")
	}
}
