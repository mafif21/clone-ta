package services

import (
	"errors"
	"pendaftaran-sidang/internal/config"
	"pendaftaran-sidang/internal/helper"
	"pendaftaran-sidang/internal/model/entity"
	"pendaftaran-sidang/internal/model/web"
	"pendaftaran-sidang/internal/repositories"
	"strings"
)

type TeamServiceImpl struct {
	TeamRepository        repositories.TeamRepository
	PengajuanRepository   repositories.PengajuanRepository
	DocumentLogRepository repositories.DocumentLogRepository
}

func NewTeamService(teamRepository repositories.TeamRepository, pengajuanRepository repositories.PengajuanRepository, documentLogRepository repositories.DocumentLogRepository) TeamService {
	return &TeamServiceImpl{TeamRepository: teamRepository, PengajuanRepository: pengajuanRepository, DocumentLogRepository: documentLogRepository}
}

func (service TeamServiceImpl) Create(request *web.TeamCreateRequest) (*web.TeamResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundTeam, _ := service.TeamRepository.FindTeamByName(db, request.Name)
	if foundTeam != nil {
		return nil, errors.New("name team is already exists")
	}

	userFound, err := service.PengajuanRepository.FindByUserId(db, request.UserId)
	if err != nil {
		return nil, err
	}

	_, err = service.DocumentLogRepository.FindLatestDocument(db, userFound.ID, "slide")
	if err != nil {
		return nil, errors.New("make sure student already upload slide")
	}

	if strings.Contains(userFound.Status, "ditolak oleh admin") || strings.Contains(userFound.Status, "belum disetujui admin") {
		_, errorResponse, err := helper.UpdateUserTeamID(request.UserId, 0)
		if errorResponse != nil {
			return nil, errorResponse
		}

		if err != nil {
			return nil, err
		}

		return nil, errors.New("created team is invalid")
	}

	if strings.Contains(userFound.Status, "tidak lulus (sudah update dokumen)") {
		_, errorResponse, err := helper.UpdateUserTeamID(request.UserId, 0)
		if errorResponse != nil {
			return nil, errorResponse
		}

		if err != nil {
			return nil, err
		}

		_, errChangeStatus := service.PengajuanRepository.UpdateStatus(db, userFound.ID, "tidak lulus (belum dijadwalkan)")
		if errChangeStatus != nil {
			return nil, errChangeStatus
		}

		_, err = helper.AddStatusLog(db, "-", request.UserId, "penjadwalan", "tidak lulus (belum dijadwalkan)", userFound.ID)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = service.PengajuanRepository.UpdateStatus(db, userFound.ID, "belum dijadwalkan")
		if err != nil {
			return nil, err
		}
		_, err = helper.AddStatusLog(db, "-", request.UserId, "penjadwalan", "belum dijadwalkan", userFound.ID)
		if err != nil {
			return nil, err
		}
	}

	newTeam := &entity.Team{
		Name: request.Name,
	}

	saveNewTeam, err := service.TeamRepository.Save(db, newTeam)
	if err != nil {
		return nil, err
	}

	_, errorResponse, err := helper.UpdateUserTeamID(request.UserId, saveNewTeam.ID)
	if errorResponse != nil {
		return nil, errorResponse
	}
	if err != nil {
		return nil, err
	}

	response := helper.ToTeamResponse(saveNewTeam)
	return &response, nil
}

func (service TeamServiceImpl) AddMember(request *web.AddMemberRequest) error {
	db, err := config.OpenConnection()
	if err != nil {
		return errors.New("cant connect to database")
	}

	_, err = service.TeamRepository.FindTeamById(db, request.TeamId)
	if err != nil {
		return err
	}

	userFound, err := service.PengajuanRepository.FindByUserId(db, request.UserId)
	if err != nil {
		return err
	}

	_, err = service.DocumentLogRepository.FindLatestDocument(db, userFound.ID, "slide")
	if err != nil {
		return errors.New("make sure student already upload slide")
	}

	if strings.Contains(userFound.Status, "ditolak oleh admin") || strings.Contains(userFound.Status, "belum disetujui admin") {
		_, errorResponse, err := helper.UpdateUserTeamID(request.UserId, 0)
		if errorResponse != nil {
			return errorResponse
		}

		if err != nil {
			return err
		}
		return errors.New("member status is not valid to join team")
	}

	_, errorResponse, err := helper.UpdateUserTeamID(request.UserId, request.TeamId)
	if errorResponse != nil {
		return errorResponse
	}

	if err != nil {
		return err
	}

	if userFound.Status == "tidak lulus (sudah update dokumen)" {
		_, errChangeStatus := service.PengajuanRepository.UpdateStatus(db, userFound.ID, "tidak lulus (belum dijadwalkan)")
		if errChangeStatus != nil {
			return errChangeStatus
		}

		_, err = helper.AddStatusLog(db, "-", request.UserId, "penjadwalan", "tidak lulus (belum dijadwalkan)", userFound.ID)
		if err != nil {
			return err
		}
	} else {
		_, errChangeStatus := service.PengajuanRepository.UpdateStatus(db, userFound.ID, "belum dijadwalkan")
		if errChangeStatus != nil {
			return errChangeStatus
		}
		_, err = helper.AddStatusLog(db, "-", request.UserId, "penjadwalan", "belum dijadwalkan", userFound.ID)
		if err != nil {
			return err
		}
	}

	_, err = helper.AddNotification(db, request.UserId, "Invite Team", "Sukses diundang ke dalam team sidang", "/team/get-team")

	return nil
}

func (service TeamServiceImpl) LeaveTeam(request *web.LeaveMemberRequest) error {
	db, err := config.OpenConnection()
	if err != nil {
		return errors.New("cant connect to database")
	}

	teamFound, err := service.TeamRepository.FindTeamById(db, request.TeamId)
	if err != nil {
		return err
	}

	foundPengajuan, err := service.PengajuanRepository.FindByUserId(db, request.UserId)
	if err != nil {
		return err
	}

	if foundPengajuan.Status == "tidak lulus (belum dijadwalkan)" {
		_, err = service.PengajuanRepository.UpdateStatus(db, foundPengajuan.ID, "tidak lulus (sudah update dokumen)")
		if err != nil {
			return err
		}
		_, err = helper.AddStatusLog(db, "-", request.UserId, "pengajuan", "tidak lulus (sudah update dokumen)", foundPengajuan.ID)
		if err != nil {
			return err
		}
	} else {
		_, err = service.PengajuanRepository.UpdateStatus(db, foundPengajuan.ID, "telah disetujui admin")
		if err != nil {
			return err
		}
		_, err = helper.AddStatusLog(db, "-", request.UserId, "pengajuan", "telah disetujui admin", foundPengajuan.ID)
		if err != nil {
			return err
		}
	}

	_, errorResponse, err := helper.UpdateUserTeamID(request.UserId, 0)
	if errorResponse != nil {
		return errorResponse
	}
	if err != nil {
		return err
	}

	teamMember, errorResponse, err := helper.GetAllTeamMember(teamFound.ID)
	if errorResponse != nil {
		return errorResponse
	}
	if err != nil {
		return err
	}

	if len(teamMember.Data) == 0 {
		err = service.TeamRepository.Delete(db, teamFound.ID)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (service TeamServiceImpl) Update(request *web.TeamUpdateRequest) (*web.TeamResponse, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	foundDataById, err := service.TeamRepository.FindTeamById(db, request.Id)
	if err != nil {
		return nil, errors.New("data not found")
	}

	_, err = service.TeamRepository.FindTeamUpdateName(db, request.Name, foundDataById.ID)
	if err == nil {
		return nil, errors.New("team name already exists")
	}

	foundDataById.Name = request.Name
	updateData, err := service.TeamRepository.Update(db, foundDataById)
	if err != nil {
		return nil, err
	}

	response := helper.ToTeamResponse(updateData)
	return &response, nil
}

func (service TeamServiceImpl) Delete(teamId int, userId int) error {
	db, err := config.OpenConnection()
	if err != nil {
		return errors.New("cant connect to database")
	}

	_, err = service.TeamRepository.FindTeamById(db, teamId)
	if err != nil {
		return err
	}

	teamMember, errorResponse, err := helper.ResetTeam(teamId)
	if errorResponse != nil {
		return errorResponse
	}
	if err != nil {
		return err
	}

	err = service.TeamRepository.Delete(db, teamId)
	if err != nil {
		return err
	}

	for _, userId := range teamMember.Data {
		foundMember, err := service.PengajuanRepository.FindByUserId(db, userId)
		if err != nil {
			return err
		}

		if foundMember.Status == "tidak lulus (belum dijadwalkan)" {
			_, err = service.PengajuanRepository.UpdateStatus(db, foundMember.ID, "tidak lulus (sudah update dokumen)")
			if err != nil {
				return err
			}

			_, err = helper.AddStatusLog(db, "-", userId, "pengajuan", "tidak lulus (sudah update dokumen) ", foundMember.ID)
			if err != nil {
				return err
			}
		} else {
			_, err = service.PengajuanRepository.UpdateStatus(db, foundMember.ID, "telah disetujui admin")
			if err != nil {
				return err
			}

			_, err = helper.AddStatusLog(db, "-", userId, "pengajuan", "telah disetujui admin", foundMember.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (service TeamServiceImpl) GetTeamByUserId(userId int) (*web.TeamResponseDetail, error) {
	db, err := config.OpenConnection()
	if err != nil {
		return nil, errors.New("cant connect to database")
	}

	student, errorResponse, err := helper.GetDetailStudent(userId)
	if errorResponse != nil {
		return nil, errorResponse
	}
	if err != nil {
		return nil, err
	}

	foundDataById, err := service.TeamRepository.FindTeamById(db, student.Data.TeamId)
	if err != nil {
		return nil, errors.New("data not found")
	}

	teamMember, errorResponse, err := helper.GetAllTeamMember(foundDataById.ID)
	if errorResponse != nil {
		return nil, errorResponse
	}

	if err != nil {
		return nil, err
	}

	response := helper.ToTeamResponseDetail(foundDataById, teamMember)
	return &response, nil
}
