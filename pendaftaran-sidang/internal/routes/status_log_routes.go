package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/middleware"
)

func StatusLogRoutes(router fiber.Router, controller controllers.StatusLogController) {
	status := router.Group("/status-log")

	status.Post("/create", controller.Create)

	status.Get("/get", middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}), controller.FindStudentStatusLog)
}
