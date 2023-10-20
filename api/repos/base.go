package repos

import "gorm.io/gorm"

type Repository interface {
	// Define common repository methods here, such as CRUD operations.
	// For example:
	FindByID(id uint) interface{}
	Create(model interface{}) error
	Update(model interface{}) error
	Delete(id uint) error
}

type BaseRepo struct {
	db *gorm.DB
}

func NewBaseRepo(db *gorm.DB) *BaseRepo {
	return &BaseRepo{db}
}
