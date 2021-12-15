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

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	repository := repository.Repository(db)

	ticketService := service.TicketService(repository)
	tikcetController := controller.TicketController(ticketService)

	taskListService := service.TaskListService(repository)
	taskListController := controller.TaskListController(taskListService)

	userService := service.UserService(repository)
	userController := controller.UserController(userService)

	teamService := service.TeamService(repository)
	teamController := controller.TeamController(teamService)

	roleService := service.RoleService(repository)
	roleController := controller.RoleController(roleService)

	categoryService := service.CategoryService(repository)
	categoryController := controller.CategoryController(categoryService)

	dir := os.Getenv("FILE_DIR")
	router.Static("/assets", dir)

	v1 := router.Group("/v1")
	{
		v1.GET("/get-all-ticket", tikcetController.GetAll)
		v1.POST("/get-count-ticket-status", tikcetController.CountTicketByStatus)
		v1.POST("/get-ticket", tikcetController.GetTicket)
		v1.GET("/get-detail-ticket/:ticket-code", tikcetController.GetDetailTicket)
		v1.POST("/create-ticket", tikcetController.CreateTicket)
		v1.PUT("/assign-ticket", tikcetController.AssignTicket)
		v1.PUT("/update-ticket-status", tikcetController.UpdateTicketStatus)

		v1.POST("/get-task-list", taskListController.GetTaskList)
		v1.POST("/update-task-list", taskListController.UpdateTaskListController)

		v1.POST("/get-user", userController.GetUser)
		v1.POST("/get-user-detail/:user-id", userController.GetDetailUser)
		v1.POST("/login", userController.Login)
		v1.POST("/register", userController.Register)
		v1.POST("/change-pass", userController.ChangePassword)
		v1.POST("/reset-pass", userController.ResetPassword)

		v1.GET("/get-team", teamController.GetAll)

		v1.GET("/get-role", roleController.GetAll)
		v1.POST("/create-role", roleController.CreateRole)

		v1.POST("/get-category", categoryController.GetCategory)
		v1.POST("/create-category", categoryController.CreateCategory)
		v1.PUT("/update-category", categoryController.UpdateCategory)
		v1.DELETE("/delete-category/:category-id", categoryController.DeleteCategory)
	}

	router.Run(os.Getenv("PORT"))
}
