package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type AuthRepo struct {
	db *gorm.DB
}

func (repo *AuthRepo) PostLogin(user models.User) (models.User, error) {
	return models.User{}, nil
}
func (repo *AuthRepo) PostLogout() (bool, error) {
	return true, nil
}
