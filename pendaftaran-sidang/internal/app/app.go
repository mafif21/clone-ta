package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/repositories"
	"pendaftaran-sidang/internal/routes"
	"pendaftaran-sidang/internal/services"
)

func StartApp() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		},
	})
	app.Use(cors.New())

	validator := validator.New()

	// log repository
	documentLogRepository := repositories.NewDocumentLog()
	statusLogRepository := repositories.NewStatusLogRepository()

	periodRepository := repositories.NewPeriodRepository()

	//pengajuan
	pengajuanRepository := repositories.NewPengajuanRepository()
	pengajuanService := services.NewPengajuanService(pengajuanRepository, periodRepository)
	sidangController := controllers.NewPengajuanController(pengajuanService, validator)

	//period
	periodService := services.NewPeriodService(periodRepository)
	periodController := controllers.NewPeriodController(periodService, validator)

	//team
	teamRepository := repositories.NewTeamRepository()
	teamService := services.NewTeamService(teamRepository, pengajuanRepository, documentLogRepository)
	teamController := controllers.NewTeamContoller(teamService, validator)

	//document logs
	documentLogService := services.NewDocumentLogService(documentLogRepository, pengajuanRepository)
	documentLogController := controllers.NewDocumentLogController(documentLogService, pengajuanService, validator)

	//status logs
	statusLogService := services.NewStatusLogService(statusLogRepository, pengajuanRepository)
	statusLogController := controllers.NewStatusLogController(statusLogService)

	//notification
	notificationRepository := repositories.NewNotificationRepository()
	notificationService := services.NewNotificationService(notificationRepository)
	notificationController := controllers.NewNotificationController(notificationService)

	api := app.Group("api")
	app.Static("/public/doc_ta", "./public/doc_ta")
	app.Static("/public/makalah", "./public/makalah")
	app.Static("/public/slides", "./public/slides")

	routes.PengajuanRoutes(api, sidangController)
	routes.PeriodRoutes(api, periodController)
	routes.TeamRoutes(api, teamController)
	routes.SlideRoutes(api, documentLogController)
	routes.StatusLogRoutes(api, statusLogController)
	routes.NotificationRoutes(api, notificationController)

	err := app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
