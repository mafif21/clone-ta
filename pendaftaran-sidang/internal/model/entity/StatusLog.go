package entity

import "time"

type StatusLog struct {
	ID           int       `gorm:"primaryKey;column:id;autoIncrement"`
	Feedback     string    `gorm:"column:feedback;type:text"`
	CreatedBy    int       `gorm:"column:created_by"`
	WorkFlowType string    `gorm:"column:type;type:enum('pengajuan', 'penjadwalan', 'sidang', 'revisi', 'lulus', 'sidang ulang', 'tidak lulus')"`
	Name         string    `gorm:"column:name"`
	PengajuanID  int       `gorm:"column:pengajuan_id"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (statusLog *StatusLog) TableName() string {
	return "status_logs"
}
