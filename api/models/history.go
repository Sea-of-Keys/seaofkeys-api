package models

type History struct {
	// gorm.Model
	ID           uint   `json:"id"            gorm:"primaryKey"`
	PermissionID uint   `json:"permission_id"`
	UserID       uint   `json:"user_id"`
	At           string `json:"at"`
	Permission   Permission
	User         User
}
