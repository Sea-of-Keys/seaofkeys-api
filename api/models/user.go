package models

import "gorm.io/gorm"


type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`

}
