package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
	"strconv"
	"time"
)

type PeriodControllerImpl struct {
	Service   services.PeriodService
	Validator *validator.Validate
}

func NewPeriodController(services services.PeriodService, validator *validator.Validate) PeriodController {
	return &PeriodControllerImpl{
		Service:   services,
		Validator: validator,
	}
}

func (controller *PeriodControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	period, err := controller.Service.GetAllPeriod()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all period data",
		Data:   period,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *PeriodControllerImpl) FindPeriodById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM", "RLMHS"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	periodId, err := ctx.ParamsInt("periodId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "period id is not valid",
		})
	}

	foundData, err := controller.Service.GetPeriodById(periodId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "period with id " + strconv.Itoa(periodId),
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *PeriodControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	periodId, err := ctx.ParamsInt("periodId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "period id is not valid",
		})
	}

	err = controller.Service.Delete(periodId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "period with id " + strconv.Itoa(periodId) + " has been deleted",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *PeriodControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	periodRequest := web.PeriodCreateRequest{}
	if err := ctx.BodyParser(&periodRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	if err := controller.Validator.Struct(&periodRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	newPeriod, err := controller.Service.Create(&periodRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "new period has been created",
		Data:   newPeriod,
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (controller *PeriodControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPPM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	updateRequest := web.PeriodUpdateRequest{}
	periodId, err := ctx.ParamsInt("periodId")

	updateRequest.Id = periodId
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "period id is not valid",
		})
	}

	if err := ctx.BodyParser(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	if err := controller.Validator.Struct(&updateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	periodUpdate, err := controller.Service.Update(&updateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "new period has been updated",
		Data:   periodUpdate,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller *PeriodControllerImpl) CheckPeriod(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLMHS", "RLADM", "RLPPM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	now := time.Now()

	foundData, err := controller.Service.GetPeriodByTime(now)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "time is valid",
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
