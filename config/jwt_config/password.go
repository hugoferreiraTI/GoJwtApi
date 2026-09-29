package jwtconfig

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//encrypted password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.New("Falhou na hash da senha")
	}
	return string(bytes), nil

}

// compare the password with password encrypt in database
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password precisa de no mínimo 8 caracteres")
	}
	return nil
}
