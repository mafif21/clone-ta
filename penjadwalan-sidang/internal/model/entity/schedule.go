package entity

import (
	"time"
)

type Schedule struct {
	ID               int       `gorm:"primaryKey;column:id;autoIncrement"`
	PengajuanId      int       `gorm:"column:pengajuan_id"`
	DateTime         time.Time `gorm:"column:date_time"`
	Room             string    `gorm:"column:ruang"`
	Penguji1Id       int       `gorm:"column:penguji1_id"`
	Penguji2Id       int       `gorm:"column:penguji2_id"`
	Status           string    `gorm:"column:status;default:belum dilaksanakan"`
	Decision         string    `gorm:"column:keputusan"`
	RevisionDuration int       `gorm:"column:durasi_revisi"`
	FlagAddRevision  bool      `gorm:"column:flag_add_revision"`
	FlagChangeScores bool      `gorm:"column:flag_change_scores"`
	CreatedAt        time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (schedule *Schedule) TableName() string {
	return "schedules"
}
