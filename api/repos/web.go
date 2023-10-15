package repos

import (
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

func NewWebRepo(db *gorm.DB) *WebRepo {
	return &WebRepo{db}
}
