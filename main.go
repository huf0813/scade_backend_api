package main

import (
	"fmt"
	"github.com/huf0813/scade_backend_api/infra/database/mysql"
	"github.com/huf0813/scade_backend_api/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func main() {
	e := echo.New()

	db, err := mysql.NewDriverMysql()
	if err != nil {
		panic(err)
	}
	timeOut := 10 * time.Second
	routes.NewRoutes(e, db, timeOut)

	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	envAppPort := os.Getenv("APP_PORT")
	echoPort := fmt.Sprintf(":%s", envAppPort)

	e.Logger.Fatal(e.Start(echoPort))
}
