package models

// ######TODO######

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	expirationTime := time.Now().Add(32 * time.Hour)
	claims := &Claims{
		ID:    1,
		Email: "mkronborg7@gmail.com",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("PSCRERT")))
	if err != nil {
		panic(err)
	}
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
			Day:  0,
			Name: "Sunday",
		},
	}
	var RoomOne uint
	RoomOne = 1

	embedded := []Embedded{
		{
			Name:   "KeypadOne123",
			RoomID: &RoomOne,
			Scret:  "kronborg",
		},
		{
			Name:   "RFIDOne",
			RoomID: &RoomOne,
		},
	}
	userOneCode, err := middleware.HashPassword("2589")
	if err != nil {
		log.Panic(err)
	}
	userOnePassword, err := middleware.HashPassword("Test")
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
		},
		{
			Name:  "Pissic",
			Email: &emailTwo,
			Code:  &userTwoCode,
		},
		{
			Name:     "Test",
			Email:    &emailTre,
			Password: &userOnePassword,
			Code:     &userTreCode,
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
		{
			RoomID: 1,
			UserID: 1,
			// TeamID:    1,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			StartTime:   datatypes.NewTime(9, 0, 0, 0),
			EndTime:     datatypes.NewTime(22, 0, 0, 0),
			Weekdays:    []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7}},
		},
		{
			RoomID:      1,
			UserID:      2,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			StartTime:   datatypes.NewTime(10, 0, 0, 0),
			EndTime:     datatypes.NewTime(12, 0, 0, 0),
			Weekdays:    []*Weekdays{{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}},
		},
		{
			RoomID:      2,
			UserID:      1,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			StartTime:   datatypes.NewTime(0, 0, 0, 0),
			EndTime:     datatypes.NewTime(23, 59, 0, 0),
			Weekdays:    []*Weekdays{{ID: 1}, {ID: 2}},
		},
		{
			RoomID:      3,
			TeamID:      1,
			StartDate:   datatypes.Date(currentTime),
			EndDate:     datatypes.Date(currentTime),
			StartDateST: currentTime.Format("2006-01-02"),
			EndDateST:   currentTime.Format("2006-01-02"),
			StartTime:   datatypes.NewTime(6, 0, 0, 0),
			EndTime:     datatypes.NewTime(23, 30, 0, 0),
			Weekdays:    []*Weekdays{{ID: 1}, {ID: 5}},
		},
	}
	userpc := []UserPC{
		{
			UserID:    1,
			Token:     tokenString,
			Password:  true,
			EmailSend: true,
		},
	}
	db.Create(&room)
	db.Create(&weekday)
	db.Create(&embedded)
	db.Create(&users)
	db.Create(&userpc)
	db.Create(&teams)
	db.Create(&permission)
	db.Create(&history)
}
