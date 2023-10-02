package router

import (
	_partnerData "capstone-tickets/features/partners/data"
	_partnerHandler "capstone-tickets/features/partners/handler"
	_partnerService "capstone-tickets/features/partners/service"

	_buyerData "capstone-tickets/features/buyers/data"
	_buyerHandler "capstone-tickets/features/buyers/handler"
	_buyerService "capstone-tickets/features/buyers/service"
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

	buyerData := _buyerData.New(db)
	buyerService := _buyerService.New(buyerData)
	buyerHandlerAPI := _buyerHandler.New(buyerService)

	c.POST("/partners/login", partnerHandlerAPI.Login)
	c.POST("/partners", partnerHandlerAPI.Add)
	c.GET("/partners", partnerHandlerAPI.GetAll)
	c.GET("/partners/:partner_id", partnerHandlerAPI.Get)
	c.DELETE("/partners/:partner_id", partnerHandlerAPI.Delete)

	c.POST("/buyers/login", buyerHandlerAPI.Login)
	c.POST("/buyers", buyerHandlerAPI.Create)
	// c.GET("/buyers", buyerHandlerAPI.GetAll)
	// c.GET("/buyers/:buyer_id", buyerHandlerAPI.Get)
	// c.DELETE("/buyers/:buyer_id", buyerHandlerAPI.Delete)

}
