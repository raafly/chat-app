package restserver

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/raafly/realtime-app/helper"
)

func middlewareJSON(next echo.HandlerFunc) echo.HandlerFunc {
	v := validator.Validate{}
	var messages []string

	return func(c echo.Context) error {
		var data map[string]interface{}
		c.Bind(&data)
		log.Printf("data %s", data)

		if err := v.Struct(data); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				switch e.Tag() {
				case "password":
					messages = append(messages, fmt.Sprintf("%s: minimal 8 caracter", e.StructField()))
				case "required":
					message := fmt.Sprintf("%s: tidak boleh kosong", e.StructField())
					messages = append(messages, message)
				case "email":
					message := fmt.Sprintf("%s: harus format email", e.StructField())
					messages = append(messages, message)
				case "min":
					message := fmt.Sprintf("%s: minimal %s", e.StructField(), e.Param())
					messages = append(messages, message)
				}
			}
		}

		err := next(c)
		return err
	}
}

func NewAuthRoute(db *sql.DB, e *echo.Echo) {
	password := helper.NewPassword()

	repo := NewAuthRepo(db)
	serv := NewAuthService(repo, password)
	handler := NewAuthHandler(serv)

	auth := e.Group("/auth")
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "working...")
	})

	auth.Use(middlewareJSON)

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
	auth.GET("/contacts/:user_id", handler.GetContacts)
	auth.GET("/users/:user_id/contacts/:contact_id/history", handler.GetHistory)
}
