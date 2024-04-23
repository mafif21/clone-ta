package services

import (
	"errors"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
)

type NotificationServiceImpl struct {
	Repositories repositories.NotificationRepository
}

func NewNotificationService(repositories repositories.NotificationRepository) NotificationService {
	return &NotificationServiceImpl{Repositories: repositories}
}

func (service NotificationServiceImpl) GetAllNotification() ([]web.NotificationResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var allNotifications []web.NotificationResponse

	notificationDatas, err := service.Repositories.GetAll(db)
	if err != nil {
		return nil, err
	}

	for _, notification := range notificationDatas {
		allNotifications = append(allNotifications, helper.ToNotificationResponse(&notification))
	}

	return allNotifications, nil
}

func (service NotificationServiceImpl) GetUserNotification(userId int) ([]web.NotificationResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var allNotifications []web.NotificationResponse

	notificationDatas, err := service.Repositories.GetUserNotification(db, userId)
	if err != nil {
		return nil, err
	}

	for _, notification := range notificationDatas {
		allNotifications = append(allNotifications, helper.ToNotificationResponse(&notification))
	}

	return allNotifications, nil
}

func (service NotificationServiceImpl) GetUnreadNotification(userId int) (int, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return 0, errors.New("cant connect to database")
	}

	total, err := service.Repositories.GetUnreadTotal(db, userId)
	if err != nil {
		return 0, err
	}

	return int(total), nil
}

func (service NotificationServiceImpl) Update(request *web.NotificationUpdateRequest) (*web.NotificationResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundById, err := service.Repositories.GetNotificationById(db, request.Id)
	if err != nil {
		return nil, errors.New("notification not found")
	}

	foundById.ReadAt = request.ReadAt
	notificationUpdated, err := service.Repositories.UpdateNotification(db, foundById)
	response := helper.ToNotificationResponse(notificationUpdated)
	return &response, nil
}

func (service NotificationServiceImpl) Create(request *web.NotificationCreateRequest) ([]web.NotificationResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var notificationCreate []entity.Notification
	for _, uid := range request.UserId {
		newNotification := entity.Notification{
			UserId:  uid,
			Title:   request.Title,
			Message: request.Message,
			Url:     request.Url,
		}

		notificationCreate = append(notificationCreate, newNotification)
	}

	var allNotifications []web.NotificationResponse
	newNotifications, err := service.Repositories.Save(db, notificationCreate)

	for _, notification := range newNotifications {
		allNotifications = append(allNotifications, helper.ToNotificationResponse(&notification))
	}

	return allNotifications, nil
}
