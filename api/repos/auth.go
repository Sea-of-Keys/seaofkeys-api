package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/middleware"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type AuthRepo struct {
	db *gorm.DB
}

func (repo *AuthRepo) PostLogin(user models.User) (*models.User, error) {
	var checkUser models.User
	if err := repo.db.Debug().Where("email = ?", user.Email).Find(&checkUser).Error; err != nil {
		return nil, nil
	}
	if !middleware.CheckPasswordHash(user.Password, checkUser.Password) {
		return nil, nil
	}
	checkUser.Password = ""
	checkUser.Code = ""

	return &checkUser, nil
}
func (repo *AuthRepo) PostLogout() (bool, error) {
	return true, nil
}

// Change Embedde Password ### CHANGE CODE TO int ?????
func (repo *AuthRepo) PutPassword(id uint, code string) (*models.User, error) {
	return nil, nil
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db}
}
