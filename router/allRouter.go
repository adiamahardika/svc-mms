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

	logService := service.LogService(repository)

	ticketService := service.TicketService(repository)
	tikcetController := controller.TicketController(ticketService, logService)

	taskListService := service.TaskListService(repository)
	taskListController := controller.TaskListController(taskListService, logService)

	userService := service.UserService(repository)
	userController := controller.UserController(userService, logService)

	teamService := service.TeamService(repository)
	teamController := controller.TeamController(teamService, logService)

	roleService := service.RoleService(repository)
	roleController := controller.RoleController(roleService, logService)

	categoryService := service.CategoryService(repository)
	categoryController := controller.CategoryController(categoryService, logService)

	preventiveService := service.PreventiveService(repository)
	preventiveController := controller.PreventiveController(preventiveService, logService)

	taskPreventiveService := service.TaskPreventiveService(repository)
	taskPreventiveController := controller.TaskPreventiveController(taskPreventiveService, logService)

	terminalService := service.TerminalService(repository)
	terminalController := controller.TerminalController(terminalService, logService)

	grapariService := service.GrapariService(repository)
	grapariController := controller.GrapariController(grapariService, logService)

	authService := service.AuthService(repository)

	dir := os.Getenv("FILE_DIR")
	router.Static("/assets", dir)

	// router.Use(CORSMiddleware())
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

		v1.POST("/get-team", teamController.GetAll)
		v1.POST("/create-team", teamController.CreateTeam)
		v1.PUT("/update-team", teamController.UpdateTeam)
		v1.DELETE("/delete-team/:team-id", teamController.DeleteTeam)

		v1.POST("/get-role", roleController.GetAll)
		v1.POST("/create-role", roleController.CreateRole)
		v1.PUT("/update-role", roleController.UpdateRole)
		v1.DELETE("/delete-role/:role-id", roleController.DeleteRole)

		v1.POST("/get-category", categoryController.GetCategory)
		v1.POST("/create-category", categoryController.CreateCategory)
		v1.PUT("/update-category", categoryController.UpdateCategory)
		v1.DELETE("/delete-category/:category-id", categoryController.DeleteCategory)

		v1.POST("/create-preventive", preventiveController.CreatePreventive)
		v1.POST("/get-preventive", preventiveController.GetPreventive)
		v1.PUT("/update-preventive", preventiveController.UpdatePreventive)
		v1.GET("/get-detail-preventive/:prev-code", preventiveController.GetDetailPreventive)

		v1.POST("/update-task-preventive", taskPreventiveController.UpdateTaskPreventiveController)
		v1.POST("/get-task-preventive", taskPreventiveController.GetTaskPreventive)

		v1.POST("/get-terminal", terminalController.GetTerminal)

		v1.POST("/get-grapari", grapariController.GetGrapari)
	}

	v2 := router.Group("/v2")
	{
		ticket := v2.Group("/ticket")
		ticket.Use(service.Authentication(), authService.Authorization())
		ticket.GET("/get-all-ticket", tikcetController.GetAll)
		ticket.POST("/get-count-ticket-status", tikcetController.CountTicketByStatus)
		ticket.POST("/get-ticket", tikcetController.GetTicket)
		ticket.GET("/get-detail-ticket/:ticket-code", tikcetController.GetDetailTicket)
		ticket.POST("/create-ticket", tikcetController.CreateTicket)
		ticket.PUT("/assign-ticket", tikcetController.AssignTicket)
		ticket.PUT("/update-ticket-status", tikcetController.UpdateTicketStatus)

		task_list := v2.Group("/task-list")
		task_list.Use(service.Authentication(), authService.Authorization())
		task_list.POST("/get-task-list", taskListController.GetTaskList)
		task_list.POST("/update-task-list", taskListController.UpdateTaskListController)

		auth := v2.Group("/auth")
		auth.POST("/login", userController.Login)
		auth.POST("/register", userController.Register)

		user := v2.Group("/user")
		user.Use(service.Authentication(), authService.Authorization())
		user.POST("/get-user", userController.GetUser)
		user.POST("/get-user-detail/:user-id", userController.GetDetailUser)
		user.POST("/change-pass", userController.ChangePassword)
		user.POST("/reset-pass", userController.ResetPassword)

		team := v2.Group("/team")
		team.Use(service.Authentication(), authService.Authorization())
		team.POST("/get-team", teamController.GetAll)
		team.POST("/create-team", teamController.CreateTeam)
		team.PUT("/update-team", teamController.UpdateTeam)
		team.DELETE("/delete-team/:team-id", teamController.DeleteTeam)

		role := v2.Group("/role")
		role.Use(service.Authentication(), authService.Authorization())
		role.POST("/get-role", roleController.GetAll)
		role.POST("/create-role", roleController.CreateRole)
		role.PUT("/update-role", roleController.UpdateRole)
		role.DELETE("/delete-role/:role-id", roleController.DeleteRole)

		category := v2.Group("/category")
		category.Use(service.Authentication(), authService.Authorization())
		category.POST("/get-category", categoryController.GetCategory)
		category.POST("/create-category", categoryController.CreateCategory)
		category.PUT("/update-category", categoryController.UpdateCategory)
		category.DELETE("/delete-category/:category-id", categoryController.DeleteCategory)

		preventive := v2.Group("/preventive")
		preventive.Use(service.Authentication(), authService.Authorization())
		preventive.POST("/create-preventive", preventiveController.CreatePreventive)
		preventive.POST("/get-preventive", preventiveController.GetPreventive)
		preventive.PUT("/update-preventive", preventiveController.UpdatePreventive)
		preventive.GET("/get-detail-preventive/:prev-code", preventiveController.GetDetailPreventive)

		task_preventive := v2.Group("/task-preventive")
		task_preventive.Use(service.Authentication(), authService.Authorization())
		task_preventive.POST("/update-task-preventive", taskPreventiveController.UpdateTaskPreventiveController)
		task_preventive.POST("/get-task-preventive", taskPreventiveController.GetTaskPreventive)

		terminal := v2.Group("/terminal")
		terminal.Use(service.Authentication(), authService.Authorization())
		terminal.POST("/get-terminal", terminalController.GetTerminal)

		grapari := v2.Group("/grapari")
		grapari.Use(service.Authentication(), authService.Authorization())
		grapari.POST("/get-grapari", grapariController.GetGrapari)
	}

	router.Run(os.Getenv("PORT"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// 	return
		// }

		c.Next()
	}
}
