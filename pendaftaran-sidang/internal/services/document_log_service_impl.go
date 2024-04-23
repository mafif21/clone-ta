package services

import (
	"errors"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
)

type DocumentLogServiceImpl struct {
	DocumentLogRepository repositories.DocumentLogRepository
	PengajuanRepository   repositories.PengajuanRepository
}

func NewDocumentLogService(documentLogRepo repositories.DocumentLogRepository, pengajuanRepo repositories.PengajuanRepository) DocumentLogService {
	return &DocumentLogServiceImpl{
		DocumentLogRepository: documentLogRepo,
		PengajuanRepository:   pengajuanRepo,
	}
}

func (service DocumentLogServiceImpl) GetAllDocumentLog() ([]web.DocumentLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var allDocuments []web.DocumentLogResponse

	documents, err := service.DocumentLogRepository.FindAll(db)
	if err != nil {
		return nil, err
	}

	for _, document := range documents {
		allDocuments = append(allDocuments, helper.ToDocumentLogResponse(&document))
	}

	return allDocuments, nil
}

func (service DocumentLogServiceImpl) GetDocumentLogById(documentId int) (*web.DocumentLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundDocument, err := service.DocumentLogRepository.FindDocumentLogById(db, documentId)
	if err != nil {
		return nil, err
	}

	respose := helper.ToDocumentLogResponse(foundDocument)
	return &respose, nil
}

func (service DocumentLogServiceImpl) GetUserSlide(userId int) (*web.DocumentLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	pengajuanId, err := service.PengajuanRepository.FindByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	document, err := service.DocumentLogRepository.FindLatestDocument(db, pengajuanId.ID, "slide")
	if err != nil {
		return nil, errors.New("user dont have any slide")
	}

	respose := helper.ToDocumentLogResponse(document)
	return &respose, nil
}

func (service DocumentLogServiceImpl) Create(request *web.DocumentLogCreateRequest) (*web.DocumentLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	newDocument := &entity.DocumentLog{
		PengajuanID: request.PengajuanId,
		FileName:    request.FileName,
		Type:        request.Type,
		FileUrl:     request.FileUrl,
		CreatedBy:   request.CreatedBy,
	}

	save, err := service.DocumentLogRepository.Save(db, newDocument)
	if err != nil {
		return nil, err
	}

	response := helper.ToDocumentLogResponse(save)
	return &response, nil
}

func (service DocumentLogServiceImpl) Update(request *web.DocumentLogUpdateRequest) (*web.DocumentLogResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundDocument, err := service.DocumentLogRepository.FindDocumentLogById(db, request.Id)
	if err != nil {
		return nil, err
	}

	foundDocument.FileName = request.FileName
	foundDocument.FileUrl = request.FileUrl

	updatedData, err := service.DocumentLogRepository.Update(db, foundDocument)
	if err != nil {
		return nil, err
	}

	response := helper.ToDocumentLogResponse(updatedData)
	return &response, nil
}

func (service DocumentLogServiceImpl) Delete(documentId int) error {
	db, err := config.OpenConnection()
	if err != nil {
		return errors.New("cant connect to database")
	}

	foundDocument, err := service.DocumentLogRepository.FindDocumentLogById(db, documentId)
	if err != nil {
		return err
	}

	err = service.Delete(foundDocument.ID)
	if err != nil {
		return err
	}

	return nil
}
