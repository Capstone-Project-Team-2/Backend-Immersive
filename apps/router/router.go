package router

import (
	"capstone-tickets/apps/middlewares"
	_eventData "capstone-tickets/features/events/data"
	_eventHandler "capstone-tickets/features/events/handler"
	_eventService "capstone-tickets/features/events/service"

	_adminData "capstone-tickets/features/admins/data"
	_adminHandler "capstone-tickets/features/admins/handler"
	_adminService "capstone-tickets/features/admins/service"

	_partnerData "capstone-tickets/features/partners/data"
	_partnerHandler "capstone-tickets/features/partners/handler"
	_partnerService "capstone-tickets/features/partners/service"

	_buyerData "capstone-tickets/features/buyers/data"
	_buyerHandler "capstone-tickets/features/buyers/handler"
	_buyerService "capstone-tickets/features/buyers/service"

	_volunteerData "capstone-tickets/features/volunteers/data"
	_volunteerHandler "capstone-tickets/features/volunteers/handler"
	_volunteerService "capstone-tickets/features/volunteers/service"

	_transactionData "capstone-tickets/features/transactions/data"
	_transactionHandler "capstone-tickets/features/transactions/handler"
	_transactionService "capstone-tickets/features/transactions/service"
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

	volunteerData := _volunteerData.New(db)
	volunteerService := _volunteerService.New(volunteerData)
	volunteerHandlerAPI := _volunteerHandler.New(volunteerService)

	transactionData := _transactionData.New(db)
	transactionService := _transactionService.New(transactionData)
	transactionHandlerAPI := _transactionHandler.New(transactionService)

	c.POST("/partners/login", partnerHandlerAPI.Login)
	c.POST("/partners", partnerHandlerAPI.Add)
	c.GET("/partners", partnerHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.GET("/partners/:partner_id", partnerHandlerAPI.Get, middlewares.JWTMiddleware())
	c.PUT("/partners/:partner_id", partnerHandlerAPI.Update, middlewares.JWTMiddleware())
	c.DELETE("/partners/:partner_id", partnerHandlerAPI.Delete, middlewares.JWTMiddleware())

	adminData := _adminData.New(db)
	adminService := _adminService.New(adminData)
	adminHandlerAPI := _adminHandler.New(adminService)

	c.POST("/admins", adminHandlerAPI.Register, middlewares.JWTMiddleware())
	c.POST("/admins/login", adminHandlerAPI.Login)

	c.POST("/buyers/login", buyerHandlerAPI.Login)
	c.POST("/buyers", buyerHandlerAPI.Create)
	c.GET("/buyers", buyerHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.GET("/buyers/:buyer_id", buyerHandlerAPI.GetById, middlewares.JWTMiddleware())
	c.DELETE("/buyers/:buyer_id", buyerHandlerAPI.DeleteById, middlewares.JWTMiddleware())
	c.PUT("/buyers/:buyer_id", buyerHandlerAPI.UpdateById, middlewares.JWTMiddleware())

	c.POST("/volunteers/login", volunteerHandlerAPI.Login)
	c.POST("/volunteers", volunteerHandlerAPI.Create, middlewares.JWTMiddleware())
	c.GET("/volunteers/events/:event_id", volunteerHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.GET("/volunteers/:volunteer_id", volunteerHandlerAPI.GetById, middlewares.JWTMiddleware())
	c.DELETE("/volunteers/:volunteer_id", volunteerHandlerAPI.DeleteById, middlewares.JWTMiddleware())
	c.PUT("/volunteers/:volunteer_id", volunteerHandlerAPI.UpdateById, middlewares.JWTMiddleware())

	c.POST("/transactions", transactionHandlerAPI.Create, middlewares.JWTMiddleware())
	c.GET("/transactions", transactionHandlerAPI.GetById)
	c.POST("/transactions/callback", transactionHandlerAPI.Update)

	c.GET("/partners/test", partnerHandlerAPI.Test, middlewares.JWTMiddleware())

	eventData := _eventData.New(db)
	eventService := _eventService.New(eventData)
	eventHandlerAPI := _eventHandler.New(eventService)

	c.POST("/events", eventHandlerAPI.Add, middlewares.JWTMiddleware())
	c.GET("/events/:event_id", eventHandlerAPI.Get)
	c.GET("/events", eventHandlerAPI.GetAll)
	c.PUT("/events/:event_id", eventHandlerAPI.Update, middlewares.JWTMiddleware())

	c.POST("/events/test", eventHandlerAPI.Test, middlewares.JWTMiddleware())

}
