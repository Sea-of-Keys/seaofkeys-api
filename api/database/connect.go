package databae

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/gofiber/storage/redis/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ######TODO######
// Fix Postgres db

type Db struct {
}

func InitRedis() (*redis.Storage, error) {
	port, err := strconv.Atoi(os.Getenv("REDISPORT"))
	if err != nil {
		port = 6379
	}
	storage := redis.New(redis.Config{
		Host:      os.Getenv("REDISHOST"),
		Port:      port,
		Username:  os.Getenv("REDISUSER"),
		Password:  os.Getenv("REDISPASSWORD"),
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
	return storage, nil
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
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Copenhagen",
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
		db.Exec("SET client_encoding TO 'UTF8'")
		return db, err
	}
	log.Fatal("failed to connect to database")
	return nil, errors.New("can't connect to database: " + database)

}
