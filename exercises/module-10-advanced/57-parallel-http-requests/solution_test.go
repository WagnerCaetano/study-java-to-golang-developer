package parallelhttp

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type staticFetcher map[string]string

func (m staticFetcher) Fetch(_ context.Context, url string) (string, error) {
	if body, ok := m[url]; ok {
		return body, nil
	}
	return "", errors.New("not found: " + url)
}

func TestFetchAllPreservesOrder(t *testing.T) {
	urls := []string{"https://a", "https://b", "https://c"}
	f := staticFetcher{
		"https://a": "A",
		"https://b": "B",
		"https://c": "C",
	}
	got, err := FetchAll(context.Background(), f, urls)
	if err != nil {
		t.Fatalf("FetchAll: %v", err)
	}
	want := []string{"A", "B", "C"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("got[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}

func TestFetchAllFirstError(t *testing.T) {
	f := staticFetcher{"https://ok": "ok"}
	_, err := FetchAll(context.Background(), f, []string{"https://ok", "https://bad"})
	if err == nil {
		t.Fatal("expected error")
	}
}

type peakConcurrencyFetcher struct {
	mu   sync.Mutex
	cur  int
	peak int
}

func (p *peakConcurrencyFetcher) Fetch(ctx context.Context, url string) (string, error) {
	p.mu.Lock()
	p.cur++
	if p.cur > p.peak {
		p.peak = p.cur
	}
	p.mu.Unlock()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(20 * time.Millisecond):
	}

	p.mu.Lock()
	p.cur--
	p.mu.Unlock()
	return url, nil
}

func TestFetchAllConcurrent(t *testing.T) {
	f := &peakConcurrencyFetcher{}
	urls := make([]string, 4)
	for i := range urls {
		urls[i] = "https://x"
	}
	_, err := FetchAll(context.Background(), f, urls)
	if err != nil {
		t.Fatalf("FetchAll: %v", err)
	}
	f.mu.Lock()
	peak := f.peak
	f.mu.Unlock()
	if peak < 2 {
		t.Errorf("peak concurrency = %d, want at least 2", peak)
	}
}

func TestFetchAllCancelsOnError(t *testing.T) {
	var started atomic.Int32
	block := make(chan struct{})
	f := fetcherFunc(func(ctx context.Context, url string) (string, error) {
		if url == "fail" {
			return "", errors.New("boom")
		}
		started.Add(1)
		select {
		case <-block:
			return url, nil
		case <-ctx.Done():
			return "", ctx.Err()
		}
	})

	urls := []string{"fail", "slow1", "slow2"}
	_, err := FetchAll(context.Background(), f, urls)
	if err == nil {
		close(block)
		t.Fatal("expected error")
	}
	close(block)
	if started.Load() > 0 {
		t.Log("slow fetches may have started before cancel; ok if error returned")
	}
}

type fetcherFunc func(context.Context, string) (string, error)

func (fn fetcherFunc) Fetch(ctx context.Context, url string) (string, error) {
	return fn(ctx, url)
}
