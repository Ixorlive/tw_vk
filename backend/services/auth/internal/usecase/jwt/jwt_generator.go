package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTGenerator struct {
	signingMethod    jwt.SigningMethod
	tokenExpireAfter time.Duration
}

func NewJWTGenerator(signingMethod jwt.SigningMethod) *JWTGenerator {
	return &JWTGenerator{
		signingMethod:    signingMethod,
		tokenExpireAfter: time.Hour * 24,
	}
}

// Generate generates an AccessToken using the username and role claims.
func (gen *JWTGenerator) Generate(login string) (*AccessToken, error) {
	token := jwt.New(gen.signingMethod)
	claims := Claims{}

	// set custom claims
	claims.Login = login

	// set standard claims
	now := time.Now()
	claims.IssuedAt = jwt.NewNumericDate(now)
	if gen.tokenExpireAfter > 0 {
		claims.ExpiresAt = jwt.NewNumericDate(now.Add(gen.tokenExpireAfter))
	}

	token.Claims = &claims
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	// create an access token
	accessToken := &AccessToken{
		Token:   tokenString,
		Expires: gen.tokenExpireAfter.Milliseconds(),
	}

	return accessToken, nil
}