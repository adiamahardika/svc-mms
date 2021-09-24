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
	Repository := repository.Repository(db)

	ticketService := service.TicketService(Repository)
	tikcetController := controller.TicketController(ticketService)

	taskListService := service.TaskListService(Repository)
	taskListController := controller.TaskListController(taskListService)
	
	userService := service.UserService(Repository)
	userController := controller.UserController(userService)


	v1 := router.Group("/v1")

	v1.GET("/get-all-ticket", tikcetController.GetAll)
	v1.GET("/get-count-ticket-status", tikcetController.CountTicketByStatus)
	v1.POST("/get-ticket", tikcetController.GetTicket)
	v1.POST("/create-ticket", tikcetController.CreateTicket)
	
	v1.POST("/get-task-list", taskListController.GetTaskList)
	
	v1.POST("/get-user", userController.GetUser)
	
	router.Run(":8888")
}