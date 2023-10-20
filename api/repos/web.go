package repos

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type WebRepo struct {
	db *gorm.DB
}

func (r *WebRepo) GetCheckToken(token string) (*models.UserPC, error) {
	var userPC models.UserPC
	if err := r.db.Debug().Where("Token = ?", token).First(&userPC).Error; err != nil {
		return nil, err
	}
	return &userPC, nil
}

func (r *WebRepo) PutPassword(code string, token string, password ...string) (bool, error) {
	var userPC models.UserPC
	var user models.User

	if err := r.db.Debug().Where("token = ?", token).First(&userPC).Error; err != nil {
		return false, errors.New("can't find uder with token")
	}
	if err := r.db.Debug().First(&user, userPC.ID).Error; err != nil {
		return false, errors.New("can't find uder with id")
	}
	user.Code = &code
	if err := r.db.Model(&user).Updates(&user).Error; err != nil {
		return false, errors.New("failed to update users code")
	}
	return true, nil
}
func NewWebRepo(db *gorm.DB) *WebRepo {
	return &WebRepo{db}
}
