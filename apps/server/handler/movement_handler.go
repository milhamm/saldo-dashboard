package handler

import (
	"net/http"
	"saldo-server/config"
	"saldo-server/domain"
	"saldo-server/service"

	"github.com/labstack/echo/v4"
)

type MovementHandler struct {
	MovementService *service.MovementService
}

func NewMovementHandler(e *echo.Group, movementService *service.MovementService) {
	handler := &MovementHandler{
		MovementService: movementService,
	}

	e.GET("", handler.ReadManyMovements)
	e.POST("/:id", handler.CreateMovement)
}

func (m *MovementHandler) ReadManyMovements(c echo.Context) error {
	ctx := c.Request().Context()

	userID := ctx.Value(config.USER_ID).(string)
	movements, err := m.MovementService.ListMovements(ctx, userID)

	if err != nil {
		return err
	}

	response := domain.SuccessResponse{
		Data: movements,
	}

	return c.JSON(http.StatusOK, response.ToGenericResponse())
}

func (m *MovementHandler) CreateMovement(c echo.Context) error {
	ctx := c.Request().Context()
	accountID := c.Param("id")
	userID := ctx.Value(config.USER_ID).(string)

	payload := new(domain.CreateMovement)

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, domain.GenericResponse{
			Code:    http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
			Data:    nil,
		})
	}

	movement, err := m.MovementService.CreateMovement(ctx, userID, accountID, payload)

	if err != nil {
		return handleGenericError(c, err)
	}

	return c.JSON(http.StatusCreated, domain.GenericResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    movement,
	})
}
