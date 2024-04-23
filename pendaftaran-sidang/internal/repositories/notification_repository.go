package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type NotificationRepository interface {
	GetAll(db *gorm.DB) ([]entity.Notification, error)
	GetUserNotification(db *gorm.DB, userId int) ([]entity.Notification, error)
	GetUnreadTotal(db *gorm.DB, userId int) (int64, error)
	GetNotificationById(db *gorm.DB, notificationId string) (*entity.Notification, error)
	UpdateNotification(db *gorm.DB, notification *entity.Notification) (*entity.Notification, error)
	Save(db *gorm.DB, notification []entity.Notification) ([]entity.Notification, error)
}
