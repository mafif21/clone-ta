package services

import (
	"pendaftaran-sidang/internal/model/web"
	"time"
)

type PeriodService interface {
	GetAllPeriod() ([]web.PeriodResponse, error)
	GetPeriodById(periodId int) (*web.PeriodResponse, error)
	Create(request *web.PeriodCreateRequest) (*web.PeriodResponse, error)
	Update(request *web.PeriodUpdateRequest) (*web.PeriodResponse, error)
	Delete(periodId int) error
	GetPeriodByTime(time time.Time) (*web.PeriodResponse, error)
}
