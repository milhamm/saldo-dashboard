package handler

import (
	"net/http"
	"saldo-server/domain"
	"saldo-server/service"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(e *echo.Group, authService *service.AuthService) {
	handler := &AuthHandler{
		AuthService: authService,
	}
	e.POST("/login", handler.Login)
}

func (a *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()

	u := new(domain.LoginRequest)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, domain.GenericResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
	}

	token, err := a.AuthService.Login(ctx, u)
	if err != nil {
		return handleGenericError(c, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
