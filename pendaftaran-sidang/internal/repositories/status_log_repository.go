package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type StatusLogRepository interface {
	FindAll(db *gorm.DB, pengajuanId int) ([]entity.StatusLog, error)
	FindStatusLogById(db *gorm.DB, statusId int) (*entity.StatusLog, error)
	Save(db *gorm.DB, statusLog *entity.StatusLog) (*entity.StatusLog, error)
	Delete(db *gorm.DB, statusId int) error
}
