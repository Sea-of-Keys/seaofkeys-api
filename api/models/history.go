package models

import "gorm.io/gorm"

type History struct {
	gorm.Model
	ID           uint `json:"id"            gorm:"primaryKey"`
	PermissionID uint `json:"permission_id"`
	UserID       uint `json:"user_id"`
	Permission   Permission
	User         User
}
