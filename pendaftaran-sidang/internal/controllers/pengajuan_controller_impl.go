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
	"strconv"
)

type PengajuanControllerImpl struct {
	Service   services.PengajuanService
	Validator *validator.Validate
}

func NewPengajuanController(service services.PengajuanService, validator *validator.Validate) PengajuanController {
	return &PengajuanControllerImpl{Service: service, Validator: validator}
}

func (controller PengajuanControllerImpl) FindAll(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPBB"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	pembimbingLoggedIn := ctx.Locals("user_id")
	userId := pembimbingLoggedIn.(int)

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan data",
	}

	isAdmin := slices.Contains(userRoles, "RLADM")
	if isAdmin {
		pengajuan, err := controller.Service.GetAllPengajuan(true, userId)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusInternalServerError,
				Message: "internal server error",
			})
		}

		response.Data = pengajuan
		return ctx.Status(fiber.StatusOK).JSON(response)
	}

	pengajuan, err := controller.Service.GetAllPengajuan(false, userId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	response.Data = pengajuan
	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (controller PengajuanControllerImpl) FindPengajuanById(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLADM", "RLPBB", "RLPIC"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	pembimbingLoggedIn := ctx.Locals("user_id")
	userId := pembimbingLoggedIn.(int)

	canPriority := []string{"RLADM", "RLPIC"}
	isValid := slices.ContainsFunc(canPriority, func(target string) bool {
		return slices.Contains(userRoles, target)
	})
	if isValid {
		foundData, err := controller.Service.GetPengajuanById(pengajuanId, userId, true)
		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusNotFound,
				Message: err.Error(),
			})
		}

		response := web.WebResponse{
			Code:   fiber.StatusOK,
			Status: "pengajuan data with id " + strconv.Itoa(foundData.Id),
			Data:   foundData,
		}

		return ctx.Status(fiber.StatusOK).JSON(response)
	}

	foundData, err := controller.Service.GetPengajuanById(pengajuanId, userId, false)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan data with id " + strconv.Itoa(foundData.Id),
		Data:   foundData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller PengajuanControllerImpl) FindPengajuanLoggedIn(ctx *fiber.Ctx) error {
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

	foundUserData, err := controller.Service.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "user with id " + strconv.Itoa(foundUserData.UserId) + " already registered",
		Data:   foundUserData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller PengajuanControllerImpl) FindPengajuanByUser(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLPIC", "RLPBB", "RLPGJ", "RLDSN", "RLMHS", "RLADM", "RLPPM", "RLSPR"}

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

	foundUserData, err := controller.Service.GetPengajuanByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "user with id " + strconv.Itoa(foundUserData.UserId) + " already registered",
		Data:   foundUserData,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller PengajuanControllerImpl) FindPengajuanByUsers(ctx *fiber.Ctx) error {
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

	pengajuanRequest := &web.GetPengajuanByUsersId{}
	if err := ctx.BodyParser(&pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	pengajuanId, err := controller.Service.GetPengajuanByUsers(pengajuanRequest)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success get data",
		Data:   pengajuanId,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (controller PengajuanControllerImpl) Create(ctx *fiber.Ctx) error {
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

	pengajuanRequest := web.PengajuanCreateRequest{}
	studentLoggedIn := ctx.Locals("user_id")
	studentUsername := ctx.Locals("username")
	userId := studentLoggedIn.(int)

	pengajuanRequest.UserId = userId

	if err := ctx.BodyParser(&pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	docTa, errDocTa := ctx.FormFile("doc_ta")
	makalah, errMakalah := ctx.FormFile("makalah")

	if docTa == nil || makalah == nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "must upload document",
		})
	}

	allowedExtensions := map[string]bool{
		".pdf":  true,
		".doc":  true,
		".docx": true,
	}

	docTaExt := filepath.Ext(docTa.Filename)
	makalahExt := filepath.Ext(makalah.Filename)

	if !allowedExtensions[docTaExt] || !allowedExtensions[makalahExt] {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Only .pdf, .doc, or .docx files are allowed",
		})
	}

	docTaFileName, errDocTaFileName := helper.FileHandler(errDocTa, docTa)
	makalahFileName, errMakalahFileName := helper.FileHandler(errMakalah, makalah)

	if errDocTaFileName != nil || docTaFileName == "" || errMakalahFileName != nil || makalahFileName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "file not valid",
		})
	}

	docTa.Filename = docTaFileName
	pengajuanRequest.DocTa = docTa.Filename

	makalah.Filename = makalahFileName
	pengajuanRequest.Makalah = makalah.Filename

	if err := controller.Validator.Struct(pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	pengajuanResponse, err := controller.Service.Create(&pengajuanRequest, studentUsername.(string))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	//save doc file
	_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/doc_ta/%s", docTa.Filename))
	_ = ctx.SaveFile(makalah, fmt.Sprintf("./public/makalah/%s", makalah.Filename))

	if err != nil {
		return ctx.Status(fiber.StatusCreated).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "data not found",
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "pengajuan has been created",
		Data:   pengajuanResponse,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller PengajuanControllerImpl) Update(ctx *fiber.Ctx) error {
	userRoles := ctx.Locals("role").([]string)
	rolesToCheck := []string{"RLMHS", "RLADM"}

	canAccess := slices.ContainsFunc(rolesToCheck, func(target string) bool {
		return slices.Contains(userRoles, target)
	})

	if canAccess != true {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized user",
		})
	}

	pengajuanId, err := ctx.ParamsInt("pengajuanId")
	studentUsername := ctx.Locals("username")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "id is not valid",
		})
	}

	pengajuanRequest := web.PengajuanUpdateRequest{}
	pengajuanRequest.Id = pengajuanId

	userLoggedIn := ctx.Locals("user_id")
	pengajuanRequest.UserId = userLoggedIn.(int)

	if err := ctx.BodyParser(&pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	if err := controller.Validator.Struct(pengajuanRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	docTa, docTaErr := ctx.FormFile("doc_ta")
	if docTa != nil {
		docTaNewFilename, err := helper.FileHandler(docTaErr, docTa)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "file not valid",
			})
		}

		docTa.Filename = docTaNewFilename
		pengajuanRequest.DocTa = docTa.Filename
		_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/doc_ta/%s", docTa.Filename))
	}

	makalah, makalahErr := ctx.FormFile("makalah")
	if makalah != nil {
		makalahNewFileName, err := helper.FileHandler(makalahErr, makalah)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
				Code:    fiber.StatusBadRequest,
				Message: "file not valid",
			})
		}

		makalah.Filename = makalahNewFileName
		pengajuanRequest.Makalah = makalah.Filename
		_ = ctx.SaveFile(docTa, fmt.Sprintf("./public/makalah/%s", makalah.Filename))
	}

	pengajuanResponse, err := controller.Service.Update(&pengajuanRequest, studentUsername.(string))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "pengajuan has been updated",
		Data:   pengajuanResponse,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller PengajuanControllerImpl) PengajuanRejected(ctx *fiber.Ctx) error {
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

	adminLoggedIn := ctx.Locals("user_id")
	adminId := adminLoggedIn.(int)
	pengajuanId, err := ctx.ParamsInt("pengajuanId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	statusRequest := &web.StatusAdminUpdate{
		Id:     pengajuanId,
		UserId: adminId,
		Value:  "rejected",
	}

	if err := ctx.BodyParser(&statusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	dataRejected, err := controller.Service.AdminStatus(statusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "status pengajuan " + strconv.Itoa(pengajuanId) + " has been rejected by admin",
		Data:   dataRejected,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller PengajuanControllerImpl) PengajuanApprove(ctx *fiber.Ctx) error {
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

	adminLoggedIn := ctx.Locals("user_id")
	adminId := adminLoggedIn.(int)
	pengajuanId, err := ctx.ParamsInt("pengajuanId")

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "pengajuan id is not valid",
		})
	}

	statusRequest := &web.StatusAdminUpdate{
		Id:     pengajuanId,
		UserId: adminId,
		Value:  "accept",
	}

	if err := ctx.BodyParser(&statusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	dataApprove, err := controller.Service.AdminStatus(statusRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "status pengajuan " + strconv.Itoa(pengajuanId) + " has been approve by admin",
		Data:   dataApprove,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller PengajuanControllerImpl) ChangeStatus(ctx *fiber.Ctx) error {
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

	changeStatusRequest := &web.ChangeStatusRequest{}
	if err := ctx.BodyParser(&changeStatusRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	pengajuanDatas, err := controller.Service.ChangeStatus(changeStatusRequest, userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success change pengajuan status",
		Data:   pengajuanDatas,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller PengajuanControllerImpl) FindApprovePengajuan(ctx *fiber.Ctx) error {
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

	pengajuan, err := controller.Service.GetAllApprovePengajuan()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan approve data",
		Data:   pengajuan,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)

}

func (controller PengajuanControllerImpl) FindPengajuanByPeriod(ctx *fiber.Ctx) error {
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

	periodId, err := ctx.ParamsInt("periodId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "period id is not valid",
		})
	}

	pengajuan, err := controller.Service.GetAllPengajuanByPeriod(periodId)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	response := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "all pengajuan data in period " + strconv.Itoa(periodId),
		Data:   pengajuan,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)

}
