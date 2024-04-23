package web

import (
	"time"
)

type ValidationCreateSchedule struct {
	Pembimbing1Id int    `json:"pembimbing1_id" validate:"required"`
	Pembimbing2Id int    `json:"pembimbing2_id" validate:"required"`
	StudentKK     string `json:"student_kk" validate:"required"`
	Members       []int  `json:"members"`
}

type ScheduleCreateRequest struct {
	PengajuanId int       `json:"pengajuan_id" validate:"required"`
	DateTime    time.Time `json:"date_time" validate:"required"`
	Room        string    `json:"room" validate:"required"`
	Penguji1    struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji1" validate:"required"`
	Penguji2 struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji2" validate:"required"`
	ValidationCreateSchedule
}

type ScheduleResponse struct {
	ID               int       `json:"id"`
	PengajuanId      int       `json:"pengajuan_id"`
	DateTime         time.Time `json:"date_time"`
	Room             string    `json:"room"`
	Penguji1Id       int       `json:"penguji1_id"`
	Penguji2Id       int       `json:"penguji2_id"`
	Status           string    `json:"status"`
	Decision         string    `json:"decision"`
	RevisionDuration int       `json:"revision_duration"`
	FlagAddRevision  bool      `json:"flag_add_revision"`
	FlagChangeScores bool      `json:"flag_change_scores"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ScheduleUpdateRequest struct {
	Id          int       `validate:"required"`
	PengajuanId int       `json:"pengajuan_id" validate:"required"`
	DateTime    time.Time `json:"date_time" validate:"required"`
	Room        string    `json:"room" validate:"required"`
	Penguji1    struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji1" validate:"required"`
	Penguji2 struct {
		Id  int    `json:"id" validate:"required"`
		Jfa string `json:"jfa" validate:"required"`
		Kk  string `json:"kk" validate:"required"`
	} `json:"penguji2" validate:"required"`
	ValidationCreateSchedule
}

type ScheduleDeleteRequest struct {
	ScheduleId int
	Members    []int `json:"members"`
}
