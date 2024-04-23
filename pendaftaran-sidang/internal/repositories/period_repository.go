package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
	"time"
)

type PeriodRepository interface {
	FindAll(db *gorm.DB) ([]entity.Period, error)
	FindPeriodById(db *gorm.DB, periodId int) (*entity.Period, error)
	FindPeriodByName(db *gorm.DB, periodName string) (*entity.Period, error)
	FindPeriodUpdateName(db *gorm.DB, periodName string, periodId int) (*entity.Period, error)
	Save(db *gorm.DB, period *entity.Period) (*entity.Period, error)
	Update(db *gorm.DB, period *entity.Period) (*entity.Period, error)
	Delete(db *gorm.DB, periodId int) error
	FindPeriodByTime(db *gorm.DB, time time.Time) (*entity.Period, error)
}
