package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	JWTProvider interface {
		GenerateToken(userID int) (string, error)
	}

	jwtProvider struct {
		secretKey  string
		expiration time.Duration
	}
)

func NewJWTProvider(secret string, expirationHours int) JWTProvider {
	return &jwtProvider{
		secretKey:  secret,
		expiration: time.Hour * time.Duration(expirationHours),
	}
}

func (j *jwtProvider) GenerateToken(userID int) (string, error) {

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(j.expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
