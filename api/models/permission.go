package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model
	ID          uint `json:"id"`
	RoomID      uint `json:"room_id"    gorm:"default:null"`
	Room        *Room
	TeamID      uint `json:"team_id"    gorm:"default:null"`
	UserID      uint `json:"user_id"    gorm:"default:null"`
	Team        *Team
	User        *User
	StartDateST string         `json:"start_date"`
	EndDateST   string         `json:"end_date"`
	StartDate   datatypes.Date `json:"-"`
	EndDate     datatypes.Date `json:"-"`
	// StartDate   time.Time      `json:"-"`
	// EndDate     time.Time      `json:"-"`
	StartTime datatypes.Time `json:"start_time"`
	EndTime   datatypes.Time `json:"end_time"`
	Weekdays  []*Weekdays    `json:"weekdays"   gorm:"many2many:permissions_weekdays;"`
}
