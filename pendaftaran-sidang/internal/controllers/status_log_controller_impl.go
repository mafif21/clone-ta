package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
)

type StatusLogControllerImpl struct {
	statusLogService services.StatusLogService
}

func NewStatusLogController(statusLogService services.StatusLogService) StatusLogController {
	return &StatusLogControllerImpl{statusLogService: statusLogService}
}

func (controller StatusLogControllerImpl) FindStudentStatusLog(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	allUserStatusLog, err := controller.statusLogService.GetUserStatusLog(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all status logs",
		Data:   allUserStatusLog,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller StatusLogControllerImpl) Create(ctx *fiber.Ctx) error {
	newStatusRequest := web.StatusLogCreateRequest{}

	if err := ctx.BodyParser(&newStatusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	newStatusLog, err := controller.statusLogService.CreateStatusLog(newStatusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "new status has been created",
		Data:   newStatusLog,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
