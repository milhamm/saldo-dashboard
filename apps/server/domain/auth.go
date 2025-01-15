package domain

import "github.com/golang-jwt/jwt/v5"

type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
