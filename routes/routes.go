package routes

import (
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewRoutes(e *echo.Echo) {
	e.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, custom_response.NewCustomResponse(
			true,
			"welcome to API, please contact the maintainer",
			nil),
		)
	})
}
