package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"penjadwalan-sidang/internal/exception"
	"penjadwalan-sidang/internal/helper"
	"penjadwalan-sidang/internal/model/entity"
	"penjadwalan-sidang/internal/model/web"
	"penjadwalan-sidang/internal/repositories"
	"sync"
	"time"
)

type ScheduleServiceImpl struct {
	ScheduleRepository repositories.ScheduleRepository
	Validator          *validator.Validate
}

func NewScheduleService(scheduleRepository repositories.ScheduleRepository, validator *validator.Validate) ScheduleService {
	return &ScheduleServiceImpl{ScheduleRepository: scheduleRepository, Validator: validator}
}

func (service ScheduleServiceImpl) GetAll() ([]web.ScheduleResponse, error) {
	var allSchedulesResponse []web.ScheduleResponse

	schedules, err := service.ScheduleRepository.FindAll()
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "INTERNAL SERVER ERROR")
	}

	for _, schedule := range schedules {
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleResponse(&schedule))
	}

	return allSchedulesResponse, nil
}

func (service ScheduleServiceImpl) GetAllMahasiswa(userId int, token string) ([]web.ScheduleResponse, error) {
	//var allSchedulesResponse []web.ScheduleResponse
	//
	//wg := &sync.WaitGroup{}
	//var pengajuansId []int
	//pengajuanChannel := make(chan []int, 1)
	//
	//wg.Add(1)
	//go helper.GetPengajuanUsers(pengajuanChannel, []int{userId}, token, wg)
	//
	//go func() {
	//	wg.Wait()
	//	close(pengajuanChannel)
	//}()
	//
	//pengajuansId = <-pengajuanChannel
	//membersSchedule, err := service.ScheduleRepository.FindScheduleByPengajuan(pengajuansId)
	//if err != nil {
	//	return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "INTERNAL SERVER ERROR")
	//}
	//
	//for _, schedule := range membersSchedule {
	//	allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleResponse(&schedule))
	//}
	//return allSchedulesResponse, nil
	panic("")
}

func (service ScheduleServiceImpl) GetAllPenguji(userId int) ([]web.ScheduleResponse, error) {
	var allSchedulesResponse []web.ScheduleResponse

	pengujiSchedules, err := service.ScheduleRepository.FindScheduleByPenguji(userId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, err.Error())
	}

	for _, schedule := range pengujiSchedules {
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleResponse(&schedule))
	}

	return allSchedulesResponse, nil
}

func (service ScheduleServiceImpl) GetAllPembimbing(token string) ([]web.ScheduleResponse, error) {
	var allSchedulesResponse []web.ScheduleResponse

	wg := &sync.WaitGroup{}
	var pengajuansId []int
	pengajuanChannel := make(chan []int)

	wg.Add(1)
	go helper.GetAllPengajuan(pengajuanChannel, token, wg)

	go func() {
		wg.Wait()
		close(pengajuanChannel)
	}()

	pengajuansId = <-pengajuanChannel
	membersSchedule, err := service.ScheduleRepository.FindScheduleByPengajuan(pengajuansId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "INTERNAL SERVER ERROR")
	}

	for _, schedule := range membersSchedule {
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleResponse(&schedule))
	}
	return allSchedulesResponse, nil
}

func (service ScheduleServiceImpl) GetScheduleById(scheduleId int) (*web.ScheduleResponse, error) {
	foundData, err := service.ScheduleRepository.FindById(scheduleId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	response := helper.ToScheduleResponse(foundData)
	return &response, nil
}

func (service ScheduleServiceImpl) Create(request *web.ScheduleCreateRequest, token string) ([]web.ScheduleResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	if time.Now().After(request.DateTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "time has passed")
	}

	validTime := helper.GetValidSidangTime()
	if request.DateTime.Before(validTime) || request.DateTime.Equal(validTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "scheduling must be exactly 2 hours in advance")
	}

	if request.Penguji1.Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 must be different")
	}

	if request.StudentKK != request.Penguji1.Kk || request.StudentKK != request.Penguji2.Kk {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "there must be at least 1 penguji from the same kk")
	}

	if request.Penguji1.Jfa != "NJFA" {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 must have jfa")
	}

	if request.Pembimbing1Id == request.Penguji1.Id || request.Pembimbing1Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 cant same with pembimbing 1")
	}

	if request.Pembimbing2Id == request.Penguji1.Id || request.Pembimbing2Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 cant same with pembimbing 2")
	}

	room, _ := service.ScheduleRepository.CheckAvailRoom(request.DateTime, request.Room, []int{request.PengajuanId})
	if len(room) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "room is not available")
	}

	foundDataPenguji1, _ := service.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji1.Id, "penguji1_id", []int{request.PengajuanId})
	if len(foundDataPenguji1) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 is not available")
	}

	foundDataPenguji2, _ := service.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji2.Id, "penguji2_id", []int{request.PengajuanId})
	if len(foundDataPenguji2) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 2 is not available")
	}

	var scheduleCreateData []entity.Schedule
	wg := &sync.WaitGroup{}
	pengajuanChannel := make(chan *web.GetPengajuanDatas, len(request.Members))

	wg.Add(1)
	go helper.ChangePengajuanStatus(pengajuanChannel, "-", "sudah dijadwalkan", "sidang", request.Members, token, wg)

	go func() {
		wg.Wait()
		close(pengajuanChannel)
	}()

	counter := 0
	for {
		select {
		case pengajuan := <-pengajuanChannel:
			for _, data := range pengajuan.Data {
				newSchedule := entity.Schedule{
					PengajuanId: data.PengajuanId,
					DateTime:    request.DateTime,
					Room:        request.Room,
					Penguji1Id:  request.Penguji1.Id,
					Penguji2Id:  request.Penguji2.Id,
				}

				scheduleCreateData = append(scheduleCreateData, newSchedule)
				counter++
			}

		}

		if counter == len(request.Members) {
			break
		}
	}

	savedNewSchedule, err := service.ScheduleRepository.Save(scheduleCreateData)
	if err != nil {
		return nil, &exception.ErrorResponse{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	helper.CreateNotification(request.Members, "Penjadwalan Sidang", "Jadwal sidang anda sudah ditetapkan!", "/schedule/get/mahasiswa", token)
	var newScheduleResponse []web.ScheduleResponse
	for _, newSavedSchedule := range savedNewSchedule {
		newScheduleResponse = append(newScheduleResponse, helper.ToScheduleResponse(&newSavedSchedule))
	}
	return newScheduleResponse, nil
}

func (service ScheduleServiceImpl) Update(request *web.ScheduleUpdateRequest, token string) ([]web.ScheduleResponse, error) {
	if err := service.Validator.Struct(request); err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, err.Error())
	}

	_, err := service.ScheduleRepository.FindById(request.Id)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	if time.Now().After(request.DateTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "time has passed")
	}

	validTime := helper.GetValidSidangTime()
	if request.DateTime.Before(validTime) || request.DateTime.Equal(validTime) {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "scheduling must be exactly 2 hours in advance")
	}

	if request.Penguji1.Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 must be different")
	}

	if request.StudentKK != request.Penguji1.Kk || request.StudentKK != request.Penguji2.Kk {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "there must be at least 1 penguji from the same kk")
	}

	if request.Penguji1.Jfa != "NJFA" {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 must have jfa")
	}

	if request.Pembimbing1Id == request.Penguji1.Id || request.Pembimbing1Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 cant same with pembimbing 1")
	}

	if request.Pembimbing2Id == request.Penguji1.Id || request.Pembimbing2Id == request.Penguji2.Id {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 and penguji 2 cant same with pembimbing 2")
	}

	pengajuansId := helper.GetPengajuanUsers(request.Members, token)
	fmt.Println(pengajuansId)

	room, _ := service.ScheduleRepository.CheckAvailRoom(request.DateTime, request.Room, pengajuansId)
	if len(room) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "room is not available")
	}

	foundDataPenguji1, _ := service.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji1.Id, "penguji1_id", pengajuansId)
	if len(foundDataPenguji1) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 1 is not available")
	}

	foundDataPenguji2, _ := service.ScheduleRepository.CheckAvailUser(request.DateTime, request.Penguji2.Id, "penguji2_id", pengajuansId)
	if len(foundDataPenguji2) > 0 {
		return nil, exception.NewErrorResponse(fiber.StatusBadRequest, "penguji 2 is not available")
	}

	// melakukan pencarian banyak schedule berdasarkan pengajuan_id -> mendapatkan data banyak pengajuan
	membersSchedule, err := service.ScheduleRepository.FindScheduleByPengajuan(pengajuansId)

	// update  satu per satu data schedule dari masing masing pengajuan_id
	var updateSchedule []entity.Schedule
	for _, oldSchedule := range membersSchedule {
		oldSchedule.DateTime = request.DateTime
		oldSchedule.Room = request.Room
		oldSchedule.Penguji1Id = request.Penguji1.Id
		oldSchedule.Penguji2Id = request.Penguji2.Id

		updateSchedule = append(updateSchedule, oldSchedule)
	}

	updatedData, err := service.ScheduleRepository.Update(updateSchedule)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, err.Error())
	}
	helper.CreateNotification(request.Members, "Penjadwalan Sidang", "Jadwal sidang anda ada perubahan!", "/schedule/get/mahasiswa", token)

	var updatedScheduleResponse []web.ScheduleResponse
	for _, newSavedSchedule := range updatedData {
		updatedScheduleResponse = append(updatedScheduleResponse, helper.ToScheduleResponse(&newSavedSchedule))
	}
	return updatedScheduleResponse, nil

}

func (service ScheduleServiceImpl) Delete(request *web.ScheduleDeleteRequest, token string) error {
	_, err := service.ScheduleRepository.FindById(request.ScheduleId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	var pengajuansId []int
	wg := &sync.WaitGroup{}
	pengajuanChannel := make(chan *web.GetPengajuanDatas, len(request.Members))

	wg.Add(1)
	go helper.ChangePengajuanStatus(pengajuanChannel, "-", "belum dijadwalkan", "penjadwalan", request.Members, token, wg)

	go func() {
		wg.Wait()
		close(pengajuanChannel)
	}()

	counter := 0
	for {
		select {
		case pengajuan := <-pengajuanChannel:
			for _, data := range pengajuan.Data {
				pengajuansId = append(pengajuansId, data.PengajuanId)
				counter++
			}
		}

		if counter == len(request.Members) {
			break
		}
	}

	_, err = service.ScheduleRepository.FindScheduleByPengajuan(pengajuansId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	err = service.ScheduleRepository.Delete(pengajuansId)
	if err != nil {
		return exception.NewErrorResponse(fiber.StatusNotFound, err.Error())
	}

	return nil
}

func (service ScheduleServiceImpl) CheckRoom(dateTime time.Time, roomName string) ([]web.ScheduleResponse, error) {
	pengajuanId := []int{10, 11, 12, 13}
	room, err := service.ScheduleRepository.CheckAvailRoom(dateTime, roomName, pengajuanId)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusNotFound, "data is not found")
	}

	var scheduleResponse []web.ScheduleResponse
	for _, schedule := range room {
		scheduleResponse = append(scheduleResponse, helper.ToScheduleResponse(&schedule))
	}
	return scheduleResponse, nil
}

func (service ScheduleServiceImpl) GetAllPagination(page int) ([]web.ScheduleResponse, error) {
	schedules, err := service.ScheduleRepository.FindAllPagination(page)
	if err != nil {
		return nil, exception.NewErrorResponse(fiber.StatusInternalServerError, "INTERNAL SERVER ERROR")
	}

	var allSchedulesResponse []web.ScheduleResponse
	for _, schedule := range schedules {
		allSchedulesResponse = append(allSchedulesResponse, helper.ToScheduleResponse(&schedule))
	}

	return allSchedulesResponse, nil
}
