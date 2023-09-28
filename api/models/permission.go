package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID        uint      `json:"id"`
	RoomID    Room      `json:"room_id"    gorm:"ForeignKey:ID"`
	TeamID    *Team     `json:"team_id"    gorm:"ForeignKey:ID"`
	UserID    *User     `json:"user_id"    gorm:"ForeignKey:ID"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Weekdays  []Weekday `json:"weekdays"   gorm:"ForeignKey:ID"`
}
