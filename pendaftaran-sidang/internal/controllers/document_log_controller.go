package controllers

import "github.com/gofiber/fiber/v2"

type DocumentLogController interface {
	FindAll(ctx *fiber.Ctx) error
	FindSlideByUserId(ctx *fiber.Ctx) error
	FindLoggedInSlide(ctx *fiber.Ctx) error
	CreateSlide(ctx *fiber.Ctx) error
}
