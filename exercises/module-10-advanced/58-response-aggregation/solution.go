package aggregate

import "context"

type User struct {
	ID   string
	Name string
}

type Order struct {
	ID     string
	UserID string
}

type Dashboard struct {
	User   User
	Orders []Order
}

type UserClient interface {
	GetUser(ctx context.Context, id string) (User, error)
}

type OrderClient interface {
	ListOrders(ctx context.Context, userID string) ([]Order, error)
}

func BuildDashboard(ctx context.Context, userID string, users UserClient, orders OrderClient) (Dashboard, error) {
	// FIX ME: call UserClient and OrderClient in parallel, then aggregate
	var dash Dashboard
	return dash, nil
}
