package router

import (
	"os"
	"svc-monitoring-maintenance/controller"
	"svc-monitoring-maintenance/repository"
	"svc-monitoring-maintenance/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AllRouter(db *gorm.DB) {

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "token", "request-by", "signature-key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}))

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

	categoryService := service.CategoryService(repository, repository)
	categoryController := controller.CategoryController(categoryService, logService)

	preventiveService := service.PreventiveService(repository)
	preventiveController := controller.PreventiveController(preventiveService, logService)

	taskPreventiveService := service.TaskPreventiveService(repository)
	taskPreventiveController := controller.TaskPreventiveController(taskPreventiveService, logService)

	terminalService := service.TerminalService(repository)
	terminalController := controller.TerminalController(terminalService, logService)

	grapariService := service.GrapariService(repository)
	grapariController := controller.GrapariController(grapariService, logService)

	authService := service.AuthService(repository, repository)
	authController := controller.AuthController(authService, logService)

	hwReplacementService := service.HwReplacementService(repository, repository)
	hwReplacementController := controller.HwReplacementController(hwReplacementService, logService)

	hardwareService := service.HardwareService(repository)
	hardwareController := controller.HardwareController(hardwareService, logService)

	subCategoryService := service.SubCategoryService(repository)
	subCategoryController := controller.SubCategoryController(subCategoryService, logService)

	areaService := service.AreaService(repository)
	areaController := controller.AreaController(areaService, logService)

	regionalService := service.RegionalService(repository)
	regionalController := controller.RegionalController(regionalService, logService)

	dir := os.Getenv("FILE_DIR")
	router.Static("/assets", dir)

	v3 := router.Group("/v3")
	{
		ticket := v3.Group("/ticket")
		ticket.Use(service.Authentication(), authService.Authorization())
		ticket.GET("/get-all", tikcetController.GetAll)
		ticket.POST("/get-count-status", tikcetController.CountTicketByStatus)
		ticket.POST("/get", tikcetController.GetTicket)
		ticket.GET("/get-detail/:ticket-code", tikcetController.GetDetailTicket)
		ticket.POST("/create", tikcetController.CreateTicket)
		ticket.PUT("/assign", tikcetController.AssignTicket)
		ticket.PUT("/update-status", tikcetController.UpdateTicketStatus)
		ticket.POST("/get-email-history", tikcetController.GetEmailHistory)
		ticket.PUT("/update", tikcetController.UpdateTicket)

		task_list := v3.Group("/task-list")
		task_list.Use(service.Authentication(), authService.Authorization())
		task_list.POST("/get", taskListController.GetTaskList)
		task_list.POST("/update", taskListController.UpdateTaskListController)

		auth := v3.Group("/auth")
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.GET("/refresh-token", authController.RefreshToken)

		user := v3.Group("/user")
		user.Use(service.Authentication(), authService.Authorization())
		user.POST("/get", userController.GetUser)
		user.POST("/get-detail/:user-id", userController.GetDetailUser)
		user.POST("/change-pass", userController.ChangePassword)
		user.POST("/reset-pass", userController.ResetPassword)

		team := v3.Group("/team")
		team.Use(service.Authentication(), authService.Authorization())
		team.POST("/get", teamController.GetAll)
		team.POST("/create", teamController.CreateTeam)
		team.PUT("/update", teamController.UpdateTeam)
		team.DELETE("/delete/:team-id", teamController.DeleteTeam)

		role := v3.Group("/role")
		role.Use(service.Authentication(), authService.Authorization())
		role.POST("/get", roleController.GetAll)
		role.POST("/create", roleController.CreateRole)
		role.PUT("/update", roleController.UpdateRole)
		role.DELETE("/delete/:role-id", roleController.DeleteRole)

		category := v3.Group("/category")
		category.Use(service.Authentication(), authService.Authorization())
		category.POST("/get", categoryController.GetCategory)
		category.POST("/create", categoryController.CreateCategory)
		category.PUT("/update", categoryController.UpdateCategory)
		category.DELETE("/delete/:category-id", categoryController.DeleteCategory)
		category.GET("/get-detail/:id", categoryController.GetDetailCategory)

		preventive := v3.Group("/preventive")
		preventive.Use(service.Authentication(), authService.Authorization())
		preventive.POST("/create", preventiveController.CreatePreventive)
		preventive.POST("/get", preventiveController.GetPreventive)
		preventive.PUT("/update", preventiveController.UpdatePreventive)
		preventive.GET("/get-detail/:prev-code", preventiveController.GetDetailPreventive)
		preventive.POST("/get-count-status", preventiveController.CountPreventiveByStatus)

		task_preventive := v3.Group("/task-preventive")
		task_preventive.Use(service.Authentication(), authService.Authorization())
		task_preventive.POST("/update", taskPreventiveController.UpdateTaskPreventiveController)
		task_preventive.POST("/get", taskPreventiveController.GetTaskPreventive)

		hw_replacement := v3.Group("/hw-replacement")
		hw_replacement.Use(service.Authentication(), authService.Authorization())
		hw_replacement.POST("/create", hwReplacementController.CreateHwReplacementController)
		hw_replacement.POST("/get", hwReplacementController.GetHwReplacementController)

		hardware := v3.Group("/hardware")
		hardware.Use(service.Authentication(), authService.Authorization())
		hardware.POST("/get", hardwareController.GetHardware)
		hardware.POST("/create", hardwareController.CreateHardware)
		hardware.PUT("/update", hardwareController.UpdateHardware)
		hardware.DELETE("/delete/:hw-id", hardwareController.DeleteHardware)

		sub_category := v3.Group("/sub-category")
		{
			sub_category.Use(service.Authentication(), authService.Authorization())
			sub_category.GET("/get", subCategoryController.GetSubCategory)
		}

		area := v3.Group("/area")
		{
			area.Use(service.Authentication(), authService.Authorization())
			area.POST("/get", areaController.GetArea)
		}

		regional := v3.Group("/regional")
		{
			regional.Use(service.Authentication(), authService.Authorization())
			regional.POST("/get", regionalController.GetRegional)
		}

		grapari := v3.Group("/grapari")
		{
			grapari.Use(service.Authentication(), authService.Authorization())
			grapari.POST("/get", grapariController.GetGrapari)
		}

		terminal := v3.Group("/terminal")
		{
			terminal.Use(service.Authentication(), authService.Authorization())
			terminal.POST("/get", terminalController.GetTerminal)
		}
	}

	router.Run(os.Getenv("PORT"))
}
