package main

import (
	"fmt"
	"merlindex/example/config"

	"github.com/labstack/echo/v5"
)

func main() {
	e := echo.New()

	e.Static("/", "public")

	err := e.Start(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
