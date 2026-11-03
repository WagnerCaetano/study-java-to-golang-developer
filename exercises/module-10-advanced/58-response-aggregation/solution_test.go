package aggregate

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakeUserClient struct {
	user User
	err  error
	delay time.Duration
}

func (f *fakeUserClient) GetUser(ctx context.Context, id string) (User, error) {
	if f.delay > 0 {
		select {
		case <-time.After(f.delay):
		case <-ctx.Done():
			return User{}, ctx.Err()
		}
	}
	if f.err != nil {
		return User{}, f.err
	}
	u := f.user
	if u.ID == "" {
		u.ID = id
	}
	return u, nil
}

type fakeOrderClient struct {
	orders []Order
	err    error
	delay  time.Duration
}

func (f *fakeOrderClient) ListOrders(ctx context.Context, userID string) ([]Order, error) {
	if f.delay > 0 {
		select {
		case <-time.After(f.delay):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
	if f.err != nil {
		return nil, f.err
	}
	return f.orders, nil
}

func TestBuildDashboardAggregates(t *testing.T) {
	users := &fakeUserClient{user: User{ID: "u1", Name: "Ada"}}
	orders := &fakeOrderClient{orders: []Order{{ID: "o1", UserID: "u1"}}}

	dash, err := BuildDashboard(context.Background(), "u1", users, orders)
	if err != nil {
		t.Fatalf("BuildDashboard: %v", err)
	}
	if dash.User.Name != "Ada" {
		t.Errorf("user name = %q, want Ada", dash.User.Name)
	}
	if len(dash.Orders) != 1 || dash.Orders[0].ID != "o1" {
		t.Errorf("orders = %+v, want one order o1", dash.Orders)
	}
}

func TestBuildDashboardUserError(t *testing.T) {
	users := &fakeUserClient{err: errors.New("user down")}
	orders := &fakeOrderClient{orders: []Order{{ID: "o1", UserID: "u1"}}}

	_, err := BuildDashboard(context.Background(), "u1", users, orders)
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestBuildDashboardParallel(t *testing.T) {
	const delay = 80 * time.Millisecond

	users := &fakeUserClient{user: User{Name: "Bob"}, delay: delay}
	orders := &fakeOrderClient{orders: []Order{{ID: "o1"}}, delay: delay}

	start := time.Now()
	_, err := BuildDashboard(context.Background(), "u1", users, orders)
	elapsed := time.Since(start)
	if err != nil {
		t.Fatalf("BuildDashboard: %v", err)
	}
	if elapsed >= delay*2 {
		t.Errorf("took %v; expected parallel (~%v), sequential would be ~%v", elapsed, delay, delay*2)
	}
}
