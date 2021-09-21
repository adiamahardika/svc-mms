package router

import (
	"svc-ticket-monitoring/controller"
	"svc-ticket-monitoring/repository"
	"svc-ticket-monitoring/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllRouter(db *gorm.DB) {
	
	router := gin.Default()
	
	ticketRepository := repository.TicketRepository(db)
	ticketService := service.TicketService(ticketRepository)
	tikcetController := controller.TicketController(ticketService)
	
	v1 := router.Group("/v1")

	v1.GET("/get-all-ticket", tikcetController.FindAll)
	v1.POST("/get-ticket", tikcetController.FindTicket)
	
	router.Run(":8888")
}