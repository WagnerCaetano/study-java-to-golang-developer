package hotreload

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestConfigWatcherHotReloadYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(path, []byte("message: hello-v1\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	updates := make(chan string, 2)
	w, err := NewConfigWatcher(path, func(cfg AppConfig) {
		select {
		case updates <- cfg.Message:
		default:
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = w.Close() })

	waitMessage(t, updates, "hello-v1", 2*time.Second)
	if got := w.Current().Message; got != "hello-v1" {
		t.Fatalf("current = %q", got)
	}

	if err := os.WriteFile(path, []byte("message: hello-v2\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	waitMessage(t, updates, "hello-v2", 3*time.Second)
	if got := w.Current().Message; got != "hello-v2" {
		t.Fatalf("current after reload = %q", got)
	}
}

func waitMessage(t *testing.T, ch <-chan string, want string, timeout time.Duration) {
	t.Helper()
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		select {
		case got := <-ch:
			if got == want {
				return
			}
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
	t.Fatalf("timed out waiting for message %q", want)
}
