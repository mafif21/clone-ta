package entity

import "time"

type Period struct {
	ID          int         `gorm:"primaryKey;column:id;autoIncrement"`
	Name        string      `gorm:"column:name"`
	StartDate   time.Time   `gorm:"column:start_date"`
	EndDate     time.Time   `gorm:"column:end_date"`
	Description string      `gorm:"column:description"`
	Pengajuans  []Pengajuan `gorm:"foreignKey:PeriodID;references:ID"`
	CreatedAt   time.Time   `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (period *Period) TableName() string {
	return "periods"
}
