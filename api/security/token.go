package security

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

// ##### make a type struck?? ######

// type Claims struct {
// 	ID    uint
// 	Email string
// 	jwt.RegisteredClaims
// }

// ##### Nedds to return a token (maby a string) ######
func NewPasswordToken(id uint, email string) (string, error) {
	expirationTime := time.Now().Add(32 * time.Hour)
	claims := &models.Claims{
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
	expirationTime := time.Now().Add(2 * time.Hour)
	claims := &models.Claims{
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
	return false, fmt.Errorf("invalid token")
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
	return nil, fmt.Errorf("invalid token")
}
func GetTokenData(tokenString, secretKey string) (uint, string, error) {
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
		return 0, "", fmt.Errorf("invalid token")
	}
	id := claims["ID"].(float64)
	email := claims["Email"].(string)
	ID := uint(id)
	Email := email
	if err != nil {
		return 0, "", errors.New("failed to make a new token")
	}
	return ID, Email, nil
}

func TokenMiddleware(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			log.Println(err)
			return fiber.NewError(fiber.StatusInternalServerError, "failed to find session")
		}
		tokenInter := sess.Get("ActiveToken")
		tokenString, ok := tokenInter.(string)
		if !ok || tokenString == "" {
			log.Println(err)
			return fiber.NewError(fiber.StatusNonAuthoritativeInformation, "M101 No token providet")
		}
		if ok, err := CheckToken(tokenString, os.Getenv("PSCRERT")); !ok || err != nil {
			log.Println(err)
			log.Println(ok)
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		return c.Next()
	}
}
func WebsiteTokenMiddleware(store *session.Store) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		sess, err := store.Get(c)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		tokenInter := sess.Get("WebToken")
		tokenString, ok := tokenInter.(string)
		// fmt.Printf("tokenString: %v\n", tokenString)
		// fmt.Printf("tokenString: %v\n", tokenString)
		fmt.Printf("tokenString: %v\n", tokenString)
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
func LoggingMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}
