package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint    `json:"id"       gorm:"primaryKey"`
	Name     string  `json:"name"`
	Email    *string `json:"email"    gorm:"uniqueIndex;size:191"`
	Password *string `json:"password"`
	Code     *string `json:"code"     gorm:"uniqueIndex;size:100"` // not null; MABY NOT NULL
	RFID     *string `json:"rfid"`
	Teams    []*Team `                gorm:"many2many:teams_users;"`
}

type Delete struct {
	ID uint `json:"id"`
}
