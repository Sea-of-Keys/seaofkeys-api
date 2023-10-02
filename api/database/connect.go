package databae

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
}

func Init(database string) (*gorm.DB, error) {
	switch database {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("USER"),
			os.Getenv("PASSWORD"),
			os.Getenv("HOST"),
			os.Getenv("DBPORT"),
			os.Getenv("DATABASE"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err)
			panic("Failed to connect to database")
		}

		return db, err
	case "postgres":
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
			os.Getenv("USER"),
			os.Getenv("PASSWORD"),
			os.Getenv("HOST"),
			os.Getenv("DBPORT"),
			os.Getenv("DATABASE"),
		)
		// dsn1 := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
			panic("Failed to connect to database")
		}
		return db, err
	}
	log.Fatal("failed to connect to database")
	return nil, errors.New("can't connect to database: " + database)

}
