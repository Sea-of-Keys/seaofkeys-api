package models

// ######TODO######

import (
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/middleware"
)

func Setup(db *gorm.DB) {
	db.Migrator().DropTable(
		&Room{},
		&Embedded{},
		&User{},
		&History{},
		&Team{},
		&Weekdays{},
		&Permission{},
		&UserPC{},
	)
	db.AutoMigrate(
		&Room{},
		&Embedded{},
		&User{},
		&History{},
		&Team{},
		&Weekdays{},
		&Permission{},
		&UserPC{},
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
	weekday := []Weekdays{
		{
			Day:  1,
			Name: "Monday",
		},
		{
			Day:  2,
			Name: "Tuesday",
		},
		{
			Day:  3,
			Name: "Wensday",
		},
		{
			Day:  4,
			Name: "Thursday",
		},
		{
			Day:  5,
			Name: "Friday",
		},
		{
			Day:  6,
			Name: "Saturday",
		},
		{
			Day:  7,
			Name: "Sunday",
		},
	}
	var RoomOne uint
	RoomOne = 1

	embedded := []Embedded{
		{
			Name: "KeypadOne",
			// Room:  Room{Name: "A1"},
			RoomID: &RoomOne,
		},
		{
			Name: "RFIDOne",
			// Room: Room{Name: "dfg"},
			RoomID: &RoomOne,
		},
	}
	userOneCode, err := middleware.HashPassword("2589")
	if err != nil {
		log.Panic(err)
	}
	userOnePassword, err := middleware.HashPassword("Test")
	// fmt.Println(userOnePassword)
	if err != nil {
		log.Panic(err)
	}
	userTwoCode, err := middleware.HashPassword("1234")
	if err != nil {
		log.Panic(err)
	}
	userTreCode, err := middleware.HashPassword("4444")
	if err != nil {
		log.Panic(err)
	}
	emailOne := "mkronborg7@gmail.com"
	emailTwo := "IMGAY@viki.lover"
	emailTre := "gg@wp.ez"
	users := []User{
		{
			Name:     "Kronborg",
			Email:    &emailOne,
			Password: &userOnePassword,
			Code:     &userOneCode,
			// Teams:    []Team{{ID: 1}, {ID: 2}},
		},
		{
			Name:  "Pissic",
			Email: &emailTwo,
			// Password: "Test",
			Code: &userTwoCode,
			// Teams: []Team{{ID: 1}},
		},
		{
			Name:     "Test",
			Email:    &emailTre,
			Password: &userOnePassword,
			Code:     &userTreCode,
			// Teams:    []Team{{ID: 1}, {ID: 2}},
		},
	}

	teams := []Team{
		{
			Name:  "Dev",
			Users: []*User{{ID: 1}, {ID: 2}},
		},
		{
			Name:  "HR",
			Users: []*User{{ID: 1}},
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
		// {
		// 	RoomID: 1,
		// 	// UserID:    1,
		// 	TeamID:    1,
		// 	StartDate: time.Now(),
		// 	EndDate:   time.Now(),
		// 	StartTime: time.Now(),
		// 	EndTime:   time.Now(),
		// 	// Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		// },
		{
			RoomID: 1,
			UserID: 1,
			// TeamID:    1,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			StartTime: time.Now(),
			EndTime:   time.Now(),
			Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		},
		{
			RoomID: 1,
			UserID: 2,
			// TeamID:    1,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			StartTime: time.Now(),
			EndTime:   time.Now(),
			// Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		},
		{
			// RoomID:    &RoomOne,
			UserID:    1,
			StartDate: time.Now(),
			EndDate:   time.Now(),
			StartTime: time.Now(),
			EndTime:   time.Now(),
			Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}},
		},
	}
	// create domy data in the database
	db.Create(&room)
	db.Create(&weekday)
	db.Create(&embedded)
	db.Create(&users)
	db.Create(&history)
	db.Create(&teams)
	db.Create(&permission)
}
