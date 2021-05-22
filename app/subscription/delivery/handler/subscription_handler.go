package handler

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/auth"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

type SubscriptionHandler struct {
	SubscriptionUseCase domain.SubscriptionUseCase
}

func NewSubscriptionHandler(e *echo.Echo,
	s domain.SubscriptionUseCase,
	authMiddleware middleware.JWTConfig) {
	handler := &SubscriptionHandler{SubscriptionUseCase: s}
	e.GET("/subscriptions",
		handler.GetSubscriptions,
		middleware.JWTWithConfig(authMiddleware))
	e.GET("/subscriptions/:id",
		handler.GetSubscriptionByID,
		middleware.JWTWithConfig(authMiddleware))
	e.POST("/subscriptions/create",
		handler.CreateSubscription,
		middleware.JWTWithConfig(authMiddleware))
}

func (s *SubscriptionHandler) GetSubscriptions(c echo.Context) error {
	ctx := c.Request().Context()
	authHeader := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	res, err := s.SubscriptionUseCase.GetSubscriptionByUser(ctx, token.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch data successfully",
		res),
	)
}

func (s *SubscriptionHandler) GetSubscriptionByID(c echo.Context) error {
	ctx := c.Request().Context()
	authHeader := c.Request().Header.Get("Authorization")
	token, err := auth.NewTokenExtraction(authHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway,
			err.Error())
	}

	id := c.Param("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway,
			err.Error())
	}

	res, err := s.SubscriptionUseCase.GetSubscriptionByID(ctx, token.Email, idInteger)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError,
			err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch data successfully",
		res),
	)
}

func (s *SubscriptionHandler) CreateSubscription(c echo.Context) error {
	ctx := c.Request().Context()
	authHeader := c.Request().Header.Get("Authorization")

	token, err := auth.NewTokenExtraction(authHeader)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusBadGateway,
			err.Error())
	}

	if err := s.SubscriptionUseCase.CreateSubscriptionByUser(ctx,
		token.Email); err != nil {
		return echo.NewHTTPError(
			http.StatusInternalServerError,
			err.Error())
	}

	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"created data successfully",
		nil))
}
