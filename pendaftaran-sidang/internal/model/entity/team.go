package entity

import "time"

type Team struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (team *Team) TableName() string {
	return "teams"
}
