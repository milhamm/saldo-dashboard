package config

import (
	"context"
	"os"
	"saldo-server/domain"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type key string

const (
	USER_ID key = "user_id"
)

func InitJWT() echojwt.Config {
	secret := os.Getenv("JWT_SECRET")
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(domain.JWTClaims)
		},
		SuccessHandler: func(c echo.Context) {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(*domain.JWTClaims)

			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, USER_ID, claims.UserID)
			c.SetRequest(c.Request().WithContext(ctx))
		},
		SigningKey: []byte(secret),
	}
}
