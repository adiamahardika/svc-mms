package router

import (
	"os"
	"svc-monitoring-maintenance/controller"
	"svc-monitoring-maintenance/repository"
	"svc-monitoring-maintenance/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllRouter(db *gorm.DB) {
	
	router := gin.Default()
	repository := repository.Repository(db)

	ticketService := service.TicketService(repository)
	tikcetController := controller.TicketController(ticketService)

	taskListService := service.TaskListService(repository)
	taskListController := controller.TaskListController(taskListService)
	
	userService := service.UserService(repository)
	userController := controller.UserController(userService)

	teamService := service.TeamService(repository)
	teamController :=  controller.TeamController(teamService)

	roleService := service.RoleService(repository)
	roleController := controller.RoleController(roleService)
	
	dir := os.Getenv("FILE_DIR")
	v1 := router.Group("/v1")

	router.Static("/assets", dir)
	v1.GET("/get-all-ticket", tikcetController.GetAll)
	v1.GET("/get-count-ticket-status", tikcetController.CountTicketByStatus)
	v1.POST("/get-ticket", tikcetController.GetTicket)
	v1.POST("/create-ticket", tikcetController.CreateTicket)
	v1.PUT("/assign-ticket", tikcetController.AssignTicketToMember)
	v1.PUT("/update-ticket-status", tikcetController.UpdateTicketStatus)
	
	v1.POST("/get-task-list", taskListController.GetTaskList)
	v1.POST("/update-task-list", taskListController.UpdateTaskListController)
	
	v1.POST("/get-user", userController.GetUser)
	
	v1.GET("/get-team", teamController.GetAll)

	v1.GET("/get-role", roleController.GetAll)
	
	router.Run(os.Getenv("PORT"))
}