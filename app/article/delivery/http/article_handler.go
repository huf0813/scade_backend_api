package http

import (
	"fmt"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"strconv"
)

type ArticleHandler struct {
	ArticleUseCase domain.ArticleUseCase
}

func NewArticleHandler(e *echo.Echo, auc domain.ArticleUseCase, authMiddleware middleware.JWTConfig) {
	handler := ArticleHandler{ArticleUseCase: auc}
	e.GET("/articles/image/:file", handler.GetArticleImage)
	e.GET("/articles", handler.GetArticles)
	e.GET("/articles/:language", handler.GetArticlesBasedOnLanguage)
	e.GET("/articles/:language/:id", handler.GetArticlesBasedOnLanguageByID)
	e.POST("/articles/create", handler.CreateArticle, middleware.JWTWithConfig(authMiddleware))
}

func (ah *ArticleHandler) GetArticleImage(c echo.Context) error {
	filename := c.Param("file")
	path := fmt.Sprintf("assets/article/%s", filename)
	f, err := os.Open(path)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.Stream(http.StatusOK, "image/jpg", f)
}

func (ah *ArticleHandler) GetArticles(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := ah.ArticleUseCase.GetArticles(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
	lang := c.Param("language")
	res, err := ah.ArticleUseCase.GetArticlesBasedOnLanguage(ctx, lang)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res, err := ah.ArticleUseCase.GetArticlesBasedOnLanguageByID(ctx, lang, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK,
		custom_response.NewCustomResponse(
			true,
			"fetch articles successfully",
			res,
		),
	)
}

func (ah *ArticleHandler) CreateArticle(c echo.Context) error {
	a := new(domain.ArticleRequest)
	if err := c.Bind(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(a); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	if err := ah.ArticleUseCase.CreateArticle(ctx, a.Title, a.Body, a.Thumbnail, a.ArticleLanguageID); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, custom_response.NewCustomResponse(
		true,
		"data created successfully",
		nil),
	)
}
