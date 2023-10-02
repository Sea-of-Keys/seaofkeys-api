package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type PermissionRepo struct {
	db *gorm.DB
}

func (r *PermissionRepo) GetPermission(id uint) (models.Permission, error) {
	return models.Permission{}, nil
}
func (r *PermissionRepo) GetPermissions() ([]models.Permission, error) {
	return nil, nil
}
func (r *PermissionRepo) PostPermission(per models.Permission) (models.Permission, error) {
	return models.Permission{}, nil
}
func (r *PermissionRepo) PutPermission(per models.Permission) (models.Permission, error) {
	return models.Permission{}, nil
}
func (r *PermissionRepo) DelPermission(id uint) (bool, error) {
	return true, nil
}
func (r *PermissionRepo) CleanPermission() error {
	return nil
}
