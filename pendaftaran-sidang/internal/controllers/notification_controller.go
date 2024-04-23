package controllers

import "github.com/gofiber/fiber/v2"

type NotificationController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByUser(ctx *fiber.Ctx) error
	FindUnreadTotal(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}
