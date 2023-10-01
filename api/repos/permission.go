package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type PermissionRepo struct {
	db *gorm.DB
}

func (repo *PermissionRepo) GetPermission(id uint) (models.Permission, error) {
	return models.Permission{}, nil
}
func (repo *PermissionRepo) GetPermissions() ([]models.Permission, error) {
	return nil, nil
}
func (repo *PermissionRepo) PostPermission(per models.Permission) (models.Permission, error) {
	return models.Permission{}, nil
}
func (repo *PermissionRepo) PutPermission(per models.Permission) (models.Permission, error) {
	return models.Permission{}, nil
}
func (repo *PermissionRepo) DelPermission(id uint) (bool, error) {
	return true, nil
}
func (repo *PermissionRepo) CleanPermission() error {
	return nil
}
