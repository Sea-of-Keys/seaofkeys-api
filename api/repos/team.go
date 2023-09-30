package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type TeamRepo struct {
	db *gorm.DB
}

func (repo *TeamRepo) GetTeam(id uint) (models.Team, error) {
	return models.Team{}, nil
}
func (repo *TeamRepo) GetTeams() ([]models.Team, error) {
	return nil, nil
}
func (repo *TeamRepo) PostTeam() (models.Team, error) {
	return models.Team{}, nil
}
func (repo *TeamRepo) AddToTeam(TeamID, userID uint) (models.Team, error) {
	return models.Team{}, nil
}
func (repo *TeamRepo) RemoveFromTeam(TeamID, userID uint) (models.Team, error) {
	return models.Team{}, nil
}
func (repo *TeamRepo) PutTeam() (models.Team, error) {
	return models.Team{}, nil
}
func (repo *TeamRepo) DelTeam() (bool, error) {
	return true, nil
}
