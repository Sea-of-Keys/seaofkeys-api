package models

type Team struct {
	ID    uint   `json:"id"    gorm:"primaryKey"`
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"many2many:team_user;"`
}
