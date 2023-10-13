package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func SessionMiddleware() func(*fiber.Ctx) error {
	// store := session.New()
	store := session.New(session.Config{
		KeyLookup:  "cookie:sessionid",
		Expiration: time.Hour * 24, // Session expiration time
	})

	return func(c *fiber.Ctx) error {
		// Attach the session to the context
		c.Locals("session", store)
		return c.Next()
	}
}
