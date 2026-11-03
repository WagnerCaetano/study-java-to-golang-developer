package fiberjson

import "github.com/gofiber/fiber/v2"

// CreateUserRequest is the JSON body for POST /users.
type CreateUserRequest struct {
	Name  string `json:"name"`  // FIX ME: wire up POST /users
	Email string `json:"email"` // FIX ME: wire up POST /users
}

// SetupApp creates a Fiber app with POST /users.
// FIX ME: BodyParser → 201 JSON echo; invalid body → 400 {"error":"invalid body"}.
func SetupApp() *fiber.App {
	app := fiber.New()
	return app
}
