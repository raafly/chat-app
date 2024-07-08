package restserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raafly/realtime-app/helper"
)

type AuthHandler interface {
	Register(c echo.Context) error
	Login(c echo.Context) error
	// GetProfile(c echo.Context) error
	GetContacts(c echo.Context) error
	GetHistory(c echo.Context) error
}

type AuthHandlerImpl struct {
	serv AuthService
}

func NewAuthHandler(serv AuthService) AuthHandler {
	return &AuthHandlerImpl{serv: serv}
}

func (h *AuthHandlerImpl) Register(c echo.Context) error {
	u := new(UserReq)
	if err := c.Bind(u); err != nil {
		return c.JSON(500, helper.ErrInternalServerError())
	}

	if err := h.serv.Create(u); err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(201, helper.NewCreated("CREATED"))
}

func (h *AuthHandlerImpl) Login(c echo.Context) error {
	u := new(UserReq)
	if err := c.Bind(u); err != nil {
		return c.JSON(500, helper.ErrInternalServerError())
	}

	_, err := h.serv.Login(u)
	if err != nil {
		return c.JSON(404, err)
	}

	return c.JSON(200, helper.NewSucces("SUCCESS LOG IN"))
}

func (h *AuthHandlerImpl) GetHistory(c echo.Context) error {
	userID := c.QueryParam("user_id")
	contactID := c.QueryParam("contact_id")

	resp, err := h.serv.GetHistory(userID, contactID)
	if err != nil {
		return err
	}

	return helper.NewContent(resp)
}

func (h *AuthHandlerImpl) GetContacts(c echo.Context) error {
	userID := c.QueryParam("user_id")
	if userID == "" {
		return c.JSON(400, http.StatusBadRequest)
	}

	resp, err := h.serv.GetContacts(userID)
	if err != nil {
		return err
	}

	return helper.NewContent(resp)
}
