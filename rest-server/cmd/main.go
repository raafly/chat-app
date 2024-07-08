package main

import (
	"github.com/labstack/echo/v4"
	"github.com/raafly/realtime-app/db"
	restserver "github.com/raafly/realtime-app/rest-server"
)

func main() {
	db := db.NewDB()
	e := echo.New()
	restserver.NewAuthRoute(db, e)

	e.Logger.Fatal(e.Start(":1323"))
}