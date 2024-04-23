package services

import (
	"errors"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
	"time"
)

type PeriodServiceImpl struct {
	Repository repositories.PeriodRepository
}

func NewPeriodService(repository repositories.PeriodRepository) PeriodService {
	return &PeriodServiceImpl{
		Repository: repository,
	}
}

func (service *PeriodServiceImpl) GetAllPeriod() ([]web.PeriodResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var allPeriodResponse []web.PeriodResponse

	periodDatas, err := service.Repository.FindAll(db)
	if err != nil {
		return nil, err
	}

	for _, period := range periodDatas {
		allPeriodResponse = append(allPeriodResponse, helper.ToPeriodResponse(&period))
	}

	return allPeriodResponse, nil
}

func (service *PeriodServiceImpl) GetPeriodById(periodId int) (*web.PeriodResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundPeriod, err := service.Repository.FindPeriodById(db, periodId)
	if err != nil {
		return nil, err
	}

	response := helper.ToPeriodResponse(foundPeriod)
	return &response, nil
}

func (service *PeriodServiceImpl) Create(request *web.PeriodCreateRequest) (*web.PeriodResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, err
	}

	_, err = service.Repository.FindPeriodByName(db, request.Name)
	if err == nil {
		return nil, errors.New("period name already exists")
	}

	if request.StartDate.After(request.EndDate) || request.StartDate.Equal(request.EndDate) {
		return nil, errors.New("period date is not valid")
	}

	newPeriod := &entity.Period{
		Name:        request.Name,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Description: request.Description,
	}

	save, err := service.Repository.Save(db, newPeriod)
	if err != nil {
		return nil, err
	}

	response := helper.ToPeriodResponse(save)
	return &response, nil

}

func (service *PeriodServiceImpl) Update(request *web.PeriodUpdateRequest) (*web.PeriodResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	if request.StartDate.After(request.EndDate) || request.StartDate.Equal(request.EndDate) {
		return nil, errors.New("period date is not valid")
	}

	foundPeriod, err := service.Repository.FindPeriodById(db, request.Id)
	if err != nil {
		return nil, err
	}

	_, err = service.Repository.FindPeriodUpdateName(db, request.Name, request.Id)
	if err == nil {
		return nil, errors.New("found same name")
	}

	foundPeriod.Name = request.Name
	foundPeriod.StartDate = request.StartDate
	foundPeriod.EndDate = request.EndDate
	foundPeriod.Description = request.Description

	update, err := service.Repository.Update(db, foundPeriod)
	if err != nil {
		return nil, err
	}

	response := helper.ToPeriodResponse(update)
	return &response, nil

}

func (service *PeriodServiceImpl) Delete(periodId int) error {
	db, err := config.OpenConnection()
	if err != nil {
		return errors.New("cant connect to database")
	}

	foundPeriod, err := service.Repository.FindPeriodById(db, periodId)
	if err != nil {
		return err
	}

	err = service.Repository.Delete(db, foundPeriod.ID)
	if err != nil {
		return err
	}

	return nil
}

func (service *PeriodServiceImpl) GetPeriodByTime(time time.Time) (*web.PeriodResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundPeriod, err := service.Repository.FindPeriodByTime(db, time)
	if err != nil {
		return nil, err
	}

	response := helper.ToPeriodResponse(foundPeriod)
	return &response, nil
}
