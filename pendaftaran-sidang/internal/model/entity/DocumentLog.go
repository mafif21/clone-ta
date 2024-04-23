package entity

import "time"

type DocumentLog struct {
	ID          int       `gorm:"primaryKey;column:id;autoIncrement"`
	PengajuanID int       `gorm:"foreignKey:DocumentLogID;references:ID"`
	FileName    string    `gorm:"column:file_name"`
	Type        string    `gorm:"column:type;type:enum('makalah', 'draft', 'slide')"`
	FileUrl     string    `gorm:"column:file_url"`
	CreatedBy   int       `gorm:"column:created_by"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (documentLog *DocumentLog) TableName() string {
	return "document_logs"
}
