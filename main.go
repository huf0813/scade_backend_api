package main

import (
	mysql_driver "github.com/huf0813/scade_backend_api/infra/database/mysql"
	"github.com/huf0813/scade_backend_api/routes"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
	e := echo.New()
	db, err := mysql_driver.NewDriverMysql()
	if err != nil {
		panic(err)
	}
	timeOut := time.Duration(10 * time.Second)
	routes.NewRoutes(e, db, timeOut)
	e.Logger.Fatal(e.Start(":8009"))
}
