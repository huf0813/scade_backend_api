package routes

import (
	"github.com/go-playground/validator"
	_articleHandler "github.com/huf0813/scade_backend_api/app/article/delivery/http"
	_articleRepoMysql "github.com/huf0813/scade_backend_api/app/article/repository/mysql"
	_articleUseCase "github.com/huf0813/scade_backend_api/app/article/usecase"
	_articleLanguageHandler "github.com/huf0813/scade_backend_api/app/article_language/delivery/http"
	_articleLanguageRepoMysql "github.com/huf0813/scade_backend_api/app/article_language/repository/mysql"
	_articleLanguageUseCase "github.com/huf0813/scade_backend_api/app/article_language/usecase"
	_userHandler "github.com/huf0813/scade_backend_api/app/user/delivery/http"
	_userRepoMysql "github.com/huf0813/scade_backend_api/app/user/repository/myql"
	_userUseCase "github.com/huf0813/scade_backend_api/app/user/usecase"
	_ "github.com/huf0813/scade_backend_api/docs"
	_migrationHandler "github.com/huf0813/scade_backend_api/migration/delivery/http"
	_migrationRepoMysql "github.com/huf0813/scade_backend_api/migration/repository/mysql"
	_migrationUseCase "github.com/huf0813/scade_backend_api/migration/usecase"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewRoutes(e *echo.Echo, db *gorm.DB, timeOut time.Duration, authMiddleware middleware.JWTConfig) {
	e.Validator = &CustomValidator{validator: validator.New()}

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

	userRepoMysql := _userRepoMysql.NewUserRepoMysql(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepoMysql, timeOut)
	_userHandler.NewUserHandler(e, userUseCase)

	articleLanguageRepoMysql := _articleLanguageRepoMysql.NewArticleLanguageRepoMysql(db)
	articleLanguageUseCase := _articleLanguageUseCase.NewArticleLanguageUseCase(articleLanguageRepoMysql, timeOut)
	_articleLanguageHandler.NewArticleLanguageHandler(e, articleLanguageUseCase, authMiddleware)

	articleRepoMysql := _articleRepoMysql.NewArticleRepoMysql(db)
	articleUseCase := _articleUseCase.NewArticleUseCase(articleRepoMysql, timeOut)
	_articleHandler.NewArticleHandler(e, articleUseCase, authMiddleware)
}
