package helper

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

func AddNotification(db *gorm.DB, userId int, title string, message string, url string) (*entity.Notification, error) {
	newNotification := &entity.Notification{
		UserId:  userId,
		Title:   title,
		Message: message,
		Url:     url,
	}

	err := db.Create(newNotification).Error
	if err != nil {
		return nil, err
	}

	return newNotification, nil
}
