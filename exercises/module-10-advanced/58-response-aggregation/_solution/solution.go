package aggregate

import (
	"context"
	"sync"
)

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
	var (
		user       User
		orderList  []Order
		userErr    error
		ordersErr  error
		wg         sync.WaitGroup
	)

	wg.Add(2)
	go func() {
		defer wg.Done()
		user, userErr = users.GetUser(ctx, userID)
	}()
	go func() {
		defer wg.Done()
		orderList, ordersErr = orders.ListOrders(ctx, userID)
	}()
	wg.Wait()

	if userErr != nil {
		return Dashboard{}, userErr
	}
	if ordersErr != nil {
		return Dashboard{}, ordersErr
	}
	return Dashboard{User: user, Orders: orderList}, nil
}
