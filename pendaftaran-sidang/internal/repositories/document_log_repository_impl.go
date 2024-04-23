package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type DocumentLogRepositoryImpl struct{}

func NewDocumentLog() DocumentLogRepository {
	return &DocumentLogRepositoryImpl{}
}

func (repository DocumentLogRepositoryImpl) FindAll(db *gorm.DB) ([]entity.DocumentLog, error) {
	var allDocuments []entity.DocumentLog

	err := db.Model(&entity.DocumentLog{}).Find(&allDocuments).Error
	if err != nil {
		return nil, err
	}

	return allDocuments, nil
}

func (repository DocumentLogRepositoryImpl) FindDocumentLogById(db *gorm.DB, documentId int) (*entity.DocumentLog, error) {
	var foundDocument entity.DocumentLog

	err := db.Take(&foundDocument, "id = ?", documentId).Error
	if err != nil {
		return nil, errors.New("document id is not found")
	}

	return &foundDocument, nil
}

func (repository DocumentLogRepositoryImpl) FindLatestDocument(db *gorm.DB, pengajuanId int, docType string) (*entity.DocumentLog, error) {
	var foundLatestDocument entity.DocumentLog
	err := db.Where("pengajuan_id = ? AND type = ?", pengajuanId, docType).Order("id DESC").First(&foundLatestDocument).Error
	if err != nil {
		return nil, errors.New("document is not found")
	}

	return &foundLatestDocument, nil
}

func (repository DocumentLogRepositoryImpl) Save(db *gorm.DB, document *entity.DocumentLog) (*entity.DocumentLog, error) {
	err := db.Create(&document).Error
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (repository DocumentLogRepositoryImpl) Update(db *gorm.DB, document *entity.DocumentLog) (*entity.DocumentLog, error) {
	err := db.Model(&entity.DocumentLog{}).Updates(&document).Error
	if err != nil {
		return nil, err
	}

	return document, nil
}

func (repository DocumentLogRepositoryImpl) Delete(db *gorm.DB, documentId int) error {
	err := db.Delete(&entity.DocumentLog{}, "id = ?", documentId).Error
	if err != nil {
		return err
	}

	return nil
}
