package models

import "github.com/golang-jwt/jwt"

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
