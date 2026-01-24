package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword transforma "123456" em algo como "$2a$10$X5..."
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

// CheckPassword verifica se a senha bate com o hash salvo no banco
func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
