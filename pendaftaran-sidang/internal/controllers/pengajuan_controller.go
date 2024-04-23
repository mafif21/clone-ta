package controllers

import "github.com/gofiber/fiber/v2"

type PengajuanController interface {
	FindAll(ctx *fiber.Ctx) error
	FindPengajuanById(ctx *fiber.Ctx) error
	FindPengajuanLoggedIn(ctx *fiber.Ctx) error
	FindPengajuanByUser(ctx *fiber.Ctx) error
	FindPengajuanByUsers(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	PengajuanRejected(ctx *fiber.Ctx) error
	PengajuanApprove(ctx *fiber.Ctx) error
	ChangeStatus(ctx *fiber.Ctx) error
	FindApprovePengajuan(ctx *fiber.Ctx) error
	FindPengajuanByPeriod(ctx *fiber.Ctx) error
}
