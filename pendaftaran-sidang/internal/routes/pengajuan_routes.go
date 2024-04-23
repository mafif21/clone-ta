package routes

import (
	"github.com/gofiber/fiber/v2"
	"pendaftaran-sidang/internal/controllers"
	"pendaftaran-sidang/internal/middleware"
)

func PengajuanRoutes(router fiber.Router, controller controllers.PengajuanController) {
	pengajuan := router.Group("/pengajuan").Use(middleware.UserAuthentication(middleware.AuthConfig{
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	}))

	pengajuan.Get("/get", controller.FindAll)
	pengajuan.Get("/user/:userId", controller.FindPengajuanByUser)
	pengajuan.Post("/users", controller.FindPengajuanByUsers)
	pengajuan.Get("/status/approve", controller.FindApprovePengajuan)
	pengajuan.Get("/get/:pengajuanId", controller.FindPengajuanById)
	pengajuan.Get("/period/:periodId", controller.FindPengajuanByPeriod)
	pengajuan.Post("/create", controller.Create)
	pengajuan.Patch("/update/:pengajuanId", controller.Update)
	pengajuan.Get("/check-user", controller.FindPengajuanLoggedIn)
	pengajuan.Patch("/change-status", controller.ChangeStatus)

	pengajuan.Patch("/rejected/:pengajuanId", controller.PengajuanRejected)
	pengajuan.Patch("/approve/:pengajuanId", controller.PengajuanApprove)
}
