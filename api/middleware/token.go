package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
