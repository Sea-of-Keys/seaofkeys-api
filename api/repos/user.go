package repos

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Debug().Preload("Teams").First(&user, id).Error; err != nil {
		return nil, errors.New("ERROR 10: " + err.Error())
	}
	return &user, nil
}
func (r *UserRepo) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Debug().Find(&users).Error; err != nil {
		return nil, errors.New("ERROR 11: " + err.Error())
	}
	return users, nil
}
func (r *UserRepo) PostUser(user models.User) (*models.User, error) {
	var UserPC models.UserPC
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return nil, errors.New("ERROR 12: " + err.Error())
	}
	UserPC.UserID = user.ID
	// if err :=
	token, err := security.NewToken(user.ID, *user.Email)
	if err != nil {
		return nil, err
	}
	UserPC.Token = token
	if err := r.db.Debug().Create(&UserPC).Error; err != nil {
		return nil, err
	}
	fmt.Println(UserPC)
	return &user, nil
}
func (r *UserRepo) PostUsers(users []models.User) ([]models.User, error) {
	if err := r.db.Debug().Create(&users).Error; err != nil {
		return nil, errors.New("ERROR 12: " + err.Error())
	}
	return users, nil
}
func (r *UserRepo) PutUser(user models.User) (*models.User, error) {

	if err := r.db.Debug().Model(&user).Updates(&user).Error; err != nil {
		return nil, errors.New("ERROR 12: " + err.Error())
	}
	return &user, nil
}
func (r *UserRepo) DelUser(id uint) (bool, error) {
	var user models.User
	if err := r.db.Debug().Model(&user).
		Where("ID = ?", id).
		Updates(map[string]interface{}{
			"Email":    nil,
			"Code":     nil,
			"Password": nil,
		}).Error; err != nil {
		return false, errors.New("ERROR 13: " + err.Error())
	}
	if err := r.db.Debug().Delete(&user, id).Error; err != nil {
		return false, errors.New("ERROR 13: " + err.Error())
	}
	return true, nil
}
func (r *UserRepo) DelUsers(id []models.Delete) (bool, error) {
	var user models.User
	// gg := []uint{2, 3}
	fmt.Println(id)
	for _, v := range id {
		if err := r.db.Debug().Model(&user).
			Where("ID = ?", v.ID).
			Updates(map[string]interface{}{
				"Email":    nil,
				"Code":     nil,
				"Password": nil,
			}).Error; err != nil {
			return false, errors.New("ERROR 13: " + err.Error())
		}
		if err := r.db.Debug().Delete(&user, v.ID).Error; err != nil {
			return false, errors.New("ERROR 13: " + err.Error())
		}
	}
	return true, nil
}
func (r *UserRepo) GetAllTeamsUserIsNotOn(UserID uint) ([]models.Team, error) {
	var teams []models.Team
	var user models.User
	if err := r.db.Debug().First(&user, UserID).Error; err != nil {
		return nil, err
	}

	if err := r.db.Debug().Where("id NOT IN (SELECT team_id FROM teams_users WHERE user_id = ?)", UserID).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
