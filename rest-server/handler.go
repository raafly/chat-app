package restserver

import (
	"github.com/labstack/echo/v4"
	"github.com/raafly/realtime-app/helper"
)

type AuthHandler interface {
	Register(c echo.Context) error
	Vertify(c echo.Context) error
	NewContact(c echo.Context) error
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
	u := new(UserDTO)
	if err := c.Bind(u); err != nil {
		return c.JSON(500, helper.ErrInternalServerError())
	}

	otp, err := h.serv.Create(u)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(201, helper.NewCreated("CREATED", map[string]string{"otp": otp}))
}

func (h *AuthHandlerImpl) Vertify(c echo.Context) error {
	telp := c.Request().Header.Get("telp")
	if telp == "" {
		return c.JSON(400, helper.ErrBadRequest("user not login", nil))
	}

	otp := new(OTP)
	if err := c.Bind(otp); err != nil {
		return c.JSON(500, helper.ErrInternalServerError())
	}

	err := h.serv.VertifyOTP(telp, otp.OTP)
	if err != nil {
		return c.JSON(400, err)
	}

	return c.JSON(200, helper.NewSucces("SUCCESS LOG IN", nil))
}

func (h *AuthHandlerImpl) GetHistory(c echo.Context) error {
	userID := c.Param("user_id")
	contactID := c.Param("contact_id")

	resp, err := h.serv.GetHistory(userID, contactID)
	if err != nil {
		return c.JSON(404, err)
	}

	return c.JSON(200, helper.NewContent("success", resp))
}

func (h *AuthHandlerImpl) GetContacts(c echo.Context) error {
	contactID := c.Param("contact_id")
	userID := c.Param("user_id")

	resp, err := h.serv.GetContacts(userID, contactID)
	if err != nil {
		return c.JSON(404, err)
	}

	return c.JSON(200, helper.NewSucces("success", resp))
}

func (h *AuthHandlerImpl) NewContact(c echo.Context) error {
	dto := new(ContactDTO)
	if err := c.Bind(dto); err != nil {
		return c.JSON(500, helper.ErrInternalServerError())
	}

	if err := h.serv.NewContact(dto.UserID, dto.ContactID); err != nil {
		return c.JSON(500, helper.ErrInternalServerError())
	}

	return c.JSON(201, helper.NewCreated("success created", nil))
}
