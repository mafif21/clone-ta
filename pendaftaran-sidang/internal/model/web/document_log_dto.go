package web

import "time"

type DocumentLogCreateRequest struct {
	PengajuanId int    `validate:"required"`
	FileName    string `validate:"required"`
	Type        string `validate:"required"`
	FileUrl     string `validate:"required"`
	CreatedBy   int    `json:"created_by"`
}

type DocumentLogUpdateRequest struct {
	Id       int    `validate:"required"`
	FileName string `json:"file_name" validate:"required"`
	FileUrl  string `json:"file_url" validate:"required"`
}

type DocumentLogResponse struct {
	Id          int       `json:"id"`
	PengajuanId int       `json:"pengajuan_id"`
	FileName    string    `json:"file_name"`
	Type        string    `json:"type"`
	FileUrl     string    `json:"file_url"`
	CreatedBy   int       `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FindSlideByUserRequest struct {
	UserId int `json:"user_id"`
}
