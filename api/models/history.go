package models

import "gorm.io/gorm"


type History struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	EmbeddedID uint `json:"embedded_id"`
	UserID uint `json:"user_id"`
}
