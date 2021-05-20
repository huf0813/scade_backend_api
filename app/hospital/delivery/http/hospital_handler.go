package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type HospitalHandler struct {
	HospitalUseCase domain.HospitalUseCase
}

func NewHospitalHandler(e *echo.Echo, h domain.HospitalUseCase) {
	handler := HospitalHandler{HospitalUseCase: h}
	e.GET("/hospitals", handler.GetHospitals)
	e.GET("/hospitals/:city", handler.GetHospitalsByCity)
	e.GET("/hospitals/:id", handler.GetHospitalByID)
}

func (h *HospitalHandler) GetHospitals(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.HospitalUseCase.GetHospitals(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get hospitals successfully",
		result),
	)
}

func (h *HospitalHandler) GetHospitalsByCity(c echo.Context) error {
	city := c.Param("city")
	ctx := c.Request().Context()
	result, err := h.HospitalUseCase.GetHospitalsByCity(ctx, city)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get hospitals by city successfully",
		result),
	)
}

func (h *HospitalHandler) GetHospitalByID(c echo.Context) error {
	id := c.Param("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	result, err := h.HospitalUseCase.GetHospitalByID(ctx, idInteger)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"get hospital by id successfully",
		result),
	)
}
