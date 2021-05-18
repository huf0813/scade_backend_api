package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	ArticleUseCase domain.ArticleUseCase
}

func NewArticleHandler(e *echo.Echo, auc domain.ArticleUseCase) {
	handler := ArticleHandler{ArticleUseCase: auc}
	e.GET("/articles", handler.GetArticles)
	e.GET("/articles/:language", handler.GetArticlesBasedOnLanguage)
	e.GET("/articles/:language/:id", handler.GetArticlesBasedOnLanguageByID)
}

func (ah *ArticleHandler) GetArticles(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := ah.ArticleUseCase.GetArticles(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				"failed to fetch articles, please try again later",
				nil,
			),
		)
	}
	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"fetch articles successfully",
			res,
		),
	)
}

func (ah *ArticleHandler) GetArticlesBasedOnLanguage(c echo.Context) error {
	ctx := c.Request().Context()
	lang := c.Param("lang")
	res, err := ah.ArticleUseCase.GetArticlesBasedOnLanguage(ctx, lang)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil,
			),
		)
	}
	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"fetch articles successfully",
			res,
		),
	)
}

func (ah *ArticleHandler) GetArticlesBasedOnLanguageByID(c echo.Context) error {
	ctx := c.Request().Context()
	lang := c.Param("language")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil,
			),
		)
	}
	res, err := ah.ArticleUseCase.GetArticlesBasedOnLanguageByID(ctx, lang, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			custom_response.NewCustomResponse(
				false,
				err.Error(),
				nil,
			),
		)
	}
	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"fetch articles successfully",
			res,
		),
	)
}
