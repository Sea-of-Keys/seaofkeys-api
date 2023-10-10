package main

// ######TODO######

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/controllers"
	databae "github.com/Sea-of-Keys/seaofkeys-api/api/database"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return "0.0.0.0:" + port
}

func Endpoints(db *gorm.DB, api fiber.Router) {
	controllers.RegisterAuthController(db, api)
	controllers.RegisterUserController(db, api)
	controllers.RegisterEmbeddedController(db, api)
	controllers.RegisterTeamController(db, api)
	controllers.RegisterHistoryController(db, api)
	controllers.RegisterRoomController(db, api)
	controllers.RegisterStatsController(db, api)

}

func main() {
	// db, err := databae.Init(os.Getenv("DATABASETYPE"))
	db, err := databae.Init("mysql")
	fmt.Printf("The type of myVar is: %T\n", db)
	models.Setup(db)
	app := fiber.New()
	// db, err := databae.Init("postgres")
	if err != nil {
		log.Panic(err)
	}
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/")
	Endpoints(db, api)

	log.Fatal(app.Listen(getPort()))

}
