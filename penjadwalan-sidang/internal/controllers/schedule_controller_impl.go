package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"penjadwalan-sidang/internal/exception"
	"penjadwalan-sidang/internal/model/web"
	"penjadwalan-sidang/internal/services"
	"strconv"
	"strings"
)

type ScheduleControllerImpl struct {
	ScheduleService services.ScheduleService
	Validator       *validator.Validate
}

func NewScheduleController(scheduleService services.ScheduleService, validator *validator.Validate) ScheduleController {
	return &ScheduleControllerImpl{ScheduleService: scheduleService, Validator: validator}
}

func (controller ScheduleControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	allDatas, err := controller.ScheduleService.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) FindAllMahasiswa(ctx *fiber.Ctx) error {
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

	allDatas, err := controller.ScheduleService.GetAllMahasiswa(userId, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) FindAllPenguji(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPGJ"}

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

	allDatas, err := controller.ScheduleService.GetAllPenguji(userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) FindAllPembimbing(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPBB"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	allDatas, err := controller.ScheduleService.GetAllPembimbing(ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all data schedule",
		Data:   allDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) FindScheduleById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	scheduleId, err := ctx.ParamsInt("scheduleId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	foundData, err := controller.ScheduleService.GetScheduleById(scheduleId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule with id " + strconv.Itoa(scheduleId) + " founded",
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) Create(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	newScheduleRequest := web.ScheduleCreateRequest{}

	if err := ctx.BodyParser(&newScheduleRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	newScheduleRequest.Room = strings.ToLower(newScheduleRequest.Room)

	newSchedule, err := controller.ScheduleService.Create(&newScheduleRequest, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "new schedule has been created",
		Data:   newSchedule,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller ScheduleControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	updateScheduleRequest := web.ScheduleUpdateRequest{}
	if err := ctx.BodyParser(&updateScheduleRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	updateScheduleRequest.Room = strings.ToLower(updateScheduleRequest.Room)
	updateScheduleRequest.Id = scheduleId

	updatedSchedule, err := controller.ScheduleService.Update(&updateScheduleRequest, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has been update",
		Data:   updatedSchedule,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) Delete(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	deleteRequest := &web.ScheduleDeleteRequest{}

	scheduleId, err := ctx.ParamsInt("scheduleId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "schedule id is not valid",
		})
	}

	if err := ctx.BodyParser(&deleteRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	deleteRequest.ScheduleId = scheduleId

	err = controller.ScheduleService.Delete(deleteRequest, ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "schedule has been delete",
	}
	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller ScheduleControllerImpl) Test(ctx *fiber.Ctx) error {
	//token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6MTQ2LCJ1c2VybmFtZSI6ImlxYmFscyIsIm5hbWEiOiJJUUJBTCBTQU5UT1NBIiwibmltIjpudWxsLCJuaXAiOiIyMDg4MDAwMS0xIiwicm9sZSI6WyJSTFBJQyIsIlJMUEJCIiwiUkxQR0oiLCJSTERTTiJdfQ.bhlPdUYS-49Ta3s2onTSN3CicFbcm_qtt1lFziNTiro"
	//dateString := "2024-04-19T18:00:00Z"

	//dateTime, _ := time.Parse(time.RFC3339, dateString)
	//foundRoom, _ := controller.ScheduleService.CheckRoom(dateTime, "Room D")

	page := ctx.Query("pages", "1")
	pageInt, err := strconv.Atoi(page)
	if pageInt < 1 {
		pageInt = 1
	}

	pagination, err := controller.ScheduleService.GetAllPagination(pageInt)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(err)
	}

	webResponse := &web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all schedule data",
		Data:   pagination,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}
