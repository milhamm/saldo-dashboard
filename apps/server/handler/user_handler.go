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
	user, _ := u.UserService.GetUserByID(ctx, userID)

	return c.JSON(http.StatusOK, user)
}
