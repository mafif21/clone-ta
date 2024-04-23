package web

import (
	"time"
)

type PeriodCreateRequest struct {
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Description string    `json:"description"`
}

type PeriodUpdateRequest struct {
	Id          int       `validate:"required"`
	Name        string    `json:"name" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Description string    `json:"description" validate:"required"`
}

type PeriodResponse struct {
	Id          int                 `json:"id"`
	Name        string              `json:"name"`
	StartDate   time.Time           `json:"start_date"`
	EndDate     time.Time           `json:"end_date"`
	Description string              `json:"description"`
	Pengajuans  []PengajuanResponse `json:"pengajuan_data"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}
