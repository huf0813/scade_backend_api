package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ArticleLanguageHandler struct {
	ArticleLanguageUseCase domain.ArticleLanguageUseCase
}

func NewArticleLanguageHandler(e *echo.Echo, a domain.ArticleLanguageUseCase) {
	handler := ArticleLanguageHandler{ArticleLanguageUseCase: a}
	e.GET("/article_languages", handler.GetArticleLanguages)
}

func (a *ArticleLanguageHandler) GetArticleLanguages(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := a.ArticleLanguageUseCase.GetArticleLanguages(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"fetch data successfully",
		res),
	)
}
