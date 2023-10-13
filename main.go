package main

// ######TODO######

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/django/v3"
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
func initApp() (*fiber.App, error) {
	engine := django.New("./web/views", ".html")
	engine.Reload(true)
	engine.AddFunc("css", func(name string) (res template.HTML) {
		link := "static/" + name
		res = template.HTML("<link rel=\"stylesheet\" href=\"/" + link + "\">")
		return
	})
	app := fiber.New(fiber.Config{
		PassLocalsToViews: true,
		Views:             engine,
	})
	return app, nil
}

func Endpoints(db *gorm.DB, api fiber.Router, store *session.Store) {
	controllers.RegisterAuthController(db, api)
	controllers.RegisterUserController(db, api)
	controllers.RegisterEmbeddedController(db, api)
	controllers.RegisterTeamController(db, api)
	controllers.RegisterHistoryController(db, api, store)
	controllers.RegisterRoomController(db, api)
	controllers.RegisterStatsController(db, api)
	controllers.RegisterPermissionController(db, api)
	controllers.RegisterWebController(db, api, store)

}

func main() {
	// db, err := databae.Init(os.Getenv("DATABASETYPE"))
	db, err := databae.Init("mysql")
	models.Setup(db)

	// engine := CreateEngine()
	// app := fiber.New()
	// db, err := databae.Init("postgres")
	fmt.Println("im gona be runed")
	app, err := initApp()
	if err != nil {
		log.Panic(err)
	}
	app.Use(logger.New())
	app.Use(cors.New())
	store := session.New(session.Config{
		KeyLookup:  "cookie:sessionid",
		Expiration: time.Hour * 24, // Session expiration time
	})

	app.Static("/static", "./web/static")
	api := app.Group("/")
	Endpoints(db, api, store)

	log.Fatal(app.Listen(getPort()))

}
