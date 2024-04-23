package services

import (
	"errors"
	"fmt"
	"os"
	"path"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
	"strconv"
)

type PengajuanServiceImpl struct {
	PengajuanRepository repositories.PengajuanRepository
	PeriodRepository    repositories.PeriodRepository
}

func NewPengajuanService(pengajuanRepository repositories.PengajuanRepository, periodRepository repositories.PeriodRepository) PengajuanService {
	return &PengajuanServiceImpl{
		PengajuanRepository: pengajuanRepository,
		PeriodRepository:    periodRepository,
	}
}

func (service *PengajuanServiceImpl) GetAllPengajuan(isAdmin bool, userId int) ([]web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var allDatasResponse []web.PengajuanResponse

	all, err := service.PengajuanRepository.FindAll(db, isAdmin, userId)
	if err != nil {
		return nil, err
	}
	for _, pengajuan := range all {
		allDatasResponse = append(allDatasResponse, helper.ToPengajuanResponse(&pengajuan))
	}
	return allDatasResponse, err
}

func (service *PengajuanServiceImpl) GetAllPengajuanPembimbing(userId int) ([]web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	all, err := service.PengajuanRepository.FindPembimbingById(db, userId)
	if err != nil {
		return nil, err
	}

	var allDatasResponse []web.PengajuanResponse
	for _, pengajuan := range all {
		allDatasResponse = append(allDatasResponse, helper.ToPengajuanResponse(&pengajuan))
	}

	return allDatasResponse, err

}

func (service *PengajuanServiceImpl) GetPengajuanById(pengajuanId int, userId int, isAdmin bool) (*web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundPengajuan, err := service.PengajuanRepository.FindByPengajuanId(db, pengajuanId)
	if err != nil {
		return nil, err
	}

	if isAdmin {
		response := helper.ToPengajuanResponse(foundPengajuan)
		return &response, nil
	}

	if foundPengajuan.Pembimbing1Id != userId && foundPengajuan.Pembimbing2Id != userId {
		return nil, errors.New("you can't acces this data")
	}

	response := helper.ToPengajuanResponse(foundPengajuan)
	return &response, nil
}

func (service *PengajuanServiceImpl) GetPengajuanByUserId(userId int) (*web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundUserPengajuan, err := service.PengajuanRepository.FindByUserId(db, userId)
	if err != nil {
		return nil, err
	}

	response := helper.ToPengajuanResponse(foundUserPengajuan)
	return &response, nil
}

func (service *PengajuanServiceImpl) GetPengajuanByUsers(request *web.GetPengajuanByUsersId) ([]int, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	pengajuanId, err := service.PengajuanRepository.FindByUsers(db, request.UsersId)
	if err != nil {
		return nil, err
	}

	return pengajuanId, err
}

func (service *PengajuanServiceImpl) Create(request *web.PengajuanCreateRequest, username string) (*web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	if request.Pembimbing1Id == request.Pembimbing2Id {
		return nil, errors.New("cant same pembimbing 1 and pembimbing 2")
	}

	foundData, _ := service.PengajuanRepository.FindByUserId(db, request.UserId)
	if foundData != nil {
		return nil, errors.New("user already registered in database")
	}

	newPengajuan := entity.Pengajuan{
		UserId:         request.UserId,
		Nim:            request.Nim,
		Pembimbing1Id:  request.Pembimbing1Id,
		Pembimbing2Id:  request.Pembimbing2Id,
		Judul:          request.Judul,
		Eprt:           request.Eprt,
		Tak:            request.Tak,
		PeriodID:       request.PeriodID,
		FormBimbingan1: request.FormBimbingan1,
		FormBimbingan2: request.FormBimbingan2,
	}

	docTaUrl := fmt.Sprintf("/public/doc_ta/%s", request.DocTa)
	newPengajuan.DocTa = docTaUrl

	makalahUrl := fmt.Sprintf("/public/makalah/%s", request.Makalah)
	newPengajuan.Makalah = makalahUrl

	save, err := service.PengajuanRepository.Save(db, &newPengajuan)
	if err != nil {
		return nil, err
	}

	_, errorResponse, err := helper.UpdatePeminatanByUserId(request.Peminatan, request.UserId)
	if errorResponse != nil {
		return nil, errorResponse
	}

	if err != nil {
		return nil, err
	}

	_, _ = helper.AddDocumentLog(db, save.ID, save.Makalah, "makalah", makalahUrl, request.UserId)
	if err != nil {
		return nil, err
	}

	_, err = helper.AddDocumentLog(db, save.ID, save.DocTa, "draft", docTaUrl, request.UserId)
	if err != nil {
		return nil, err
	}

	_, err = helper.AddStatusLog(db, "-", save.UserId, "pengajuan", "pengajuan", save.ID)
	if err != nil {
		return nil, err
	}

	_, err = helper.AddNotification(db, request.Pembimbing1Id, "Mahasiswa daftar sidang", "Mahasiswa dengan username "+username+" telah mendaftarkan sidang", "/pengajuan/get")
	_, err = helper.AddNotification(db, request.Pembimbing2Id, "Mahasiswa daftar sidang", "Mahasiswa dengan username "+username+" telah mendaftarkan sidang", "/pengajuan/get")

	response := helper.ToPengajuanResponse(save)
	return &response, nil
}

func (service *PengajuanServiceImpl) Update(request *web.PengajuanUpdateRequest, username string) (*web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	oldData, err := service.PengajuanRepository.FindByPengajuanId(db, request.Id)
	if err != nil {
		return nil, err
	}

	if request.Pembimbing1Id == request.Pembimbing2Id {
		return nil, errors.New("cant same pembimbing 1 and pembimbing 2")
	}

	studentFound, err := service.PengajuanRepository.FindByUserId(db, request.UserId)
	if err != nil {
		return nil, errors.New("user not registered")
	}

	// delete old file
	if request.DocTa != "" {
		_ = os.Remove("./public/doc_ta/" + path.Base(studentFound.DocTa))
		studentFound.DocTa = fmt.Sprintf("/public/doc_ta/%s", request.DocTa)
	} else {
		studentFound.DocTa = oldData.DocTa
	}

	if request.Makalah != "" {
		_ = os.Remove("./public/makalah/" + path.Base(studentFound.Makalah))
		studentFound.Makalah = fmt.Sprintf("/public/makalah/%s", request.Makalah)
	} else {
		studentFound.Makalah = oldData.Makalah
	}

	studentFound.Nim = request.Nim
	studentFound.Pembimbing1Id = request.Pembimbing1Id
	studentFound.Pembimbing2Id = request.Pembimbing2Id
	studentFound.Judul = request.Judul
	studentFound.Eprt = request.Eprt
	studentFound.Tak = request.Tak

	studentFound.PeriodID = request.PeriodID
	studentFound.FormBimbingan1 = request.FormBimbingan1
	studentFound.FormBimbingan2 = request.FormBimbingan2

	update, err := service.PengajuanRepository.Update(db, studentFound)
	if err != nil {
		return nil, err
	}

	_, errorResponse, err := helper.UpdatePeminatanByUserId(request.Peminatan, request.UserId)
	if errorResponse != nil {
		return nil, errorResponse
	}

	if err != nil {
		return nil, err
	}

	_, err = helper.AddStatusLog(db, "-", update.UserId, "pengajuan", "perbaikan berkas", update.ID)
	if err != nil {
		return nil, err
	}

	_, err = helper.AddNotification(db, request.Pembimbing1Id, "Mahasiswa Edit Data Pengajuan Sidang", "Mahasiswa dengan username "+username+" melakukan perubahan data daftar sidang", "/pengajuan/get")
	_, err = helper.AddNotification(db, request.Pembimbing2Id, "Mahasiswa Edit Data Pengajuan Sidang", "Mahasiswa dengan username "+username+" melakukan perubahan data daftar sidang", "/pengajuan/get")

	response := helper.ToPengajuanResponse(update)
	return &response, nil
}

func (service *PengajuanServiceImpl) AdminStatus(request *web.StatusAdminUpdate) (*web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundPengajuan, err := service.PengajuanRepository.FindByPengajuanId(db, request.Id)
	if err != nil {
		return nil, err
	}

	var message string
	var title string
	var messageNotification string
	var url string

	if request.Value == "accept" {
		message = "telah disetujui admin"
		title = "Approved"
		messageNotification = "Sidang anda telah di approve admin, silahkan membuat team"
		url = "team/get-team"
	} else {
		message = "ditolak oleh admin"
		title = "Rejected"
		messageNotification = "Berkas Anda Ditolak Admin, Silahkan Perbaiki Berkas Anda"
		url = "pengajuan/update/" + strconv.Itoa(foundPengajuan.ID)
	}

	if request.Feedback == "" {
		request.Feedback = "-"
	}

	updatedData, err := service.PengajuanRepository.UpdateAdminStatus(db, foundPengajuan.ID, message, request.IsEnglish)
	if err != nil {
		return nil, err
	}

	_, err = helper.AddStatusLog(db, request.Feedback, request.UserId, "pengajuan", message, updatedData.ID)
	if err != nil {
		return nil, err
	}

	_, err = helper.AddNotification(db, foundPengajuan.UserId, title, messageNotification, url)

	response := helper.ToPengajuanResponse(updatedData)
	return &response, nil
}

func (service *PengajuanServiceImpl) ChangeStatus(request *web.ChangeStatusRequest, createdBy int) ([]web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var pengajuanResponse []web.PengajuanResponse
	for _, memberId := range request.MemberId {
		pengajuanData, err := service.PengajuanRepository.FindByUserId(db, memberId)
		if err != nil {
			return nil, err
		}

		if request.Status == "sudah dijadwalkan" {
			if pengajuanData.Status == "tidak lulus (belum dijadwalkan)" {
				_, err := service.PengajuanRepository.UpdateStatus(db, pengajuanData.ID, "tidak lulus (sudah dijadwalkan)")
				if err != nil {
					return nil, err
				}

				_, err = helper.AddStatusLog(db, request.Feedback, createdBy, request.WorkFlowType, "tidak lulus(sudah dijadwalkan)", pengajuanData.ID)
				if err != nil {
					return nil, err
				}
			} else {
				_, err := service.PengajuanRepository.UpdateStatus(db, pengajuanData.ID, request.Status)
				if err != nil {
					return nil, err
				}

				_, err = helper.AddStatusLog(db, request.Feedback, createdBy, request.WorkFlowType, request.Status, pengajuanData.ID)
				if err != nil {
					return nil, err
				}
			}
		} else if request.Status == "belum dijadwalkan" {
			if pengajuanData.Status == "tidak lulus (sudah dijadwalkan)" {
				_, err := service.PengajuanRepository.UpdateStatus(db, pengajuanData.ID, "tidak lulus (belum dijadwalkan)")
				if err != nil {
					return nil, err
				}

				_, err = helper.AddStatusLog(db, request.Feedback, createdBy, request.WorkFlowType, "tidak lulus (belum dijadwalkan)", pengajuanData.ID)
				if err != nil {
					return nil, err
				}
			} else {
				_, err := service.PengajuanRepository.UpdateStatus(db, pengajuanData.ID, request.Status)
				if err != nil {
					return nil, err
				}

				_, err = helper.AddStatusLog(db, request.Feedback, createdBy, request.WorkFlowType, request.Status, pengajuanData.ID)
				if err != nil {
					return nil, err
				}
			}
		}
		pengajuanResponse = append(pengajuanResponse, helper.ToPengajuanResponse(pengajuanData))
	}

	return pengajuanResponse, nil
}

func (service *PengajuanServiceImpl) GetAllApprovePengajuan() ([]web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	var allDatasResponse []web.PengajuanResponse

	all, err := service.PengajuanRepository.FindAllApprovePengajuan(db)
	if err != nil {
		return nil, err
	}
	for _, pengajuan := range all {
		allDatasResponse = append(allDatasResponse, helper.ToPengajuanResponse(&pengajuan))
	}
	return allDatasResponse, err
}

func (service *PengajuanServiceImpl) GetAllPengajuanByPeriod(periodId int) ([]web.PengajuanResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundPeriod, err := service.PeriodRepository.FindPeriodById(db, periodId)
	if err != nil {
		return nil, err
	}

	var allDatasResponse []web.PengajuanResponse

	all, err := service.PengajuanRepository.FindPengajuanByPeriod(db, foundPeriod.ID)
	if err != nil {
		return nil, err
	}

	for _, pengajuan := range all {
		allDatasResponse = append(allDatasResponse, helper.ToPengajuanResponse(&pengajuan))
	}
	return allDatasResponse, err
}
