package controllers

import (
	"database/sql"
	"log"
	"merlindex/example/config"
	"merlindex/example/database"
	"merlindex/example/utils"
	"net/http"
	"time"

	q "github.com/doug-martin/goqu/v9"
	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
	"golang.org/x/crypto/bcrypt"
)

type Login struct{}

func NewLogin() *Login {
	d := Login{}
	return &d
}

func (c *Login) Execute(ctx *echo.Context) error {
	var payload struct {
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

	query, _, err := q.From("user").Where(q.Ex{"email": payload.Email}).ToSQL()
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	var user database.User
	rows, err := db.QueryContext(ctx.Request().Context(), query)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("error: %v", err)
			return utils.UserNotFound.Handle(ctx)
		}
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	err = sqlscan.ScanOne(&user, rows)
	if err != nil {
		log.Printf("error: %v", err)
		return utils.ErrorInternal.Handle(ctx)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		log.Printf("error: %v", err)
		return utils.IncorrectPassword.Handle(ctx)
	}

	claims := &utils.JwtClaims{
		UserID: uint(user.Id),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JwtSecret))
	if err != nil {
		log.Printf("error: %v", err)
		return utils.IncorrectPassword.Handle(ctx)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = true
	cookie.SameSite = http.SameSiteStrictMode
	ctx.SetCookie(cookie)

	return utils.SendJsonMessage(ctx, "Successfully Logged In")
}
