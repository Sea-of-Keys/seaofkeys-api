package main

// ######TODO######

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

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

func main() {
	// db, err := databae.Init(os.Getenv("DATABASETYPE"))
	db, err := databae.Init("mysql")
	models.Setup(db)
	app := fiber.New()
	// db, err := databae.Init("postgres")
	if err != nil {
		log.Panic(err)
	}
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/")

	controllers.RegisterAuthController(db, api)
	controllers.RegisterUserController(db, api)
	controllers.RegisterEmbeddedController(db, api)
	controllers.RegisterTeamController(db, api)
	controllers.RegisterHistoryController(db, api)

	log.Fatal(app.Listen(getPort()))

}
