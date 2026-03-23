package controllers

import (
	"log"
	"merlindex/example/database"
	"merlindex/example/utils"

	q "github.com/doug-martin/goqu/v9"
	"github.com/labstack/echo/v5"
)

type CreateList struct{}

func NewCreateList() *CreateList {
	d := CreateList{}
	return &d
}

func (c *CreateList) Execute(ctx *echo.Context) error {
	var payload struct {
		Title string `json:"title"`
	}

	if err := ctx.Bind(&payload); err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorBadRequest.Handle(ctx)
	}

	jwt := utils.CurrentUser(ctx)

	db, err := database.Db()
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}
	defer db.Close()

	sql, _, err := q.Insert("list").
		Cols("title", "user_id").
		Vals(q.Vals{payload.Title, jwt.ID}).
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
