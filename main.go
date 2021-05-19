package main

import (
	"fmt"
	"github.com/huf0813/scade_backend_api/infra/database/mysql"
	"github.com/huf0813/scade_backend_api/routes"
	"github.com/huf0813/scade_backend_api/utils/auth"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"time"
)

// @title Echo Swagger API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support

// @host 35.213.130.133:8080
// @BasePath /
// @schemes http
func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	authMiddleware, err := auth.NewAuthMiddleware()
	if err != nil {
		panic(err)
	}

	db, err := mysql.NewDriverMysql()
	if err != nil {
		panic(err)
	}
	timeOut := 10 * time.Second
	routes.NewRoutes(e, db, timeOut, authMiddleware)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	envAppPort := os.Getenv("APP_PORT")
	echoPort := fmt.Sprintf(":%s", envAppPort)

	e.Logger.Fatal(e.Start(echoPort))
}
