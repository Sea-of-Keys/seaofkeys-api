package repos

import "gorm.io/gorm"

type PermissionRepo struct {
	db *gorm.DB
}
