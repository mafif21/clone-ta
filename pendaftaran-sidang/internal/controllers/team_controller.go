package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type TeamController interface {
	GetUserTeam(ctx *fiber.Ctx) error
	CreateTeam(ctx *fiber.Ctx) error     // create team
	CreatePersonal(ctx *fiber.Ctx) error // create personal
	AddMember(ctx *fiber.Ctx) error
	LeaveTeam(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	//GetAvailableMember(ctx *fiber.Ctx) error
}
