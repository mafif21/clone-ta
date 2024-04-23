package routes

import (
	"github.com/gofiber/fiber/v2"
	"penjadwalan-sidang/internal/controllers"
	"penjadwalan-sidang/internal/middleware"
)

func NewRoutes(router fiber.Router, scheduleController controllers.ScheduleController) {
	schedule := router.Group("/schedule").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	schedule.Get("/get", scheduleController.FindAll)
	schedule.Get("/get/mahasiswa", scheduleController.FindAllMahasiswa)
	schedule.Get("/get/penguji", scheduleController.FindAllPenguji)
	schedule.Get("/get/pembimbing", scheduleController.FindAllPembimbing)

	schedule.Get("/get/:scheduleId", scheduleController.FindScheduleById)
	schedule.Post("/create", scheduleController.Create)
	schedule.Patch("/edit/:scheduleId", scheduleController.Update)
	schedule.Delete("/delete/:scheduleId", scheduleController.Delete)
	schedule.Get("/test", scheduleController.Test)
}
