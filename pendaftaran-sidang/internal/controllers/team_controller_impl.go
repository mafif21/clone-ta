package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
	"pendaftaran-sidang/internal/exception"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/services"
	"strconv"
)

type TeamControllerImpl struct {
	Service   services.TeamService
	Validator *validator.Validate
}

func NewTeamContoller(services services.TeamService, validator *validator.Validate) TeamController {
	return &TeamControllerImpl{
		Service:   services,
		Validator: validator,
	}
}

func (controller TeamControllerImpl) GetUserTeam(ctx *fiber.Ctx) error {
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

	foundedTeam, err := controller.Service.GetTeamByUserId(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "team has found",
		Data:   foundedTeam,
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)
}

func (controller TeamControllerImpl) CreateTeam(ctx *fiber.Ctx) error {
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

	teamRequest := web.TeamCreateRequest{}
	teamRequest.UserId = userId
	if err := ctx.BodyParser(&teamRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	if err := controller.Validator.Struct(&teamRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	newTeam, err := controller.Service.Create(&teamRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "team " + newTeam.Name + " has been created",
		Data:   newTeam,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller TeamControllerImpl) CreatePersonal(ctx *fiber.Ctx) error {
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
	personalRequest := web.TeamCreateRequest{}
	personalRequest.UserId = userId

	student, errorResponse, err := helper.GetDetailStudent(userId)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	personalRequest.Name = strconv.Itoa(student.Data.Nim) + " Sidang Individu"
	newTeam, err := controller.Service.Create(&personalRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "team " + newTeam.Name + " has been created",
		Data:   newTeam,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller TeamControllerImpl) AddMember(ctx *fiber.Ctx) error {
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

	addMemberRequest := &web.AddMemberRequest{}
	//studentLoggedIn := ctx.Locals("user_id")
	//userId := studentLoggedIn.(int)

	if err := ctx.BodyParser(&addMemberRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	// nembak api
	//student, errorResponse, err := helper.GetDetailStudent(userId)
	//if errorResponse != nil {
	//	return ctx.Status(errorResponse.Code).JSON(errorResponse)
	//}

	//if err != nil {
	//	return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
	//		Code:    fiber.StatusNotFound,
	//		Message: err.Error(),
	//	})
	//}

	//addMemberRequest.TeamId = student.Data.TeamId
	err := controller.Service.AddMember(addMemberRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "success add new member",
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller TeamControllerImpl) LeaveTeam(ctx *fiber.Ctx) error {
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

	//student, errorResponse, err := helper.GetDetailStudent(userId)
	//if errorResponse != nil {
	//	return ctx.Status(errorResponse.Code).JSON(errorResponse)
	//}

	//if err != nil {
	//	return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
	//		Code:    fiber.StatusNotFound,
	//		Message: err.Error(),
	//	})
	//}

	//request := &web.LeaveMemberRequest{
	//	UserId: userId,
	//	TeamId: student.Data.TeamId,
	//}

	leaveTeamRequest := &web.LeaveMemberRequest{
		UserId: userId,
	}

	if err := ctx.BodyParser(&leaveTeamRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	err := controller.Service.LeaveTeam(leaveTeamRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "success leave the team",
	}

	return ctx.Status(fiber.StatusOK).JSON(webResponse)

}

func (controller TeamControllerImpl) Update(ctx *fiber.Ctx) error {
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
	teamId, err := ctx.ParamsInt("teamId")

	teamUpdateRequest := web.TeamUpdateRequest{}
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "team id is not valid",
		})
	}

	teamUpdateRequest.Id = teamId
	if err := ctx.BodyParser(&teamUpdateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "failed to parse JSON",
		})
	}

	if err := controller.Validator.Struct(&teamUpdateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	student, errorResponse, err := helper.GetDetailStudent(userId)
	if errorResponse != nil {
		return ctx.Status(errorResponse.Code).JSON(errorResponse)
	}

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	if student.Data.TeamId != teamUpdateRequest.Id {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusUnauthorized,
			Message: "cant update the team name",
		})
	}

	updatedTeam, err := controller.Service.Update(&teamUpdateRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusCreated,
		Status: "team name has been updated",
		Data:   updatedTeam,
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller TeamControllerImpl) Delete(ctx *fiber.Ctx) error {
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
	userLoggedIn := ctx.Locals("user_id")
	userId := userLoggedIn.(int)

	teamId, err := ctx.ParamsInt("teamId")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "team id is not valid",
		})
	}

	err = controller.Service.Delete(teamId, userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&exception.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: err.Error(),
		})
	}

	webResponse := web.WebResponse{
		Code:   fiber.StatusOK,
		Status: "team name has been deleted",
	}

	return ctx.Status(fiber.StatusCreated).JSON(webResponse)

}
