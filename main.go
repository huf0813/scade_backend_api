package main

import (
	mysql_driver "github.com/huf0813/scade_backend_api/infra/database/mysql"
	"github.com/huf0813/scade_backend_api/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db, err := mysql_driver.NewDriverMysql()
	if err != nil {
		panic(err)
	}
	routes.NewRoutes(e, db)
	e.Logger.Fatal(e.Start(":8009"))
}
