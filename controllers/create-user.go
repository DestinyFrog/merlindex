package controllers

import (
	"log"
	"merlindex/example/database"
	"merlindex/example/utils"

	q "github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct{}

func NewCreateUser() *CreateUser {
	d := CreateUser{}
	return &d
}

func (c *CreateUser) Execute(ctx *echo.Context) error {
	var payload struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.Bind(&payload); err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorBadRequest.Handle(ctx)
	}

	db, err := database.Db()
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}
	defer db.Close()

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	sql, _, err := q.Insert("user").
		Cols("name", "email", "password").
		Vals(q.Vals{payload.Name, payload.Email, string(passwordBytes)}).
		Returning("id").
		ToSQL()
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	var id int
	err = db.QueryRow(sql).Scan(&id)
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	return ctx.JSON(200, struct {
		Id int `json:"id"`
	}{
		Id: id,
	})
}
