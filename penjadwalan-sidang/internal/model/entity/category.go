package entity

import "github.com/google/uuid"

// example
type Category struct {
	ID   uuid.UUID `gorm:"primaryKey;column:id;autoIncrement;type:uuid"`
	Name string    `gorm:"column:name"`
}
