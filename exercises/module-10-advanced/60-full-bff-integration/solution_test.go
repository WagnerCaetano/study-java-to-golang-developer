package bffinteg

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type memCache struct {
	mu   sync.Mutex
	data map[string]string
}

func newMemCache() *memCache {
	return &memCache{data: make(map[string]string)}
}

func (m *memCache) Get(k string) (string, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[k]
	return v, ok
}

func (m *memCache) Set(k, v string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[k] = v
}

type stubUsers struct {
	user User
	err  error
	calls int
}

func (s *stubUsers) GetUser(_ context.Context, id string) (User, error) {
	s.calls++
	if s.err != nil {
		return User{}, s.err
	}
	u := s.user
	if u.ID == "" {
		u.ID = id
	}
	return u, nil
}

type stubOrders struct {
	orders []Order
	err    error
	calls  int
}

func (s *stubOrders) ListOrders(_ context.Context, _ string) ([]Order, error) {
	s.calls++
	if s.err != nil {
		return nil, s.err
	}
	return s.orders, nil
}

func TestHealth(t *testing.T) {
	app := SetupApp(Deps{Cache: newMemCache()})
	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Test: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "ok" {
		t.Errorf("body = %q, want ok", body)
	}
}

func TestDashboardCacheMissThenHit(t *testing.T) {
	users := &stubUsers{user: User{ID: "u1", Name: "Ada"}}
	orders := &stubOrders{orders: []Order{{ID: "o1", UserID: "u1"}}}
	cache := newMemCache()
	app := SetupApp(Deps{Users: users, Orders: orders, Cache: cache})

	req := httptest.NewRequest("GET", "/dashboard/u1", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Test: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want 200", resp.StatusCode)
	}
	if users.calls != 1 || orders.calls != 1 {
		t.Errorf("upstream calls users=%d orders=%d, want 1 each on miss", users.calls, orders.calls)
	}

	req2 := httptest.NewRequest("GET", "/dashboard/u1", nil)
	resp2, err := app.Test(req2, -1)
	if err != nil {
		t.Fatalf("Test: %v", err)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != fiber.StatusOK {
		t.Fatalf("cached status = %d, want 200", resp2.StatusCode)
	}
	if users.calls != 1 || orders.calls != 1 {
		t.Errorf("after hit users=%d orders=%d, want still 1 each", users.calls, orders.calls)
	}
}

func TestDashboardUserError502(t *testing.T) {
	users := &stubUsers{err: errors.New("user service unavailable")}
	orders := &stubOrders{orders: []Order{{ID: "o1"}}}
	app := SetupApp(Deps{Users: users, Orders: orders, Cache: newMemCache()})

	req := httptest.NewRequest("GET", "/dashboard/u1", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Test: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != fiber.StatusBadGateway {
		t.Errorf("status = %d, want 502", resp.StatusCode)
	}
	var payload map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if payload["error"] == "" {
		t.Error("expected error field in JSON body")
	}
}
