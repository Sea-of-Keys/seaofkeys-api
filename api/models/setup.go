package models

// ######TODO######

import (
	"time"

	"gorm.io/gorm"
)

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Room{},
		&Embedded{},
		&User{},
		&History{},
		&Team{},
		&Weekday{},
		&Permission{},
	)
	db.AutoMigrate(
		&Room{},
		&Embedded{},
		&User{},
		&History{},
		&Team{},
		&Weekday{},
		&Permission{},
	)

	room := []Room{
		{
			Name: "A1",
		},
		{
			Name: "A2",
		},
		{
			Name: "A3",
		},
	}
	weekdays := []Weekday{
		{
			Day: 1,
		},
		{
			Day: 2,
		},
		{
			Day: 3,
		},
		{
			Day: 4,
		},
		{
			Day: 5,
		},
		{
			Day: 6,
		},
		{
			Day: 7,
		},
	}
	embedded := []Embedded{
		{
			Name: "KeypadOne",
			// Room:  Room{Name: "A1"},
			RoomID: 1,
		},
		{
			Name: "RFIDOne",
			// Room: Room{Name: "dfg"},
			RoomID: 1,
		},
	}
	users := []User{
		{
			Name:     "Kronborg",
			Email:    "mkronborg7@gmail.com",
			Password: "Test",
			Code:     "2589",
		},
		{
			Name:     "Pissic",
			Email:    "IMGAY@gmail.com",
			Password: "Test",
			Code:     "2589",
		},
	}
	teams := []Team{
		{
			Name:  "Dev",
			Users: []User{{ID: 1}, {ID: 2}},
		},
		{
			Name:  "HR",
			Users: []User{{ID: 1}},
		},
	}
	history := []History{
		{
			EmbeddedID: 1,
			UserID:     1,
		},
		{
			EmbeddedID: 2,
			UserID:     1,
		},
		{
			EmbeddedID: 1,
			UserID:     2,
		},
	}
	permission := []Permission{
		{
			RoomID:    1,
			UserID:    1,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			StartTime: time.Now(),
			EndTime:   time.Now(),
			Weekdays:  []*Weekday{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		},
		{
			RoomID:    3,
			TeamID:    1,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			StartTime: time.Now(),
			EndTime:   time.Now(),
			Weekdays:  []*Weekday{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}},
		},
	}
	// create domy data in the database
	db.Create(&room)
	db.Create(&weekdays)
	db.Create(&embedded)
	db.Create(&users)
	db.Create(&history)
	db.Create(&teams)
	db.Create(&permission)
}
