package http

import (
	"fmt"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/auth"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strconv"
)

type DiagnoseHandler struct {
	DiagnoseUseCase domain.DiagnoseUseCase
}

func NewDiagnoseHandler(e *echo.Echo, d domain.DiagnoseUseCase, authMiddleware middleware.JWTConfig) {
	handler := &DiagnoseHandler{DiagnoseUseCase: d}
	e.GET("/diagnoses",
		handler.GetDiagnoses,
		middleware.JWTWithConfig(authMiddleware))
	e.GET("/diagnoses/:id",
		handler.GetDiagnoseByID,
		middleware.JWTWithConfig(authMiddleware))
	e.GET("/diagnoses/image/:file",
		handler.GetDiagnoseImage)
	e.POST("/diagnoses/create",
		handler.CreateDiagnose,
		middleware.JWTWithConfig(authMiddleware))
}

func (d *DiagnoseHandler) GetDiagnoses(c echo.Context) error {
	authorization := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	result, err := d.DiagnoseUseCase.GetDiagnoses(ctx, token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch history of diagnoses successfully",
		result),
	)
}

func (d *DiagnoseHandler) GetDiagnoseImage(c echo.Context) error {
	filename := c.Param("file")
	path := fmt.Sprintf("assets/skin_image/%s", filename)
	f, err := os.Open(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if os.IsNotExist(err) {
		file, _ := os.Open("assets/default/default.jpg")
		return c.Stream(http.StatusInternalServerError, "image/jpg", file)
	}
	return c.Stream(http.StatusOK, "image/jpg", f)
}

func (d *DiagnoseHandler) GetDiagnoseByID(c echo.Context) error {
	id := c.Param("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	authorization := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	result, err := d.DiagnoseUseCase.GetDiagnoseByID(ctx, token.Email, uint(idInteger))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch history of diagnoses successfully",
		result),
	)
}

func (d *DiagnoseHandler) CreateDiagnose(c echo.Context) error {
	cancerName := c.FormValue("cancer_name")
	cancerPosition := c.FormValue("position")
	cancerImage, err := c.FormFile("cancer_image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	authorization := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authorization)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()

	create := domain.DiagnoseRequest{
		CancerName: cancerName,
		Position:   cancerPosition,
		Price:      10,
		UserEmail:  token.Email,
	}

	lastID, err := d.DiagnoseUseCase.CreateDiagnose(ctx, &create, cancerImage)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"data created successfully",
		lastID),
	)
}
