package helper

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

func AddStatusLog(db *gorm.DB, feedback string, createdBy int, workflowType string, name string, pengajuanId int) (*entity.StatusLog, error) {
	newStatusLog := &entity.StatusLog{
		Feedback:     feedback,
		CreatedBy:    createdBy,
		WorkFlowType: workflowType,
		Name:         name,
		PengajuanID:  pengajuanId,
	}

	err := db.Create(newStatusLog).Error
	if err != nil {
		return nil, err
	}

	return newStatusLog, nil
}

func AddDocumentLog(db *gorm.DB, id int, fileName string, fileType string, fileUrl string, createdBy int) (*entity.DocumentLog, error) {
	newDocLog := &entity.DocumentLog{
		PengajuanID: id,
		FileName:    fileName,
		Type:        fileType,
		FileUrl:     fileUrl,
		CreatedBy:   createdBy,
	}

	err := db.Create(newDocLog).Error
	if err != nil {
		return nil, err
	}

	return newDocLog, nil
}
