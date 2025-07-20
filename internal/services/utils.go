package services

import (
	"errors"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func validateCredentials(login, password string) error {
	if len(login) < 4 || len(login) > 32 {
		return errors.New("login must be 4-32 characters")
	}
	if len(password) < 6 || len(password) > 64 {
		return errors.New("password must be 6-64 characters")
	}
	for _, r := range login {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
			return errors.New("login contains invalid characters")
		}
	}
	for _, r := range password {
		if unicode.IsSpace(r) {
			return errors.New("password must not contain spaces")
		}
	}
	return nil
}
