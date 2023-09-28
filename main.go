package main

import (
	"capstone-tickets/app/config"
	"capstone-tickets/app/database"
	"capstone-tickets/app/router"

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
