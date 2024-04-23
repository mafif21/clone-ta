package controllers

import "github.com/gofiber/fiber/v2"

type StatusLogController interface {
	FindStudentStatusLog(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
}
