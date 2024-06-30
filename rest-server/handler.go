package restserver

import "github.com/labstack/echo/v4"

type AuthRest interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	GetProfile(c echo.Context) error
	GetContacts(c echo.Context) error
	GetHistory(c echo.Context) error
}

type AuthRestImpl struct {
	
}

func NewAuthRest()