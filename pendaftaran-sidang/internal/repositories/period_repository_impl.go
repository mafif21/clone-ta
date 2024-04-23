package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
	"time"
)

type PeriodRepositoryImpl struct{}

func NewPeriodRepository() PeriodRepository {
	return &PeriodRepositoryImpl{}
}

func (repository PeriodRepositoryImpl) FindAll(db *gorm.DB) ([]entity.Period, error) {
	var periods []entity.Period
	err := db.Model(&entity.Period{}).Find(&periods).Error
	if err != nil {
		return nil, err
	}

	return periods, nil
}

func (repository PeriodRepositoryImpl) FindPeriodById(db *gorm.DB, periodId int) (*entity.Period, error) {
	var period entity.Period
	err := db.Model(&entity.Period{}).Preload("Pengajuans").Take(&period, "id = ?", periodId).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	return &period, nil
}

func (repository PeriodRepositoryImpl) FindPeriodByName(db *gorm.DB, periodName string) (*entity.Period, error) {
	var period entity.Period
	err := db.Take(&period, "name = ?", periodName).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	return &period, nil
}

func (repository PeriodRepositoryImpl) FindPeriodUpdateName(db *gorm.DB, periodName string, periodId int) (*entity.Period, error) {
	var period entity.Period
	err := db.Take(&period, "name = ?", periodName).Where("id != ?", periodId).Error
	if err != nil {
		return nil, err
	}

	return &period, nil
}

func (repository PeriodRepositoryImpl) Save(db *gorm.DB, period *entity.Period) (*entity.Period, error) {
	err := db.Create(&period).Error
	if err != nil {
		return nil, err
	}

	return period, nil
}

func (repository PeriodRepositoryImpl) Update(db *gorm.DB, period *entity.Period) (*entity.Period, error) {
	err := db.Model(&entity.Period{}).Where("id = ?", period.ID).Updates(&period).Error
	if err != nil {
		return nil, err
	}

	return period, nil
}

func (repository PeriodRepositoryImpl) Delete(db *gorm.DB, periodId int) error {
	err := db.Delete(&entity.Period{}, "id = ?", periodId).Error
	if err != nil {
		return errors.New("a foreign key constraint fails")
	}

	return nil
}

func (repository PeriodRepositoryImpl) FindPeriodByTime(db *gorm.DB, time time.Time) (*entity.Period, error) {
	var foundPeriod entity.Period
	err := db.Where("? BETWEEN start_date AND end_date", time).Take(&foundPeriod).Error
	if err != nil {
		return nil, errors.New("time is not valid with period data")
	}

	return &foundPeriod, nil
}
