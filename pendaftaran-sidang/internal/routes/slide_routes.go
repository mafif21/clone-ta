package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/middleware"
)

func SlideRoutes(router fiber.Router, controller controllers.DocumentLogController) {
	slide := router.Group("/slide").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	slide.Get("/get", controller.FindAll)
	slide.Get("/get-latest-slide", controller.FindLoggedInSlide)
	slide.Get("/get/user/:userId", controller.FindSlideByUserId)
	slide.Post("/create-slide", controller.CreateSlide)
}
