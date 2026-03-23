package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

type JwtClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func CurrentUser(ctx *echo.Context) *JwtClaims {
	user, _ := ctx.Get("token").(*JwtClaims)
	return user
}
