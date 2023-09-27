package models


type Embedded struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	// RoomID uint `json:"room_id"`
	Room uint `json:"room_id" gorm:"ForeignKey:ID"`
}
