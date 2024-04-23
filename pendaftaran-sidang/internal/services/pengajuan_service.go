package services

import "pendaftaran-sidang/internal/model/web"

type PengajuanService interface {
	GetAllPengajuan(isAdmin bool, userId int) ([]web.PengajuanResponse, error)
	GetAllPengajuanPembimbing(userId int) ([]web.PengajuanResponse, error)
	GetPengajuanById(pengajuanId int, userId int, isAdmin bool) (*web.PengajuanResponse, error)
	GetPengajuanByUserId(userId int) (*web.PengajuanResponse, error)
	GetPengajuanByUsers(request *web.GetPengajuanByUsersId) ([]int, error)
	Create(request *web.PengajuanCreateRequest, username string) (*web.PengajuanResponse, error)
	Update(request *web.PengajuanUpdateRequest, username string) (*web.PengajuanResponse, error)
	AdminStatus(request *web.StatusAdminUpdate) (*web.PengajuanResponse, error)
	ChangeStatus(request *web.ChangeStatusRequest, createdBy int) ([]web.PengajuanResponse, error)
	GetAllApprovePengajuan() ([]web.PengajuanResponse, error)
	GetAllPengajuanByPeriod(periodId int) ([]web.PengajuanResponse, error)
}
