package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type PengajuanRepositoryImpl struct{}

func NewPengajuanRepository() PengajuanRepository {
	return &PengajuanRepositoryImpl{}
}

func (repository PengajuanRepositoryImpl) FindAll(db *gorm.DB, isAdmin bool, userId int) ([]entity.Pengajuan, error) {
	var allPengajuan []entity.Pengajuan
	if !isAdmin {
		err := db.Where("pembimbing1_id = ?", userId).Or("pembimbing2_id = ?", userId).Find(&allPengajuan).Error
		if err != nil {
			return nil, err
		}
		return allPengajuan, nil
	}

	err := db.Find(&allPengajuan).Error
	if err != nil {
		return nil, err
	}
	return allPengajuan, nil
}

func (repository PengajuanRepositoryImpl) FindByPengajuanId(db *gorm.DB, pengajuanId int) (*entity.Pengajuan, error) {
	var pengajuan *entity.Pengajuan
	err := db.Take(&pengajuan, "id = ?", pengajuanId).Error

	if err != nil {
		return nil, errors.New("data pengajuan not found")
	}

	return pengajuan, nil
}

func (repository PengajuanRepositoryImpl) FindPembimbingById(db *gorm.DB, userId int) ([]entity.Pengajuan, error) {
	var allPengajuan []entity.Pengajuan
	err := db.Model(&entity.Pengajuan{}).Where("pembimbing1_id = ? ", userId).Or("pembimbing2_id = ?", userId).Find(&allPengajuan).Error
	if err != nil {
		return nil, errors.New("data pengajuan not found")
	}

	return allPengajuan, nil
}

func (repository PengajuanRepositoryImpl) CheckPembimbing(db *gorm.DB, userId int, pengajuanId int) (*entity.Pengajuan, error) {
	var userPengajuan *entity.Pengajuan
	err := db.Model(&entity.Pengajuan{}).Where("pembimbing").Where("pembimbing1_id = ? ", userId).Or("pembimbing2_id = ?", userId).Take(&userPengajuan, "id = ?", pengajuanId).Error

	if err != nil {
		return nil, errors.New("pembimbing dont have permission in this data")
	}

	return userPengajuan, nil
}

func (repository PengajuanRepositoryImpl) FindByUserId(db *gorm.DB, userId int) (*entity.Pengajuan, error) {
	var userPengajuan *entity.Pengajuan
	err := db.Take(&userPengajuan, "user_id = ?", userId).Error

	if err != nil {
		return nil, errors.New("user dont have pengajuan data")
	}

	return userPengajuan, nil
}

func (repository PengajuanRepositoryImpl) FindByUsers(db *gorm.DB, users []int) ([]int, error) {
	var pengajuansId []int
	err := db.Model(&entity.Pengajuan{}).Where("user_id IN ?", users).Pluck("id", &pengajuansId).Error

	if err != nil {
		return nil, errors.New("user dont have pengajuan data")
	}

	return pengajuansId, nil
}

func (repository PengajuanRepositoryImpl) Save(db *gorm.DB, pengajuan *entity.Pengajuan) (*entity.Pengajuan, error) {
	err := db.Create(&pengajuan).Error

	if err != nil {
		return nil, err
	}

	return pengajuan, err
}

func (repository PengajuanRepositoryImpl) Update(db *gorm.DB, pengajuan *entity.Pengajuan) (*entity.Pengajuan, error) {
	err := db.Save(&pengajuan).Error
	if err != nil {
		return nil, err
	}

	return pengajuan, nil
}

func (repository PengajuanRepositoryImpl) UpdateStatus(db *gorm.DB, pengajuanId int, value string) (*entity.Pengajuan, error) {
	var pengajuan *entity.Pengajuan
	err := db.Model(&entity.Pengajuan{}).
		Where("id = ?", pengajuanId).
		Updates(map[string]interface{}{
			"status": value,
		}).Take(&pengajuan).Error
	if err != nil {
		return nil, err
	}

	return pengajuan, nil
}

func (repository PengajuanRepositoryImpl) UpdateAdminStatus(db *gorm.DB, pengajuanId int, value string, isEnglish bool) (*entity.Pengajuan, error) {
	var pengajuan *entity.Pengajuan
	err := db.Model(&entity.Pengajuan{}).
		Where("id = ?", pengajuanId).
		Updates(map[string]interface{}{
			"status":     value,
			"is_english": isEnglish,
		}).Take(&pengajuan).Error
	if err != nil {
		return nil, err
	}

	return pengajuan, nil
}

func (repository PengajuanRepositoryImpl) UpdateManyStatus(db *gorm.DB, message string, userId []int) ([]entity.Pengajuan, error) {
	err := db.Model(&entity.Pengajuan{}).Where("user_id IN ?", userId).Update("status", message).Error
	if err != nil {
		return nil, err
	}

	var allPengajuan []entity.Pengajuan
	err = db.Model(&entity.Pengajuan{}).Where("user_id IN ?", userId).Find(&allPengajuan).Error
	if err != nil {
		return nil, err
	}

	return allPengajuan, err
}

func (repository PengajuanRepositoryImpl) FindAllApprovePengajuan(db *gorm.DB) ([]entity.Pengajuan, error) {
	var allPengajuan []entity.Pengajuan

	condition := []string{"telah disetujui admin", "tidak lulus (sudah update dokumen)"}
	err := db.Where("status in ? ", condition).Find(&allPengajuan).Error
	if err != nil {
		return nil, err
	}
	return allPengajuan, nil
}

func (repository PengajuanRepositoryImpl) FindPengajuanByPeriod(db *gorm.DB, periodId int) ([]entity.Pengajuan, error) {
	var allPengajuan []entity.Pengajuan

	err := db.Joins("JOIN periods ON pengajuans.period_id = periods.id").Where(&entity.Pengajuan{PeriodID: periodId}).Find(&allPengajuan).Error
	if err != nil {
		return nil, err
	}
	return allPengajuan, nil
}
