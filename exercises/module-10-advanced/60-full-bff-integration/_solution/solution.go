package bffinteg

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

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

func buildDashboard(ctx context.Context, userID string, users UserClient, orders OrderClient) (Dashboard, error) {
	var (
		user      User
		orderList []Order
		userErr   error
		ordersErr error
		wg        sync.WaitGroup
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

func cacheKey(userID string) string {
	return fmt.Sprintf("dash:%s", userID)
}

func SetupApp(d Deps) *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Get("/dashboard/:userID", func(c *fiber.Ctx) error {
		userID := c.Params("userID")
		key := cacheKey(userID)

		if cached, ok := d.Cache.Get(key); ok {
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
			return c.SendString(cached)
		}

		dash, err := buildDashboard(c.Context(), userID, d.Users, d.Orders)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		body, err := json.Marshal(dash)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		s := string(body)
		d.Cache.Set(key, s)

		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.SendString(s)
	})

	return app
}
