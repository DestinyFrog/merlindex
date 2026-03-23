package controllers

import (
	"log"
	"merlindex/example/database"
	"merlindex/example/models"
	"merlindex/example/utils"

	q "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/labstack/echo/v5"
)

type SearchList struct{}

func NewSearchList() *SearchList {
	d := SearchList{}
	return &d
}

func (c *SearchList) Execute(ctx *echo.Context) error {
	var payload struct {
		Search string `query:"search"`
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

	sql, _, err := q.From("list").ToSQL()
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	var lists []*models.List
	err = sqlscan.Select(ctx.Request().Context(), db, &lists, sql)
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	return ctx.JSON(200, &lists)
}
