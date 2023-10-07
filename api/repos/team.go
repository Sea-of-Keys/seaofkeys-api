package repos

import (
	"errors"

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
func (r *TeamRepo) AddToTeam(TeamID, userID uint) (*models.Team, error) {
	var team models.Team
	var user models.User
	team.ID = TeamID
	if err := r.db.Debug().Preload("Users").First(&team, TeamID).Error; err != nil {
		return nil, err
	}
	if err := r.db.Debug().First(&user, userID).Error; err != nil {
		return nil, err
	}
	for _, v := range team.Users {
		if v.ID == userID {
			return nil, errors.New("R27: User all ready on the team")
		}
	}
	team.Users = append(team.Users, &user)
	if err := r.db.Debug().Save(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) RemoveFromTeam(TeamID, userID uint) (*models.Team, error) {
	var team models.Team
	var user models.User
	// var users []models.User

	if err := r.db.Preload("Users").First(&team, TeamID).Error; err != nil {
		return nil, err
	}
	user.ID = userID
	if err := r.db.Debug().Model(&team).Association("Users").Delete(&user); err != nil {
		return nil, err
	}
	return &team, nil
}
func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{db}
}
