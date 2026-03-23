package middlewares

import (
	"fmt"
	"merlindex/example/config"
	"merlindex/example/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func OptionalJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx *echo.Context) error {
		cookie, err := ctx.Cookie("token")
		if err != nil {
			return next(ctx)
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &utils.JwtClaims{}, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método inválido: %v", t.Header["alg"])
			}
			return config.JwtSecret, nil
		})

		if err == nil && token.Valid {
			ctx.Set("token", token.Claims.(*utils.JwtClaims))
		}

		return next(ctx)
	}
}
