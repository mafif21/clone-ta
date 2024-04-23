package repositories

import (
	"penjadwalan-sidang/internal/model/entity"
	"time"
)

type ScheduleRepository interface {
	FindAll() ([]entity.Schedule, error)
	FindById(scheduleId int) (*entity.Schedule, error)
	CheckAvailRoom(dateTime time.Time, room string, pengajuanId []int) ([]entity.Schedule, error)
	FindScheduleByPengajuan(pengajuanId []int) ([]entity.Schedule, error)
	FindScheduleByPenguji(pengujiId int) ([]entity.Schedule, error)
	CheckAvailUser(dateTime time.Time, userId int, column string, pengajuanId []int) ([]entity.Schedule, error)
	FindByDate(date time.Time) (*entity.Schedule, error)
	Save(schedules []entity.Schedule) ([]entity.Schedule, error)
	Update(schedule []entity.Schedule) ([]entity.Schedule, error)
	Delete(scheduleId []int) error
	FindAllPagination(page int) ([]entity.Schedule, error)
}
