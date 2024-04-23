package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type DocumentLogRepository interface {
	FindAll(db *gorm.DB) ([]entity.DocumentLog, error)
	FindDocumentLogById(db *gorm.DB, documentId int) (*entity.DocumentLog, error)
	FindLatestDocument(db *gorm.DB, pengajuanId int, docType string) (*entity.DocumentLog, error)
	Save(db *gorm.DB, document *entity.DocumentLog) (*entity.DocumentLog, error)
	Update(db *gorm.DB, document *entity.DocumentLog) (*entity.DocumentLog, error)
	Delete(db *gorm.DB, documentId int) error
}
