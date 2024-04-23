package controllers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"path/filepath"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
)

type DocumentLogControllerImpl struct {
	DocumentLogService services.DocumentLogService
	PengajuanService   services.PengajuanService
	Validator          *validator.Validate
}

func NewDocumentLogController(documentLogService services.DocumentLogService, pengajuanService services.PengajuanService, validator *validator.Validate) DocumentLogController {
	return &DocumentLogControllerImpl{
		DocumentLogService: documentLogService,
		PengajuanService:   pengajuanService,
		Validator:          validator,
	}
}

func (controller DocumentLogControllerImpl) FindAll(ctx *fiber.Ctx) error {
	allDocumentLogs, err := controller.DocumentLogService.GetAllDocumentLog()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all document logs",
		Data:   allDocumentLogs,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller DocumentLogControllerImpl) FindLoggedInSlide(ctx *fiber.Ctx) error {
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

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	slideDocument, err := controller.DocumentLogService.GetUserSlide(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "get the latest slide",
		Data:   slideDocument,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller DocumentLogControllerImpl) FindSlideByUserId(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLADM", "RLPPM", "RLSPR"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	userId, err := ctx.ParamsInt("userId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "user id is not valid",
		})
	}

	slideDocument, err := controller.DocumentLogService.GetUserSlide(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "get the latest slide",
		Data:   slideDocument,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller DocumentLogControllerImpl) CreateSlide(ctx *fiber.Ctx) error {
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

	studentLoggedIn := ctx.Locals("user_id")
	userId := studentLoggedIn.(int)

	foundPengajuanUser, err := controller.PengajuanService.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	//student, err := helper.GetDetailStudent(userId)
	//if foundPengajuanUser.Status != "tidak lulus (belum dijadwalkan)" || foundPengajuanUser.Status != "belum dijadwalkan" {
	//	return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
	//		Code:    fiber.StatusBadRequest,
	//		Message: "make sure student status 'telah disetujui admin' and student have team or individual sidang",
	//	})
	//}

	if foundPengajuanUser.Status == "belum disetujui admin" || foundPengajuanUser.Status == "ditolak oleh admin" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "make sure student status 'telah disetujui admin'",
		})
	}

	requestSlide := web.DocumentLogCreateRequest{
		PengajuanId: foundPengajuanUser.Id,
		Type:        "slide",
		CreatedBy:   foundPengajuanUser.UserId,
	}

	slide, err := ctx.FormFile("slide")
	slideName, errSlideName := helper.FileHandler(err, slide)

	if errSlideName != nil || slideName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "file not valid",
		})
	}

	ext := filepath.Ext(slide.Filename)
	if ext != ".ppt" && ext != ".pptx" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "only .ppt or .pptx files are allowed",
		})
	}

	slide.Filename = slideName
	requestSlide.FileName = slide.Filename
	requestSlide.FileUrl = fmt.Sprintf("/public/slides/%s", slide.Filename)

	if err := controller.Validator.Struct(&requestSlide); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	slideResponse, err := controller.DocumentLogService.Create(&requestSlide)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	//save doc file
	_ = ctx.SaveFile(slide, fmt.Sprintf("./public/slides/%s", slide.Filename))
	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "slide has been upload",
		Data:   slideResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
