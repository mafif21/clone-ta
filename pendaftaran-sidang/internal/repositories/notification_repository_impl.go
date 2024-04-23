package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type NotificationRepositoryImpl struct{}

func NewNotificationRepository() NotificationRepository {
	return &NotificationRepositoryImpl{}
}

func (repo NotificationRepositoryImpl) GetAll(db *gorm.DB) ([]entity.Notification, error) {
	var allNotification []entity.Notification
	err := db.Model(&entity.Notification{}).Find(&allNotification).Error
	if err != nil {
		return nil, err
	}

	return allNotification, nil
}

func (repo NotificationRepositoryImpl) GetUserNotification(db *gorm.DB, userId int) ([]entity.Notification, error) {
	var allNotification []entity.Notification
	err := db.Model(&entity.Notification{}).Where("read_at = ?", nil).Where("user_id = ?", userId).Order("created_at DESC").Find(&allNotification).Error
	if err != nil {
		return nil, err
	}

	return allNotification, nil
}

func (repo NotificationRepositoryImpl) GetUnreadTotal(db *gorm.DB, userId int) (int64, error) {
	var totalUnread int64
	err := db.Model(&entity.Notification{}).Where("read_at = ?", nil).Where("user_id = ?", userId).Count(&totalUnread).Error
	if err != nil {
		return 0, err
	}

	return totalUnread, nil
}

func (repo NotificationRepositoryImpl) GetNotificationById(db *gorm.DB, notificationId string) (*entity.Notification, error) {
	var foundNotification entity.Notification
	err := db.Model(&entity.Notification{}).Where("id = ?", notificationId).First(&foundNotification).Error
	if err != nil {
		return nil, err
	}

	return &foundNotification, nil
}

func (repo NotificationRepositoryImpl) UpdateNotification(db *gorm.DB, notification *entity.Notification) (*entity.Notification, error) {
	err := db.Model(&entity.Notification{}).Where("id = ?", notification.ID).Updates(&notification).Error
	if err != nil {
		return nil, err
	}

	return notification, nil
}

func (repo NotificationRepositoryImpl) Save(db *gorm.DB, notification []entity.Notification) ([]entity.Notification, error) {
	err := db.Create(&notification).Error
	if err != nil {
		return nil, err
	}

	return notification, nil
}
