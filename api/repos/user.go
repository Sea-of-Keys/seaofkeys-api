package repos

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.Debug().First(&user, id).Error; err != nil {
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
	if err := r.db.Debug().Create(&user).Error; err != nil {
		return nil, errors.New("ERROR 12: " + err.Error())
	}
	return &user, nil
}
func (r *UserRepo) PostUsers(users []models.User) ([]models.User, error) {
	if err := r.db.Debug().Create(&users).Error; err != nil {
		return nil, errors.New("ERROR 12: " + err.Error())
	}
	return users, nil
}
func (r *UserRepo) PutUser(user models.User) (models.User, error) {
	return models.User{}, nil
}
func (r *UserRepo) DelUser(id uint) (bool, error) {
	var user models.User
	if err := r.db.Debug().Delete(&user, id).Error; err != nil {
		return false, errors.New("ERROR 13: " + err.Error())
	}
	return true, nil
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}
