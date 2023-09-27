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

	// create domy data in the database
	// db.Create(&room)
	// db.Create(&embedded)
	// db.Create(&user)
	// db.Create(&history)
	// db.Create(&team)
	// db.Create(&history)
	// db.Create(&permission)
}
