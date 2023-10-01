package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type UserRepo struct {
	db *gorm.DB
}

func (repo *UserRepo) GetUser(id uint) (models.User, error) {
	return models.User{}, nil
}
