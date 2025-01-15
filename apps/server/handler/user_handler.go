package handler

import (
	"net/http"
	"saldo-server/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(routeGroup *echo.Group, userService *service.UserService) {
	handler := &UserHandler{
		UserService: userService,
	}
	routeGroup.GET("/:id", handler.ReadUserByID)
}

func (u *UserHandler) ReadUserByID(c echo.Context) error {
	ctx := c.Request().Context()

	userID := c.Param("id")
	user, err := u.UserService.GetUserByID(ctx, userID)
	if err != nil {
		return handleGenericError(c, err)
	}

	return c.JSON(http.StatusOK, user)
}
