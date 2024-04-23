package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/middleware"
)

func TeamRoutes(router fiber.Router, controller controllers.TeamController) {
	team := router.Group("/team").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	team.Get("/get-team", controller.GetUserTeam)
	team.Post("/create-individual", controller.CreatePersonal)
	team.Post("/create-team", controller.CreateTeam)
	team.Post("/leave-team", controller.LeaveTeam)
	team.Post("/add-member", controller.AddMember)
	team.Patch("/update/:teamId", controller.Update)
	team.Delete("/delete/:teamId", controller.Delete)
}
