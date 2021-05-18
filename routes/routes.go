package routes

import (
	_articleHandler "github.com/huf0813/scade_backend_api/app/article/delivery/http"
	_articleRepoMysql "github.com/huf0813/scade_backend_api/app/article/repository/mysql"
	_articleUseCase "github.com/huf0813/scade_backend_api/app/article/usecase"
	_articleLanguageHandler "github.com/huf0813/scade_backend_api/app/article_language/delivery/http"
	_articleLanguageRepoMysql "github.com/huf0813/scade_backend_api/app/article_language/repository/mysql"
	_articleLanguageUseCase "github.com/huf0813/scade_backend_api/app/article_language/usecase"
	_migrationHandler "github.com/huf0813/scade_backend_api/migration/delivery/http"
	_migrationRepoMysql "github.com/huf0813/scade_backend_api/migration/repository/mysql"
	_migrationUseCase "github.com/huf0813/scade_backend_api/migration/usecase"

	"github.com/huf0813/scade_backend_api/utils/custom_response"

	_ "github.com/huf0813/scade_backend_api/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func NewRoutes(e *echo.Echo, db *gorm.DB, timeOut time.Duration) {
	e.GET("/", func(context echo.Context) error {
		return context.JSON(http.StatusOK, custom_response.NewCustomResponse(
			true,
			"welcome to API, please contact the maintainer",
			nil),
		)
	})
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	migrationRepoMysql := _migrationRepoMysql.NewMigrationRepoMysql(db)
	migrationUseCase := _migrationUseCase.NewMigrationUseCase(migrationRepoMysql, timeOut)
	_migrationHandler.NewMigrationHandler(e, migrationUseCase)

	articleLanguageRepoMysql := _articleLanguageRepoMysql.NewArticleLanguageRepoMysql(db)
	articleLanguageUseCase := _articleLanguageUseCase.NewArticleLanguageUseCase(articleLanguageRepoMysql, timeOut)
	_articleLanguageHandler.NewArticleLanguageHandler(e, articleLanguageUseCase)

	articleRepoMysql := _articleRepoMysql.NewArticleRepoMysql(db)
	articleUseCase := _articleUseCase.NewArticleUseCase(articleRepoMysql, timeOut)
	_articleHandler.NewArticleHandler(e, articleUseCase)
}
