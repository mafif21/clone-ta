package services

import (
	"penjadwalan-sidang/internal/model/web"
	"time"
)

type ScheduleService interface {
	GetAll() ([]web.ScheduleResponse, error)
	GetAllMahasiswa(userId int, token string) ([]web.ScheduleResponse, error)
	GetAllPenguji(userId int) ([]web.ScheduleResponse, error)
	GetAllPembimbing(token string) ([]web.ScheduleResponse, error)
	GetScheduleById(scheduleId int) (*web.ScheduleResponse, error)

	Create(request *web.ScheduleCreateRequest, token string) ([]web.ScheduleResponse, error)
	Update(request *web.ScheduleUpdateRequest, token string) ([]web.ScheduleResponse, error)
	Delete(request *web.ScheduleDeleteRequest, token string) error
	CheckRoom(dateTime time.Time, roomName string) ([]web.ScheduleResponse, error)
	GetAllPagination(page int) ([]web.ScheduleResponse, error)
}
