package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

func NewEmbeddedToken() (string, error) {
	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	token := base64.URLEncoding.EncodeToString(bytes)

	token = strings.ReplaceAll(token, "+", "")
	token = strings.ReplaceAll(token, "/", "")
	token = strings.ReplaceAll(token, "\\", "")

	fmt.Printf("Token: %v\n", token)
	return token, nil
}
