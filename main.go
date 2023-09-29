package main

import (
	"capstone-tickets/apps/config"
	"capstone-tickets/apps/database"
	"capstone-tickets/apps/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	database.InitMigration(mysql)

	e := echo.New()
	router.InitRouter(mysql, e)
	e.Logger.Fatal(e.Start(":80"))
}
