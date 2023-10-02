package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint    `json:"id"       gorm:"primaryKey"`
	Email    string  `json:"email"    gorm:"uniqueIndex;size:191"`
	Password string  `json:"password"`
	Code     string  `json:"code"     gorm:"not null;"`
	Name     string  `json:"name"`
	Teams    []*Team `                gorm:"many2many:teams_users;"`
}
