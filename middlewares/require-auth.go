package middlewares

import (
	"log"
	"merlindex/example/utils"

	"github.com/labstack/echo/v5"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx *echo.Context) error {
		if utils.CurrentUser(ctx) == nil {
			log.Printf("error: %v", "none user logged")
			return utils.AuthenticationRequired.Handle(ctx)
		}
		return next(ctx)
	}
}
