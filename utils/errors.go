package utils

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type HttpError struct {
	statusCode int
	message    string
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{
		statusCode: statusCode,
		message:    message,
	}
}

func (e *HttpError) Error() string {
	return e.message
}

type HttpErrorPayload struct {
	Error string `json:"error"`
}

func (e *HttpError) Handle(ctx *echo.Context) error {
	payload := HttpErrorPayload{Error: e.Error()}
	return ctx.JSON(e.statusCode, payload)
}

var (
	ErrorBadRequest        = NewHttpError(http.StatusBadRequest, "payload has invalid format")
	ErrorInternal          = NewHttpError(http.StatusInternalServerError, "internal error")
	UserNotFound           = NewHttpError(http.StatusNotFound, "user not found")
	IncorrectPassword      = NewHttpError(http.StatusUnauthorized, "incorrect password")
	AuthenticationRequired = NewHttpError(http.StatusUnauthorized, "authentication required")
)
