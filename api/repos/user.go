package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type UserRepo struct {
	db *gorm.DB
}

func (r *UserRepo) GetUser(id uint) (models.User, error) {
	return models.User{}, nil
}
func (r *UserRepo) GetUsers() ([]models.User, error) {
	return nil, nil
}
func (r *UserRepo) PostUser(user models.User) (models.User, error) {
	return models.User{}, nil
}
func (r *UserRepo) PostUsers(user []models.User) ([]models.User, error) {
	return nil, nil
}
func (r *UserRepo) PutUser(user models.User) (models.User, error) {
	return models.User{}, nil
}
func (r *UserRepo) DelUser(id uint) (bool, error) {
	return true, nil
}
