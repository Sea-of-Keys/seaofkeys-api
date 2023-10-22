package models

// ######TODO######

import (
	"log"
	"time"

	"gorm.io/datatypes"
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
		"teams_users",
		"permissions_weekdays",
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
	currentTime := time.Now()
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
			Name: "Wednesday",
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
	// hashed, err := middleware.HashPassword("Kronborg")
	// if err != nil {
	// 	panic(err)
	// }
	embedded := []Embedded{
		{
			Name: "KeypadOne123",
			// Room:  Room{Name: "A1"},
			RoomID: &RoomOne,
			Scret:  "kronborg",
		},
		{
			Name: "RFIDOne",
			// Room: Room{Name: "dfg"},
			RoomID: &RoomOne,
		},
		// {
		// 	Name: "TestOne",
		// 	// Room: Room{Name: "dfg"},
		// 	RoomID: &RoomOne,
		// 	Scret:  "Kronborg er kogen",
		// },
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
	layout := "2006-01-02 15:04:05"
	timeTwo := currentTime.Add(-2 * time.Hour)
	timeTre := currentTime.Add(-1 * time.Hour)
	formattedTime := currentTime.Format(layout)
	formattedTimeTwo := timeTwo.Format(layout)
	formattedTimeTre := timeTre.Format(layout)
	history := []History{
		{
			PermissionID: 1,
			UserID:       1,
			At:           formattedTime,
		},
		{
			PermissionID: 2,
			UserID:       1,
			At:           formattedTimeTwo,
		},
		{
			PermissionID: 1,
			UserID:       2,
			At:           formattedTimeTre,
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
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			// StartDate: time.Now(),
			// EndDate:   time.Now(),
			StartTime: datatypes.NewTime(9, 0, 0, 0),
			EndTime:   datatypes.NewTime(22, 0, 0, 0),
			Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		},
		{
			RoomID: 1,
			UserID: 2,
			// TeamID:    1,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			// StartDate: time.Now(),
			// EndDate:   time.Now(),
			StartTime: datatypes.NewTime(10, 0, 0, 0),
			EndTime:   datatypes.NewTime(12, 0, 0, 0),
			// Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		},
		{
			// RoomID:    &RoomOne,
			RoomID:      2,
			UserID:      1,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			// StartDate: time.Now(),
			// EndDate:   time.Now(),
			StartTime: datatypes.NewTime(0, 0, 0, 0),
			EndTime:   datatypes.NewTime(23, 59, 0, 0),
			Weekdays:  []*Weekdays{{ID: 1}, {ID: 2}},
		},
		{
			RoomID: 3,
			// UserID:    1,
			TeamID:      1,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			// StartDate: time.Now(),
			// EndDate:   time.Now(),
			StartTime: datatypes.NewTime(6, 0, 0, 0),
			EndTime:   datatypes.NewTime(23, 30, 0, 0),
			Weekdays:  []*Weekdays{{ID: 1}, {ID: 5}},
		},
	}
	userpc := []UserPC{
		{
			UserID:    1,
			Token:     "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NCwiRW1haWwiOiJta3JvbmJvcmc2NkBnbWFpbC5jb20iLCJleHAiOjE2OTc0NjkwOTV9.-0JXwf6-vAKuCxB8g0br9ZVaWvHUOHQq7ikRr2EbVJk",
			Password:  true,
			EmailSend: true,
		},
	}
	// create domy data in the database
	db.Create(&room)
	db.Create(&weekday)
	db.Create(&embedded)
	db.Create(&users)
	db.Create(&userpc)
	db.Create(&teams)
	db.Create(&permission)
	db.Create(&history)
}
