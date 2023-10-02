package models

import "time"

type Permission struct {
	// gorm.Model
	ID        uint `json:"id"`
	RoomID    uint `json:"room_id"`
	Room      Room
	TeamID    uint `json:"team_id"    gorm:"default:null"`
	UserID    uint `json:"user_id"    gorm:"default:null"`
	Team      *Team
	User      *User
	StartDate time.Time  `json:"start_date"`
	EndDate   time.Time  `json:"end_date"`
	StartTime time.Time  `json:"start_time"`
	EndTime   time.Time  `json:"end_time"`
	Weekdays  []*Weekday `json:"weekdays"   gorm:"many2many:permissions_weekdays;"`
}
