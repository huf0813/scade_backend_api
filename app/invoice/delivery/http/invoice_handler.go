package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/auth"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type InvoiceHandler struct {
	InvoiceUseCase domain.InvoiceUseCase
}

func NewInvoiceHandler(e *echo.Echo, i domain.InvoiceUseCase, authMiddleware middleware.JWTConfig) {
	handler := InvoiceHandler{InvoiceUseCase: i}
	e.GET("/invoices",
		handler.GetInvoices,
		middleware.JWTWithConfig(authMiddleware))
	e.GET("/invoices/:id",
		handler.GetInvoices,
		middleware.JWTWithConfig(authMiddleware))
	e.POST("/invoices/create",
		handler.CreateInvoice,
		middleware.JWTWithConfig(authMiddleware))
}

func (i *InvoiceHandler) GetInvoices(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	result, err := i.InvoiceUseCase.GetInvoices(ctx, token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch invoices successfully",
		result))
}

func (i *InvoiceHandler) GetInvoiceByID(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	id := c.Param("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	res, err := i.InvoiceUseCase.GetInvoiceByID(ctx, idInteger, token.Email)
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch invoice by id successfully",
		res))
}

func (i *InvoiceHandler) CreateInvoice(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(bearer)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	create := new(domain.InvoiceRequest)
	if err := c.Bind(create); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(create); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if err := i.InvoiceUseCase.CreateInvoice(ctx,
		create,
		token.Email); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"invoice created successfully",
		nil))
}
