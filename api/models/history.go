package models

import "gorm.io/gorm"


type History struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`
	EmbeddedID Embedded `json:"embedded_id" gorm:"ForeignKey:ID"`
	UserID User `json:"user_id" gorm:"ForeignKey:ID"`
}
