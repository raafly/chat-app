package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raafly/realtime-app/db"
	restserver "github.com/raafly/realtime-app/rest-server"
)

func main() {
	db := db.NewDB()
	e := echo.New()
	restserver.NewAuthRoute(db, e)
	
	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.Logger.Fatal(e.Start(":1323"))
}