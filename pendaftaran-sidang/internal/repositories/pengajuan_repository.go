package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type PengajuanRepository interface {
	FindAll(db *gorm.DB, isAdmin bool, userId int) ([]entity.Pengajuan, error)
	FindByPengajuanId(db *gorm.DB, pengajuanId int) (*entity.Pengajuan, error)
	FindPembimbingById(db *gorm.DB, userId int) ([]entity.Pengajuan, error)
	CheckPembimbing(db *gorm.DB, userId int, pengajuanId int) (*entity.Pengajuan, error)
	FindByUserId(db *gorm.DB, userId int) (*entity.Pengajuan, error)
	FindByUsers(db *gorm.DB, users []int) ([]int, error)
	Save(db *gorm.DB, pengajuan *entity.Pengajuan) (*entity.Pengajuan, error)
	Update(db *gorm.DB, pengajuan *entity.Pengajuan) (*entity.Pengajuan, error)
	UpdateStatus(db *gorm.DB, pengajuanId int, value string) (*entity.Pengajuan, error)
	UpdateAdminStatus(db *gorm.DB, pengajuanId int, value string, isEnglish bool) (*entity.Pengajuan, error)
	UpdateManyStatus(db *gorm.DB, message string, pengajuanId []int) ([]entity.Pengajuan, error)
	FindAllApprovePengajuan(db *gorm.DB) ([]entity.Pengajuan, error)
	FindPengajuanByPeriod(db *gorm.DB, periodId int) ([]entity.Pengajuan, error)
}
