package web

import (
	"time"
)

type PengajuanCreateRequest struct {
	UserId         int    `validate:"required"`
	Nim            string `form:"nim" json:"nim" validate:"required"`
	Pembimbing1Id  int    `form:"pembimbing1_id" json:"pembimbing1_id" validate:"required"`
	Pembimbing2Id  int    `form:"pembimbing2_id" json:"pembimbing2_id" validate:"required"`
	Judul          string `form:"judul" json:"judul"  validate:"required"`
	Eprt           int    `form:"eprt" json:"eprt"  validate:"required"`
	DocTa          string `form:"doc_ta" json:"doc_ta"`
	Makalah        string `form:"makalah" json:"makalah"`
	Tak            int    `form:"tak" json:"tak"  validate:"required"`
	PeriodID       int    `form:"period_id" json:"period_id"  validate:"required"`
	FormBimbingan1 int    `form:"form_bimbingan1" json:"form_bimbingan1"`
	FormBimbingan2 int    `form:"form_bimbingan2" json:"form_bimbingan2"`
	Peminatan      int    `form:"peminatan" json:"peminatan"  validate:"required"`
}

type PengajuanResponse struct {
	Id             int       `json:"id"`
	UserId         int       `json:"user_id"`
	Nim            string    `json:"nim"`
	Pembimbing1Id  int       `json:"pembimbing1_id"`
	Pembimbing2Id  int       `json:"pembimbing2_id"`
	Judul          string    `json:"judul"`
	Eprt           int       `json:"eprt"`
	DocTa          string    `json:"doc_ta"`
	Makalah        string    `json:"makalah"`
	Tak            int       `json:"tak"`
	Status         string    `json:"status"`
	SksLulus       int       `json:"sks_lulus"`
	SksBelumLulus  int       `json:"sks_belum_lulus"`
	IsEnglish      bool      `json:"is_english"`
	PeriodID       int       `json:"period_id"`
	SkPenguji      string    `json:"sk_penguji"`
	FormBimbingan1 int       `json:"form_bimbingan1"`
	FormBimbingan2 int       `json:"form_bimbingan2"`
	CreatedAt      time.Time `json:"created_at"`
	Updated_at     time.Time `json:"updated_at"`
}

type PengajuanUpdateRequest struct {
	Id             int    `validate:"required"`
	UserId         int    `validate:"required"`                                             //clear
	Nim            string `form:"nim" json:"nim" validate:"required"`                       //clear
	Pembimbing1Id  int    `form:"pembimbing1_id" json:"pembimbing1_id" validate:"required"` //clear
	Pembimbing2Id  int    `form:"pembimbing2_id" json:"pembimbing2_id" validate:"required"` //clear
	Judul          string `form:"judul" json:"judul" validate:"required"`                   //clear
	Eprt           int    `form:"eprt" json:"eprt" validate:"required"`                     //clear
	DocTa          string `form:"doc_ta" json:"doc_ta"`                                     //clear
	Makalah        string `form:"makalah" json:"makalah"`                                   //clear (file)
	Tak            int    `form:"tak" json:"tak" validate:"required"`                       //clear
	PeriodID       int    `form:"period_id" json:"period_id" validate:"required"`           //clear
	FormBimbingan1 int    `form:"form_bimbingan1" json:"form_bimbingan_1"`                  //clear
	FormBimbingan2 int    `form:"form_bimbingan2" json:"form_bimbingan_2"`                  //clear
	Peminatan      int    `form:"peminatan" json:"peminatan" validate:"required"`           //clear
}

type GetPengajuanByUsersId struct {
	UsersId []int `json:"users_id"`
}

type StatusAdminUpdate struct {
	Id        int
	Feedback  string `json:"feedback" validate:"required"`
	IsEnglish bool   `json:"is_english" validate:"required"`
	UserId    int
	Value     string
}

type ChangeStatusRequest struct {
	Feedback     string `json:"feedback"`
	Status       string `json:"status"`
	WorkFlowType string `json:"workflow_type"`
	MemberId     []int  `json:"member_id"`
}
