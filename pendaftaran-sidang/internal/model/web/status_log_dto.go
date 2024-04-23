package web

import "time"

type StatusLogCreateRequest struct {
	Feedback     string `json:"feedback" validate:"required"`
	CreatedBy    int    `json:"created_by" validate:"required"`
	WorkFlowType string `json:"workflow_type" validate:"required"`
	Name         string `json:"name" validate:"required"`
	PengajuanId  int    `json:"pengajuan_id" validate:"required"`
}

type StatusLogResponse struct {
	Id           int       `json:"id"`
	Feedback     string    `json:"feedback"`
	CreatedBy    int       `json:"created_by"`
	WorkFlowType string    `json:"workflow_type"`
	Name         string    `json:"name"`
	PengajuanId  int       `json:"pengajuan_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
