package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
	"time"
)

type NotificationControllerImpl struct {
	Services services.NotificationService
}

func NewNotificationController(services services.NotificationService) NotificationController {
	return &NotificationControllerImpl{Services: services}
}

func (controller NotificationControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	allDatas, err := controller.Services.GetAllNotification()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all period data",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (controller NotificationControllerImpl) FindByUser(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

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

	notificationDatas, err := controller.Services.GetUserNotification(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "data not found",
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all period data",
		Data:   notificationDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller NotificationControllerImpl) FindUnreadTotal(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

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

	total, err := controller.Services.GetUnreadNotification(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "data not found",
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "total unread notification",
		Data:   total,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller NotificationControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	notificationId := ctx.Params("notificationId")
	updateRequest := &web.NotificationUpdateRequest{
		Id:     notificationId,
		ReadAt: time.Now(),
	}

	updatedData, err := controller.Services.Update(updateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success read the notification",
		Data:   updatedData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (controller NotificationControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	createRequest := &web.NotificationCreateRequest{}
	if err := ctx.BodyParser(&createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	newData, err := controller.Services.Create(createRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to create notification",
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "success add new notification",
		Data:   newData,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
