package utils

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func SendJsonMessage(ctx *echo.Context, message string) error {
	return ctx.JSON(http.StatusOK, struct {
		Message string
	}{
		Message: message,
	})
}
