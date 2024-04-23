package services

import (
	"pendaftaran-sidang/internal/model/web"
)

type DocumentLogService interface {
	GetAllDocumentLog() ([]web.DocumentLogResponse, error)
	GetDocumentLogById(documentId int) (*web.DocumentLogResponse, error)
	GetUserSlide(userId int) (*web.DocumentLogResponse, error)
	Create(request *web.DocumentLogCreateRequest) (*web.DocumentLogResponse, error)
	Update(request *web.DocumentLogUpdateRequest) (*web.DocumentLogResponse, error)
	Delete(documentId int) error
}
