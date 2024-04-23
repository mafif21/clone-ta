package services

import "pendaftaran-sidang/internal/model/web"

type StatusLogService interface {
	GetUserStatusLog(userId int) ([]web.StatusLogResponse, error)
	CreateStatusLog(request web.StatusLogCreateRequest) (*web.StatusLogResponse, error)
}
