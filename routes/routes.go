package routes

import (
	_migrationHandler "github.com/huf0813/scade_backend_api/migration/delivery/http"
	_migrationRepoMysql "github.com/huf0813/scade_backend_api/migration/repository/mysql"
	_migrationUseCase "github.com/huf0813/scade_backend_api/migration/usecase"
	"github.com/huf0813/scade_backend_api/utils/custom_response"
	"github.com/labstack/echo/v4"
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

	migrationRepoMysql := _migrationRepoMysql.NewMigrationRepoMysql(db)
	migrationUseCase := _migrationUseCase.NewMigrationUseCase(migrationRepoMysql, timeOut)
	_migrationHandler.NewMigrationHandler(e, migrationUseCase)
}
