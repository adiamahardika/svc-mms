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

	roleService := service.RoleService(repository, repository, repository)
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

	reportService := service.ReportService(repository)
	reportController := controller.ReportController(reportService, logService)

	dir := os.Getenv("FILE_DIR")
	router.Static("/assets", dir)

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
		ticket.POST("/get-email-history", tikcetController.GetEmailHistory)
		ticket.PUT("/update", tikcetController.UpdateTicket)

		task_list := v2.Group("/task-list")
		task_list.Use(service.Authentication(), authService.Authorization())
		task_list.POST("/get-task-list", taskListController.GetTaskList)
		task_list.POST("/update-task-list", taskListController.UpdateTaskListController)

		auth := v2.Group("/auth")
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.GET("/refresh-token", authController.RefreshToken)

		user := v2.Group("/user")
		{
			user.Use(service.Authentication(), authService.Authorization())
			user.POST("/get-user", userController.GetUser)
			user.POST("/get-user-detail/:user-id", userController.GetDetailUser)
			user.POST("/change-pass", userController.ChangePassword)
			user.POST("/reset-pass", userController.ResetPassword)
			user.PUT("/update-key-hp", userController.UpdateKeyHp)
		}

		team := v2.Group("/team")
		team.Use(service.Authentication(), authService.Authorization())
		team.POST("/get-team", teamController.GetAll)
		team.POST("/create-team", teamController.CreateTeam)
		team.PUT("/update-team", teamController.UpdateTeam)
		team.DELETE("/delete-team/:team-id", teamController.DeleteTeam)

		role := v2.Group("/role")
		{
			role.Use(service.Authentication(), authService.Authorization())
			role.POST("/get-role", roleController.GetAll)
			role.POST("/create-role", roleController.CreateRole)
			role.PUT("/update-role", roleController.UpdateRole)
			role.DELETE("/delete-role/:role-id", roleController.DeleteRole)
			role.GET("/get-detail/:role-id", roleController.GetDetailRole)
		}

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
		preventive.POST("/get-count-preventive-status", preventiveController.CountPreventiveByStatus)

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

		hw_replacement := v2.Group("/hw-replacement")
		hw_replacement.Use(service.Authentication(), authService.Authorization())
		hw_replacement.POST("/create", hwReplacementController.CreateHwReplacementController)
		hw_replacement.POST("/get", hwReplacementController.GetHwReplacementController)

		hardware := v2.Group("/hardware")
		hardware.Use(service.Authentication(), authService.Authorization())
		hardware.POST("/get", hardwareController.GetHardware)
		hardware.POST("/create", hardwareController.CreateHardware)
		hardware.PUT("/update", hardwareController.UpdateHardware)
		hardware.DELETE("/delete/:hw-id", hardwareController.DeleteHardware)

		sub_category := v2.Group("/sub-category")
		{
			sub_category.Use(service.Authentication(), authService.Authorization())
			sub_category.GET("/get", subCategoryController.GetSubCategory)
		}

		area := v2.Group("/area")
		{
			area.Use(service.Authentication(), authService.Authorization())
			area.POST("/get", areaController.GetArea)
		}

		regional := v2.Group("/regional")
		{
			regional.Use(service.Authentication(), authService.Authorization())
			regional.POST("/get", regionalController.GetRegional)
		}

		report := v2.Group("/report")
		{
			report.Use(service.Authentication(), authService.Authorization())
			report.POST("/get-corrective", reportController.GetReportCorrective)
			report.POST("/get-preventive", reportController.GetReportPreventive)
		}

	}

	router.Run(os.Getenv("PORT"))
}
