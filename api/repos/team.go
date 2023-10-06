package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type TeamRepo struct {
	db *gorm.DB
}

func (r *TeamRepo) GetTeam(id uint) (*models.Team, error) {
	var team models.Team
	if err := r.db.Debug().First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) GetTeams() ([]models.Team, error) {
	var team []models.Team
	if err := r.db.Debug().Find(&team).Error; err != nil {
		return nil, err
	}
	return team, nil
}
func (r *TeamRepo) PostTeam(team models.Team) (*models.Team, error) {
	if err := r.db.Debug().Create(&team).Error; err != nil {
		return nil, err
	}
	return nil, nil
}
func (r *TeamRepo) PutTeam(team models.Team) (*models.Team, error) {
	if err := r.db.Debug().Updates(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) DelTeam(id uint) (bool, error) {
	var team models.Team
	if err := r.db.Debug().Delete(&team, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *TeamRepo) AddToTeam(TeamID, userID uint) (models.Team, error) {
	return models.Team{}, nil
}
func (r *TeamRepo) RemoveFromTeam(TeamID, userID uint) (models.Team, error) {
	return models.Team{}, nil
}
