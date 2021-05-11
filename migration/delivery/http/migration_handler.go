package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type MigrationHandler struct {
	MigrationUseCase domain.MigrationUseCase
}

func NewMigrationHandler(e *echo.Echo, m domain.MigrationUseCase) {
	handler := &MigrationHandler{MigrationUseCase: m}
	e.GET("/migrate", handler.Migrate)
}

func (m *MigrationHandler) Migrate(c echo.Context) error {
	ctx := c.Request().Context()
	if err := m.MigrationUseCase.Migrate(ctx); err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				"migration failed due to server error, please try again later",
				nil,
			),
		)
	}
	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"migration is succeed",
			nil,
		),
	)
}
