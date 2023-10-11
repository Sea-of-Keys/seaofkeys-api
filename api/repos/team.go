package repos

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type TeamRepo struct {
	db *gorm.DB
}

func (r *TeamRepo) GetTeam(id uint) (*models.Team, error) {
	var team models.Team
	if err := r.db.Debug().Preload("Users").First(&team, id).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) GetTeams() ([]models.Team, error) {
	var team []models.Team
	if err := r.db.Debug().Preload("Users").Find(&team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

// ##### Maby make a find after create
func (r *TeamRepo) PostTeam(team models.Team) (*models.Team, error) {
	if err := r.db.Debug().Preload("User").Create(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) PutTeam(team models.Team) (*models.Team, error) {
	if err := r.db.Debug().Preload("Users").Updates(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) DelTeam(id uint) (bool, error) {
	var team models.Team
	team.ID = id
	r.db.Debug().Model(&team).Association("Users").Clear()
	if err := r.db.Debug().Delete(&team, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *TeamRepo) DelTeams(id []models.Delete) (bool, error) {
	var team models.Team

	for _, v := range id {
		team.ID = v.ID
		r.db.Debug().Model(&team).Association("Users").Clear()
		if err := r.db.Debug().Delete(&team, v.ID).Error; err != nil {
			return false, err
		}
	}
	return true, nil
}
func (r *TeamRepo) GetAllUserNotOnTheTeam(TeamID uint) ([]models.User, error) {
	var users []models.User
	var team models.Team
	if err := r.db.Debug().First(&team, TeamID).Error; err != nil {
		return nil, err
	}
	if err := r.db.Debug().Where("id NOT IN (SELECT user_id FROM teams_users WHERE team_id = ?)", TeamID).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (r *TeamRepo) AddToTeam(UT models.RemoveUsersFromTeam) (*models.Team, error) {
	var team models.Team
	var user []models.User
	if err := r.db.Debug().Preload("Users").First(&team, UT.TeamID).Error; err != nil {
		return nil, err
	}
	if err := r.db.Debug().Find(&user, UT.UserID).Error; err != nil {
		return nil, err
	}
	for _, v := range user {
		var yes bool
		for _, u := range v.Teams {
			if u.ID == UT.TeamID {
				yes = true
				break
			}
		}
		if !yes {
			newUser := v // Copy the user
			team.Users = append(team.Users, &newUser)
		}
	}
	for _, v := range team.Users {
		fmt.Printf("User: %v\n", v.Name)
	}
	if err := r.db.Debug().Save(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) AddUsersToTeam(TeamID, userID uint) (*models.Team, error) {
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

	if err := r.db.Preload("Users").First(&team, TeamID).Error; err != nil {
		return nil, err
	}
	user.ID = userID
	if err := r.db.Debug().Model(&team).Association("Users").Delete(&user); err != nil {
		return nil, err
	}
	return &team, nil
}
func (r *TeamRepo) RemoveUsersFromTeam(UT models.RemoveUsersFromTeam) (*models.Team, error) {
	var team models.Team
	var user models.User

	if err := r.db.Preload("Users").First(&team, UT.TeamID).Error; err != nil {
		return nil, err
	}
	fmt.Println(UT)
	for _, v := range UT.UserID {
		fmt.Println(v)
		user.ID = v
		if err := r.db.Debug().Model(&team).Association("Users").Delete(&user); err != nil {
			return nil, err
		}
	}
	return &team, nil
}
func NewTeamRepo(db *gorm.DB) *TeamRepo {
	return &TeamRepo{db}
}
