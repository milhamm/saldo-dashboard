package handler

import (
	"net/http"
	"saldo-server/config"
	"saldo-server/domain"
	"saldo-server/service"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	AccountService *service.AccountService
}

func NewAccountHandler(e *echo.Group, accountService *service.AccountService) {
	handler := &AccountHandler{
		AccountService: accountService,
	}
	e.GET("", handler.ListAccounts)
	e.POST("", handler.CreateAccount)
}

func (a *AccountHandler) ListAccounts(c echo.Context) error {
	ctx := c.Request().Context()
	userID := ctx.Value(config.USER_ID).(string)

	accounts, err := a.AccountService.ListAccounts(ctx, userID)
	if err != nil {
		return handleGenericError(c, err)
	}

	response := domain.SuccessResponse{
		Data: accounts,
	}

	return c.JSON(http.StatusOK, response.ToGenericResponse())

}

func (a *AccountHandler) CreateAccount(c echo.Context) error {
	ctx := c.Request().Context()
	userID := ctx.Value(config.USER_ID).(string)

	var ca domain.CreateAccount
	if err := c.Bind(&ca); err != nil {
		return c.JSON(http.StatusBadRequest, domain.GenericResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    err.Error(),
		})
	}

	account, err := a.AccountService.CreateAccount(ctx, userID, &ca)
	if err != nil {
		return handleGenericError(c, err)
	}

	return c.JSON(http.StatusCreated, domain.GenericResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    account,
	})
}
