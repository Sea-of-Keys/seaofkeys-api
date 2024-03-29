package security

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

// ##### make a type struck?? ######

type Claims struct {
	ID    uint
	Email string
	jwt.RegisteredClaims
}

// ##### Nedds to return a token (maby a string) ######
func NewPasswordToken(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("PSCRERT")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func NewToken(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("PSCRERT")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func CheckToken(tokenString, secretKey string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return false, err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}
	return false, fmt.Errorf("Invalid Token")
}
func DecodeToken(tokenString, secretKey string, test models.Token) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Invalid Token")
}
func RefreshToken(tokenString, secretKey string) (uint, string, error) {
	var mToken models.Token
	fmt.Print(mToken)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return 0, "", fmt.Errorf("Invalid Token")
	}
	id := claims["ID"].(float64)
	email := claims["Email"].(string)
	ID := uint(id)
	Email := email
	// newToken, err := NewToken(id, email)
	if err != nil {
		return 0, "", errors.New("Failed to make a new token")
	}
	return ID, Email, nil
}

func TokenMiddleware(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		sess, err := store.Get(c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		fmt.Printf("Middleware Session: %v\n", sess)
		fmt.Printf("Middleware Session: %v\n", sess)
		tokenInter := sess.Get("ActiveToken")
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
func WebsiteTokenmiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		cookieHeader3 := c.Request().Header.Peek("Authorization")
		cookies3 := string(cookieHeader3)

		if cookies3 == "" {
			return c.Status(401).JSON(fiber.Map{
				"message": "No Token Provided",
			})
		}
		if len(cookies3) < 8 {
			return c.Status(401).JSON(fiber.Map{
				"message": "No Token Provided",
			})
		}
		tokenString := cookies3[7:]
		if ok, err := CheckToken(tokenString, os.Getenv("PSCRERT")); !ok || err != nil {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}
func LoggingMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

// func RefreshTokenV2(tokenString, secretKey string) (*models.Token, error) {
// 	var mToken models.Token
// 	fmt.Print(mToken)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secretKey), nil
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok && !token.Valid {
// 		return nil, fmt.Errorf("Invalid Token")
// 	}
// 	id := claims["ID"].(uint)
// 	email := claims["Email"].(string)
// 	mToken.ID = id
// 	mToken.Email = email
// 	// newToken, err := NewToken(id, email)
// 	if err != nil {
// 		return nil, errors.New("Failed to make a new token")
// 	}
// 	return &mToken, nil
// }
