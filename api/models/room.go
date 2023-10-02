package models

type Room struct {
	ID   uint   `json:"id"   gorm:"primaryKey"`
	Name string `json:"name"`
}
