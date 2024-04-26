package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("vk_love")

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

type AccessToken struct {
	Token   string `json:"access_token"`
	Expires int64  `json:"expires_in"`
}
