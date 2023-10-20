package databae

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/gofiber/storage/redis/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ######TODO######
// Fix Postgres db

type Db struct {
}

type SingletonRedis struct {
	Storage *redis.Storage
}

var (
	redisInstance *SingletonRedis
	redisOnce     sync.Once
)

func InitRedis() (*redis.Storage, error) {
	var err error
	redisOnce.Do(func() {
		port, err := strconv.Atoi(os.Getenv("REDISPORT"))
		if err != nil {
			port = 6379
		}
		// time.Sleep(3 * time.Second)
		storage := redis.New(redis.Config{
			Host:      os.Getenv("REDISHOST"),
			Port:      port,
			Username:  os.Getenv("REDISUSER"),
			Password:  os.Getenv("REDISPASSWORD"),
			Database:  0,
			Reset:     false,
			TLSConfig: nil,
			PoolSize:  80,
		})

		redisInstance = &SingletonRedis{
			Storage: storage,
		}
	})

	return redisInstance.Storage, err
}

type Singleton struct {
	DB *gorm.DB
}

var (
	instance *Singleton
	once     sync.Once
)

func Init(database string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	once.Do(func() {
		switch database {
		case "mysql":
			dsn := fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				os.Getenv("MYSQLUSER"),
				os.Getenv("MYSQLPASSWORD"),
				os.Getenv("MYSQLHOST"),
				os.Getenv("MYSQLPORT"),
				os.Getenv("MYSQLDATABASE"),
			)
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		case "postgres":
			dsn := fmt.Sprintf(
				"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=Europe/Copenhagen",
				os.Getenv("USER"),
				os.Getenv("PASSWORD"),
				os.Getenv("HOST"),
				os.Getenv("DBPORT"),
				os.Getenv("DATABASE"),
			)
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			db.Exec("SET client_encoding TO 'UTF8'")
		default:
			log.Fatal("Invalid database type")
		}
		instance = &Singleton{
			DB: db,
		}
	})

	return instance.DB, err
}
