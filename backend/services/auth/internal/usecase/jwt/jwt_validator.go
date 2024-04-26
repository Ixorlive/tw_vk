package jwt

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type JWTValidator struct{}

func (v *JWTValidator) ValidateToken(tokenString string) (string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.Login, nil
}
