package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserId    int       `gorm:"column:user_id"`
	Title     string    `gorm:"column:title"`
	Message   string    `gorm:"column:message"`
	Url       string    `gorm:"column:url"`
	ReadAt    time.Time `gorm:"column:read_at;default:null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (notification *Notification) BeforeCreate(tx *gorm.DB) error {
	notification.ID = uuid.NewString()
	return nil
}

func (notification *Notification) TableName() string {
	return "notifications"
}
