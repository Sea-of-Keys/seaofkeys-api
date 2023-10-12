package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type PermissionRepo struct {
	db *gorm.DB
}

func (r *PermissionRepo) GetPermission(id uint) (*models.Permission, error) {
	var permission models.Permission
	if err := r.db.Debug().Preload("User").Preload("Team").Preload("Room").Preload("Weekdays").First(&permission, id).Error; err != nil {
		return nil, err
	}
	return &permission, nil
}
func (r *PermissionRepo) GetPermissions() ([]models.Permission, error) {
	var permission []models.Permission
	if err := r.db.Debug().Preload("User").Preload("Team").Preload("Room").Preload("Weekdays").Find(&permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}
func (r *PermissionRepo) PostPermission(per models.Permission) (*models.Permission, error) {
	if err := r.db.Debug().Preload("User").Preload("Team").Preload("Room").Preload("Weekdays").Create(&per).Error; err != nil {
		return nil, err
	}
	return &per, nil
}
func (r *PermissionRepo) PutPermission(per models.Permission) (*models.Permission, error) {
	if err := r.db.Debug().Model(&per).Preload("User").Preload("Team").Preload("Room").Preload("Weekdays").Updates(&per).Error; err != nil {
		return nil, err
	}
	return &per, nil
}
func (r *PermissionRepo) DelPermission(id uint) (bool, error) {
	var permission models.Permission
	if err := r.db.Debug().Delete(&permission, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *PermissionRepo) GetUsersPermissions(UserID uint) ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.db.Debug().Preload("Room").Preload("Weekdays").Where("user_id = ? AND deleted_at IS NULL", UserID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
func (r *PermissionRepo) GetTeamsPermissions(TeamID uint) ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.db.Debug().Preload("Room").Preload("Weekdays").Where("team_id = ? AND deleted_at IS NULL", TeamID).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}
func (r *PermissionRepo) CleanPermission() error {
	return nil
}

func NewPermissionRepo(db *gorm.DB) *PermissionRepo {
	return &PermissionRepo{db}
}
