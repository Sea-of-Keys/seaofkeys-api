package repos

import "gorm.io/gorm"

type WebRepo struct {
	db *gorm.DB
}

func NewWebRepo(db *gorm.DB) *WebRepo {
	return &WebRepo{db}
}
