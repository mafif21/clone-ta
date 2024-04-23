package services

import (
	"errors"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
)

type StatusLogServiceImpl struct {
	StatusLogRepository repositories.StatusLogRepository
	PengajuanRepository repositories.PengajuanRepository
}

func NewStatusLogService(statusLogRepository repositories.StatusLogRepository, pengajuanRepository repositories.PengajuanRepository) StatusLogService {
	return &StatusLogServiceImpl{StatusLogRepository: statusLogRepository, PengajuanRepository: pengajuanRepository}
}

func (service StatusLogServiceImpl) GetUserStatusLog(userId int) ([]web.StatusLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundUserPengajuan, err := service.PengajuanRepository.FindByUserId(db, userId)
	if err != nil {
		return nil, errors.New("user dont have pengajuan request")
	}

	statusLogs, err := service.StatusLogRepository.FindAll(db, foundUserPengajuan.ID)
	if err != nil {
		return nil, err
	}

	var allStatusLogs []web.StatusLogResponse
	for _, status := range statusLogs {
		allStatusLogs = append(allStatusLogs, helper.ToStatusLogResponse(&status))
	}

	return allStatusLogs, nil
}

func (service StatusLogServiceImpl) CreateStatusLog(request web.StatusLogCreateRequest) (*web.StatusLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	_, err = service.PengajuanRepository.FindByPengajuanId(db, request.PengajuanId)
	if err != nil {
		return nil, err
	}

	newStatusLog := &entity.StatusLog{
		Feedback:     request.Feedback,
		CreatedBy:    request.CreatedBy,
		WorkFlowType: request.WorkFlowType,
		Name:         request.Name,
		PengajuanID:  request.PengajuanId,
	}

	save, err := service.StatusLogRepository.Save(db, newStatusLog)
	if err != nil {
		return nil, errors.New("cant add new status logs")
	}

	response := helper.ToStatusLogResponse(save)
	return &response, nil
}
