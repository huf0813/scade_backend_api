package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type UserHandler struct {
	UserUseCase domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, u domain.UserUseCase) {
	handler := &UserHandler{u}
	g := e.Group("/auth")
	g.POST("/sign_in", handler.SignIn)
	g.POST("/sign_up", handler.SignUp)
}

func (u *UserHandler) SignIn(c echo.Context) error {
	user := new(domain.UserSignInRequest)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	result, err := u.UserUseCase.SignIn(ctx, user.Email, user.Password, 10, time.Hour)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"sign in successfully",
		result),
	)
}

func (u *UserHandler) SignUp(c echo.Context) error {
	user := new(domain.UserSignUpRequest)
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if err := u.UserUseCase.SignUp(ctx,
		user.Name,
		user.Address,
		user.Email,
		user.Phone,
		user.Password); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"sign up successfully",
		nil),
	)
}
