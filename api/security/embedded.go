package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"
)

func NewBase64Token() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(bytes)

	fmt.Printf("Token: %v\n", token)
	return token, nil
}

func CheckEmbeddedToken(tokenString, secretKey string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}
	return false, fmt.Errorf("invalid token")
}
func NewEmbeddedToken(token string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": token,
		"exp":   time.Now().Add(time.Hour * 28).Unix(),
	})
	tokenString, err := claims.SignedString([]byte(os.Getenv("PSCRERT")))
	fmt.Printf("TokenString: %v\n", tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func TokenEmbeddedMiddleware(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		tokenInter := sess.Get("EmbeddedSession")
		tokenString, ok := tokenInter.(string)
		if !ok || tokenString == "" {
			return fiber.NewError(fiber.StatusNonAuthoritativeInformation, "M101 No token providet")
		}
		if ok, err := CheckToken(tokenString, os.Getenv("PSCRERT")); !ok || err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}
