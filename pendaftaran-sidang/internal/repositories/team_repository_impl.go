package repositories

import (
	"errors"
	"gorm.io/gorm"
	"pendaftaran-sidang/internal/model/entity"
)

type TeamRepositoryImpl struct{}

func NewTeamRepository() TeamRepository {
	return &TeamRepositoryImpl{}
}

func (repository TeamRepositoryImpl) FindTeamById(db *gorm.DB, teamId int) (*entity.Team, error) {
	var team *entity.Team
	err := db.Take(&team, "id = ?", teamId).Error
	if err != nil {
		return nil, errors.New("team not found")
	}

	return team, nil
}

func (repository TeamRepositoryImpl) FindTeamByName(db *gorm.DB, teamName string) (*entity.Team, error) {
	var team *entity.Team
	err := db.Where("name = ?", teamName).Take(&team).Error
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (repository TeamRepositoryImpl) FindTeamUpdateName(db *gorm.DB, teamName string, teamId int) (*entity.Team, error) {
	var team entity.Team
	err := db.Take(&team, "name = ?", teamName).Where("id != ?", teamId).Error
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (repository TeamRepositoryImpl) Save(db *gorm.DB, team *entity.Team) (*entity.Team, error) {
	err := db.Create(&team).Error
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (repository TeamRepositoryImpl) Update(db *gorm.DB, team *entity.Team) (*entity.Team, error) {
	err := db.Model(&entity.Team{}).Where("id = ?", team.ID).Update("name", team.Name).Error
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (repository TeamRepositoryImpl) Delete(db *gorm.DB, teamId int) error {
	err := db.Delete(&entity.Team{}, "id = ?", teamId).Error
	if err != nil {
		return errors.New("data not found")
	}

	return nil
}
