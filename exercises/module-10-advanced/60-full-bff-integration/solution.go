package bffinteg

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
}

type Dashboard struct {
	User   User    `json:"user"`
	Orders []Order `json:"orders"`
}

type UserClient interface {
	GetUser(ctx context.Context, id string) (User, error)
}

type OrderClient interface {
	ListOrders(ctx context.Context, userID string) ([]Order, error)
}

type Deps struct {
	Users  UserClient
	Orders OrderClient
	Cache  interface {
		Get(k string) (string, bool)
		Set(k, v string)
	}
}

func SetupApp(d Deps) *fiber.App {
	// FIX ME: /health and /dashboard/:userID with cache-aside
	app := fiber.New()
	return app
}

func buildDashboard(ctx context.Context, userID string, users UserClient, orders OrderClient) (Dashboard, error) {
	// FIX ME: parallel aggregation (used by dashboard handler)
	return Dashboard{}, nil
}
