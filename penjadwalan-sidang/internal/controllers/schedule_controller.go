package controllers

import "github.com/gofiber/fiber/v2"

type ScheduleController interface {
	FindAll(ctx *fiber.Ctx) error
	FindAllMahasiswa(ctx *fiber.Ctx) error
	FindAllPenguji(ctx *fiber.Ctx) error
	FindAllPembimbing(ctx *fiber.Ctx) error
	FindScheduleById(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Test(ctx *fiber.Ctx) error
}
