package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/middleware"
)

func NotificationRoutes(router fiber.Router, controller controllers.NotificationController) {
	period := router.Group("/notification").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	period.Get("/get", controller.FindAll)
	period.Get("/get/user", controller.FindByUser)
	period.Get("/get/user/count", controller.FindUnreadTotal)
	period.Patch("/update/:notificationId", controller.Update)
	period.Post("/create", controller.Create)

}
