package services

import "pendaftaran-sidang/internal/model/web"

type NotificationService interface {
	GetAllNotification() ([]web.NotificationResponse, error)
	GetUserNotification(userId int) ([]web.NotificationResponse, error)
	GetUnreadNotification(userId int) (int, error)
	Update(request *web.NotificationUpdateRequest) (*web.NotificationResponse, error)
	Create(request *web.NotificationCreateRequest) ([]web.NotificationResponse, error)
}
