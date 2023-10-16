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
	// store := session.New(session.Config{
	// 	Expiration: 1 * time.Minute,
	// 	Storage:    storage,
	// })
	// store24Hours := session.New(session.Config{
	// 	Expiration: 1 * time.Minute,
	// 	Storage:    storage,
	// })
	controllers.RegisterAuthController(reg, stores[0])
	controllers.RegisterUserController(reg)
	controllers.RegisterEmbeddedController(reg, stores[0])
	controllers.RegisterTeamController(reg)
	controllers.RegisterHistoryController(reg, stores[0])
	controllers.RegisterRoomController(reg)
	controllers.RegisterStatsController(reg)
	controllers.RegisterPermissionController(reg)
	controllers.RegisterWebController(reg.Db, reg.Router, stores[0])

}

func main() {
	// db, err := databae.Init(os.Getenv("DATABASETYPE"))
	// pkg.SendEmail("mkronborg7@gmail.com", "Kronborg", "KronborgErGud!@#")
	// token, _ := security.NewEmbeddedToken()
	// fmt.Println(token)

	db, err := databae.Init("mysql")
	models.Setup(db)
	storage, err := databae.InitRedis()
	if err != nil {
		panic(err)
	}
	stores := []*session.Store{
		session.New(session.Config{
			Expiration: 32 * time.Hour,
			Storage:    storage,
		}),
		session.New(session.Config{
			Expiration: 15 * time.Minute,
			// Expiration: 24 * time.Hour,
			Storage: storage,
		}),
	}
	// fmt.Println("im gona be runed")
	app, err := initApp()
	if err != nil {
		log.Panic(err)
	}
	app.Use(logger.New())
	app.Use(cors.New())

	app.Static("/static", "./web/static")
	api := app.Group("/")
	reg := &models.RegisterController{
		Db:     db,
		Router: api,
		Store:  stores[1],
	}
	InitRoutes(*reg, stores)
	// InitRoutes(db, api, stores)

	log.Fatal(app.Listen(getPort()))

}

// func main() {
// 	// db, err := databae.Init(os.Getenv("DATABASETYPE"))
// 	db, err := databae.Init("mysql")
// 	models.Setup(db)

// 	// engine := CreateEngine()
// 	// app := fiber.New()
// 	// db, err := databae.Init("postgres")
// 	fmt.Println("im gona be runed")
// 	app, err := initApp()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	app.Use(logger.New())
// 	app.Use(cors.New())
// 	store := session.New(session.Config{
// 		KeyLookup:  "cookie:sessionid",
// 		Expiration: time.Hour * 24, // Session expiration time
// 	})
// 	// app.Use(store)
// 	app.Static("/static", "./web/static")
// 	api := app.Group("/")
// 	Endpoints(db, api, store)

// 	log.Fatal(app.Listen(getPort()))
// 	log.Fatal(app.Listen(os.Getenv("PORT")))
// 	log.Fatal(app.Listen(":8001"))
// 	log.Fatal(app.Listen(getPort()))

// }
