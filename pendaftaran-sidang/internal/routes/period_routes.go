package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/middleware"
)

func PeriodRoutes(router fiber.Router, controller controllers.PeriodController) {
	period := router.Group("/period").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	period.Get("/get", controller.FindAll)
	period.Get("/check-period", controller.CheckPeriod)
	period.Get("/get/:periodId", controller.FindPeriodById)
	period.Post("/create", controller.Create)
	period.Patch("/update/:periodId", controller.Update)
	period.Delete("/delete/:periodId", controller.Delete)
}
