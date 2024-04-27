package jwt

import (
	"errors"

	"github.com/Ixorlive/tw_vk/backend/services/auth/internal/entity"
	"github.com/golang-jwt/jwt/v4"
)

type JWTValidator struct{}

func NewJWTValidator() JWTValidator {
	return JWTValidator{}
}

func (v *JWTValidator) ValidateToken(tokenString string) (entity.User, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	var user entity.User

	if err != nil {
		return user, err
	}

	if !token.Valid {
		return user, errors.New("invalid token")
	}
	user.Id = claims.Id
	user.Login = claims.Login
	return user, nil
}
