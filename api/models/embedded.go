package models

type Embedded struct {
	ID     uint   `json:"id"      gorm:"primaryKey"`
	Name   string `json:"name"`
	RoomID *uint  `json:"room_id"`
	Room   Room
	Scret  string `json:"-"`
}

type EmbeddedCheck struct {
	ID         uint   `json:"id"          gorm:"primaryKey"`
	EmbeddedID uint   `json:"embedded_id"`
	Scret      string `json:"-"`
	Embedded   Embedded
}
