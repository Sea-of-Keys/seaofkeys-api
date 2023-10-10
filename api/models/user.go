package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint    `json:"id"    gorm:"primaryKey"`
	Name     string  `json:"name"`
	Email    *string `json:"email" gorm:"uniqueIndex;size:191"`
	Password *string `json:"-"`
	Code     *string `json:"-"     gorm:"uniqueIndex;size:100"` // not null; MABY NOT NULL
	RFID     *string `json:"-"`
	Teams    []*Team `json:"teams" gorm:"many2many:teams_users;"`
}
