package http

import (
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ArticleHandler struct {
	ArticleUseCase domain.ArticleUseCase
}

func NewArticleHandler(e *echo.Echo, auc domain.ArticleUseCase) {
	handler := ArticleHandler{ArticleUseCase: auc}
	e.GET("/articles", handler.GetArticles)
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
