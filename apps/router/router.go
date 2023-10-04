package router

import (
	"capstone-tickets/apps/middlewares"
	_eventData "capstone-tickets/features/events/data"
	_eventHandler "capstone-tickets/features/events/handler"
	_eventService "capstone-tickets/features/events/service"

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

	c.POST("/partners/login", partnerHandlerAPI.Login)
	c.POST("/partners", partnerHandlerAPI.Add)
	c.GET("/partners", partnerHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.GET("/partners/:partner_id", partnerHandlerAPI.Get, middlewares.JWTMiddleware())
	c.PUT("/partners/:partner_id", partnerHandlerAPI.Update, middlewares.JWTMiddleware())
	c.DELETE("/partners/:partner_id", partnerHandlerAPI.Delete, middlewares.JWTMiddleware())

	c.GET("/partners/test", partnerHandlerAPI.Test, middlewares.JWTMiddleware())

	eventData := _eventData.New(db)
	eventService := _eventService.New(eventData)
	eventHandlerAPI := _eventHandler.New(eventService)

	c.POST("/events", eventHandlerAPI.Add, middlewares.JWTMiddleware())
	c.POST("/events/test", eventHandlerAPI.Test, middlewares.JWTMiddleware())

}
