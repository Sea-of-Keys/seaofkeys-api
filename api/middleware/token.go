package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

// ##### make a type struck?? ######

// ##### Nedds to return a token (maby a string) ######
func TokenMiddleware(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			// return fiber.newError(fiber.StatusInternalServerError, err.Error())
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		tokenInter := sess.Get("ActiveToken")
		fmt.Printf("Token1: %v\n", tokenInter)
		tokenString, ok := tokenInter.(string)
		fmt.Printf("Token2: %v\n", tokenString)
		if !ok || tokenString == "" {
			return fiber.NewError(fiber.StatusNonAuthoritativeInformation, "M101 No token providet")
		}
		if ok, err := security.CheckToken(tokenString, os.Getenv("PSCRERT")); !ok || err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
		// fmt.Printf("Token: %v\n", name)
	}
}
