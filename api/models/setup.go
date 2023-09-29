package models

// ######TODO######

import "gorm.io/gorm"

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
	}
	// create domy data in the database
	db.Create(&room)
	db.Create(&weekdays)
	db.Create(&embedded)
	db.Create(&users)
	// db.Create(&history)
	// db.Create(&team)
	// db.Create(&history)
	// db.Create(&permission)
}
