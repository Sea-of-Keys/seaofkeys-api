package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/middleware"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type AuthRepo struct {
	db *gorm.DB
}

func (r *AuthRepo) PostLogin(user models.User) (models.User, error) {
	var checkUser models.User
	if err := r.db.Debug().Where("email = ?", user.Email).Find(&checkUser).Error; err != nil {
		return models.User{}, nil
	}
	if !middleware.CheckPasswordHash(user.Password, checkUser.Password) {
		return models.User{}, nil
	}

	checkUser.Password = ""
	checkUser.Code = ""

	return checkUser, nil
}
func (r *AuthRepo) PostLogout() (bool, error) {
	return true, nil
}
