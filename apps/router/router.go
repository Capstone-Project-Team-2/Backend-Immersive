package router

import (
	_partnerData "capstone-tickets/features/partners/data"
	_partnerHandler "capstone-tickets/features/partners/handler"
	_partnerService "capstone-tickets/features/partners/service"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})

	partnerData := _partnerData.New(db)
	partnerService := _partnerService.New(partnerData)
	partnerHandlerAPI := _partnerHandler.New(partnerService)

	c.POST("/partners", partnerHandlerAPI.Add)
}
