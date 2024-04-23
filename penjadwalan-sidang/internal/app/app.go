package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"
	"penjadwalan-sidang/internal/config"
	"penjadwalan-sidang/internal/controllers"
	"penjadwalan-sidang/internal/repositories"
	"penjadwalan-sidang/internal/routes"
	"penjadwalan-sidang/internal/services"
)

func StartApp() {
	app := fiber.New()

	app.Use(cors.New())
	validator := validator.New()

	db, err := config.OpenConnection()
	if err != nil {
		panic(err)
	}

	scheduleRepository := repositories.NewScheduleRepository(db)
	scheduleService := services.NewScheduleService(scheduleRepository, validator)
	scheduleController := controllers.NewScheduleController(scheduleService, validator)

	api := app.Group("/api")
	routes.NewRoutes(api, scheduleController)

	err = app.Listen(":" + os.Getenv("APP_PORT"))
	if err != nil {
		panic(err)
	}
}
