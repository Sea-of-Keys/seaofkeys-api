package models

type Team struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"many2many:teams_users;"`
}
