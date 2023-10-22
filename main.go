package main

// ######TODO######

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/django/v3"

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

func InitRoutes(reg models.RegisterController, stores []*session.Store) {
	controllers.RegisterAuthController(reg)
	controllers.RegisterUserController(reg)
	controllers.RegisterEmbeddedController(reg, stores[0])
	controllers.RegisterTeamController(reg)
	controllers.RegisterHistoryController(reg)
	controllers.RegisterRoomController(reg)
	controllers.RegisterStatsController(reg)
	controllers.RegisterPermissionController(reg)
	controllers.RegisterWebController(reg)

}

func main() {

	db, err := databae.Init("mysql")
	if err != nil {
		panic(err)
	}
	if os.Getenv("DEV") == "dev" {
		models.Setup(db)
	}
	storage, err := databae.InitRedis()
	if err != nil {
		panic(err)
	}
	stores := []*session.Store{
		session.New(session.Config{
			KeyLookup:  "cookie:session_em_id",
			Expiration: 32 * time.Hour,
			Storage:    storage,
		}),
		session.New(session.Config{
			KeyLookup:  "cookie:kronborg_id",
			Expiration: 5 * time.Hour,
			Storage:    storage,
		}),
	}
	app, err := initApp()
	if err != nil {
		log.Panic(err)
	}
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:8000, http://localhost:8000, http://127.0.0.1, https://seaofkeys.com, https://www.seaofkeys.com, http://localhost:8006",
		AllowCredentials: true,
	}))
	app.Static("/static", "./web/static")
	api := app.Group("/")
	reg := &models.RegisterController{
		Db:     db,
		Router: api,
		Store:  stores[1],
	}
	InitRoutes(*reg, stores)

	log.Fatal(app.Listen(getPort()))

}
