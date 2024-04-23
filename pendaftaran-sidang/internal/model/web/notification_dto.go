package web

import (
	"time"
)

type NotificationCreateRequest struct {
	UserId  []int  `json:"user_id"`
	Title   string `json:"title"`
	Message string `json:"message"`
	Url     string `json:"url"`
}

type NotificationUpdateRequest struct {
	Id     string `validate:"required"`
	ReadAt time.Time
}

type NotificationResponse struct {
	Id        string    `json:"id"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Url       string    `json:"url"`
	ReadAt    time.Time `json:"read_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
