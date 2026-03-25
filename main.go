package main

import (
	"fmt"
	"log"
	"merlindex/example/config"
	"merlindex/example/controllers"
	"merlindex/example/database"
	"merlindex/example/middlewares"
	"os"

	"github.com/labstack/echo/v5"
)

func main() {
	f, err := os.OpenFile("home.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	err = database.Migrate()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middlewares.OptionalJWT)

	e.Static("/", "build")
	e.POST("/user", controllers.NewCreateUser().Execute)
	e.POST("/login", controllers.NewLogin().Execute)
	e.POST("/logout", controllers.NewLogin().Execute)
	e.GET("/list", controllers.NewSearchList().Execute)
	e.POST("/list", controllers.NewCreateList().Execute, middlewares.RequireAuth)

	err = e.Start(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
