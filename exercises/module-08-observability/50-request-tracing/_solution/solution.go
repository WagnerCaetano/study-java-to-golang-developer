package reqtrace

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gofiber/fiber/v2"
)

const HeaderRequestID = "X-Request-ID"

const localRequestID = "request_id"

func GenerateRequestID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Get(HeaderRequestID)
		if id == "" {
			id = GenerateRequestID()
		}
		c.Locals(localRequestID, id)
		c.Set(HeaderRequestID, id)
		return c.Next()
	}
}

func GetRequestID(c *fiber.Ctx) string {
	v := c.Locals(localRequestID)
	if v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}
