package repositories

import (
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type TeamRepository interface {
	FindTeamById(db *gorm.DB, teamId int) (*entity.Team, error)
	FindTeamByName(db *gorm.DB, teamName string) (*entity.Team, error)
	FindTeamUpdateName(db *gorm.DB, teamName string, teamId int) (*entity.Team, error)
	Save(db *gorm.DB, team *entity.Team) (*entity.Team, error)
	Update(db *gorm.DB, team *entity.Team) (*entity.Team, error)
	Delete(db *gorm.DB, teamId int) error
}
