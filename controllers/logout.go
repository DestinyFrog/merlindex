package controllers

import (
	"merlindex/example/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type Logout struct{}

func NewLogout() *Logout {
	d := Logout{}
	return &d
}

func (c *Logout) Execute(ctx *echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	cookie.HttpOnly = true

	ctx.SetCookie(cookie)
	return utils.SendJsonMessage(ctx, "Successfully Logout")
}
