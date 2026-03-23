package controllers

import "github.com/labstack/echo/v5"

type Controller interface {
	Execute(ctx *echo.Context) error
}
