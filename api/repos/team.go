package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type TeamRepo struct {
	db *gorm.DB
}

func (r *TeamRepo) GetTeam(id uint) (models.Team, error) {
	return models.Team{}, nil
}
func (r *TeamRepo) GetTeams() ([]models.Team, error) {
	return nil, nil
}
func (r *TeamRepo) PostTeam() (models.Team, error) {
	return models.Team{}, nil
}
func (r *TeamRepo) AddToTeam(TeamID, userID uint) (models.Team, error) {
	return models.Team{}, nil
}
func (r *TeamRepo) RemoveFromTeam(TeamID, userID uint) (models.Team, error) {
	return models.Team{}, nil
}
func (r *TeamRepo) PutTeam() (models.Team, error) {
	return models.Team{}, nil
}
func (r *TeamRepo) DelTeam() (bool, error) {
	return true, nil
}
