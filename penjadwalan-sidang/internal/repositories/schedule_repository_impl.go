package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"penjadwalan-sidang/internal/model/entity"
	"strconv"
	"time"
)

type ScheduleRepositoryImpl struct {
	DB *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{DB: db}
}

func (repo ScheduleRepositoryImpl) FindAll() ([]entity.Schedule, error) {
	var allSchedules []entity.Schedule
	err := repo.DB.Find(&allSchedules).Error
	if err != nil {
		return nil, err
	}

	return allSchedules, nil
}

func (repo ScheduleRepositoryImpl) FindById(scheduleId int) (*entity.Schedule, error) {
	var foundSchedule *entity.Schedule
	err := repo.DB.Take(&foundSchedule, "id = ?", scheduleId).Error
	if err != nil {
		return nil, errors.New("schedule id " + strconv.Itoa(scheduleId) + " is not found")
	}

	return foundSchedule, nil
}

func (repo ScheduleRepositoryImpl) FindScheduleByPengajuan(PengajuanId []int) ([]entity.Schedule, error) {
	var pengajuanSchedules []entity.Schedule
	err := repo.DB.Where("pengajuan_id IN ?", PengajuanId).Find(&pengajuanSchedules).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no schedules found for the provided users")
		}
		return nil, err
	}

	return pengajuanSchedules, nil
}

func (repo ScheduleRepositoryImpl) FindScheduleByPenguji(PengujiId int) ([]entity.Schedule, error) {
	var pengajuanSchedules []entity.Schedule
	err := repo.DB.Where("penguji1_id = ?", PengujiId).Or("penguji2_id = ?", PengujiId).Find(&pengajuanSchedules).Error
	if err != nil {
		return nil, errors.New("penguji is not found in schedule")
	}

	return pengajuanSchedules, nil
}

func (repo ScheduleRepositoryImpl) FindByDate(date time.Time) (*entity.Schedule, error) {
	var foundSchedule *entity.Schedule
	err := repo.DB.Take(&foundSchedule, "date_time = ?", date).Error
	if err != nil {
		return nil, errors.New("schedule data dont have match date")
	}

	return foundSchedule, nil
}

func (repo ScheduleRepositoryImpl) CheckAvailRoom(dateTime time.Time, room string, pengajuanId []int) ([]entity.Schedule, error) {
	timeEnd := dateTime.Add(2 * time.Hour)
	var foundData []entity.Schedule

	db := repo.DB.Where("ruang = ?", room).
		Where("((date_time >= ? AND date_time < ?) OR (date_time < ? AND date_time > ?))", dateTime, timeEnd, timeEnd, dateTime.Add(-2*time.Hour)).
		Find(&foundData)

	if len(pengajuanId) > 1 {
		db = db.Where("pengajuan_id NOT IN ?", pengajuanId).
			Find(&foundData)
	}

	err := db.Error
	if err != nil {
		return nil, err
	}

	return foundData, nil
}

func (repo ScheduleRepositoryImpl) CheckAvailUser(dateTime time.Time, userId int, column string, pengajuanId []int) ([]entity.Schedule, error) {
	timeEnd := dateTime.Add(2 * time.Hour)
	var foundData []entity.Schedule

	db := repo.DB.Where(column+" = ?", userId).
		Where("((date_time >= ? AND date_time < ?) OR (date_time < ? AND date_time > ?))", dateTime, timeEnd, timeEnd, dateTime.Add(-2*time.Hour)).
		Find(&foundData)

	if len(pengajuanId) > 1 {
		db = db.Where("pengajuan_id NOT IN ?", pengajuanId).
			Find(&foundData)
	}

	err := db.Error
	if err != nil {
		return nil, err
	}

	fmt.Println(foundData)

	return foundData, nil
}

func (repo ScheduleRepositoryImpl) Save(schedule []entity.Schedule) ([]entity.Schedule, error) {
	err := repo.DB.Create(&schedule).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (repo ScheduleRepositoryImpl) Update(schedule []entity.Schedule) ([]entity.Schedule, error) {
	err := repo.DB.Save(&schedule).Error
	if err != nil {
		return nil, err
	}

	return schedule, nil
}

func (repo ScheduleRepositoryImpl) Delete(scheduleId []int) error {
	err := repo.DB.Delete(&entity.Schedule{}, "pengajuan_id IN  ?", scheduleId).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo ScheduleRepositoryImpl) FindAllPagination(page int) ([]entity.Schedule, error) {
	var allSchedules []entity.Schedule
	limit := 3
	err := repo.DB.Limit(limit).Offset((page - 1) * limit).Find(&allSchedules).Order("created_at DESC").Error
	if err != nil {
		return nil, err
	}

	return allSchedules, nil
}
