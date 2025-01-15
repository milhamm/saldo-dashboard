package handler

import (
	"net/http"
	"saldo-server/domain"

	"github.com/labstack/echo/v4"
)

func handleGenericError(c echo.Context, err error) error {
	if ge, ok := err.(*domain.GenericError); ok {
		return c.JSON(ge.Code, domain.GenericResponse{
			Code:    ge.Code,
			Message: ge.Message,
			Data:    ge.Err.Error(),
		})
	}

	return c.JSON(http.StatusInternalServerError, domain.GenericResponse{
		Code:    http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Data:    err.Error(),
	})
}
