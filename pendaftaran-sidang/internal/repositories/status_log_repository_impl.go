package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type StatusLogRepositoryImpl struct{}

func NewStatusLogRepository() StatusLogRepository {
	return &StatusLogRepositoryImpl{}
}

func (repository StatusLogRepositoryImpl) FindAll(db *gorm.DB, pengajuanId int) ([]entity.StatusLog, error) {
	var allStatusLog []entity.StatusLog

	err := db.Where("pengajuan_id = ?", pengajuanId).Where("type = ? OR type = ?", "pengajuan", "penjadwalan").Order("created_at desc").Find(&allStatusLog).Error
	if err != nil {
		return nil, err
	}

	return allStatusLog, nil
}

func (repository StatusLogRepositoryImpl) FindStatusLogById(db *gorm.DB, statusId int) (*entity.StatusLog, error) {
	var foundStatus entity.StatusLog

	err := db.Take(&foundStatus, "id = ?", statusId).Error
	if err != nil {
		return nil, errors.New("status log id is not found")
	}

	return &foundStatus, nil
}

func (repository StatusLogRepositoryImpl) Save(db *gorm.DB, statusLog *entity.StatusLog) (*entity.StatusLog, error) {
	err := db.Create(&statusLog).Error
	if err != nil {
		return nil, err
	}

	return statusLog, nil
}

func (repository StatusLogRepositoryImpl) Delete(db *gorm.DB, statusId int) error {
	err := db.Delete(&entity.StatusLog{}, "id = ?", statusId).Error
	if err != nil {
		return err
	}

	return nil
}
