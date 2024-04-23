package controllers

import "github.com/gofiber/fiber/v2"

type PeriodController interface {
	FindAll(ctx *fiber.Ctx) error
	FindPeriodById(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	CheckPeriod(ctx *fiber.Ctx) error
}
