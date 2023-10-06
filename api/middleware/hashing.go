package middleware

import "golang.org/x/crypto/bcrypt"

// ######TODO######
// Change password hasing salt from 5 to ??

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func CheckPasswordHash2(password, hash *string) bool {
	passwordBytes := []byte(*password)
	hashBytes := []byte(*hash)
	err := bcrypt.CompareHashAndPassword([]byte(hashBytes), []byte(passwordBytes))
	return err == nil
}
