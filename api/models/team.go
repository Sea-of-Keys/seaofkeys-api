package models


type Team struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	
}
