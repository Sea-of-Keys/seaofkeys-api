package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type WebController struct {
	repo *repos.WebRepo
}

func (con *WebController) GetPage(c *fiber.Ctx) error {
	// session, err := c.Locals("session").(*session.Session)
	store := session.New(session.Config{
		KeyLookup:  "cookie:sessionid",
		Expiration: time.Hour * 24, // Session expiration time
	})
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Set("name", "john")

	data := fiber.Map{
		"password": true,
	}
	return c.Render("web/index", data)
}
func (con *WebController) PostPasswordAndCode(c *fiber.Ctx) error {
	var FormData models.SetPasswordAndCode

	if err := c.BodyParser(&FormData); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C10: "+err.Error())
	}

	fmt.Printf(
		"PasswordOne: %v\nPaswordTwo: %v\nCodeOne: %v\nCodeTwo: %v\n",
		FormData.PasswordOne,
		FormData.PasswordTwo,
		FormData.CodeOne,
		FormData.CodeTwo,
	)

	return c.Redirect("https://api.seaofkeys.com")
}

func NewWebController(repo *repos.WebRepo) *WebController {
	return &WebController{repo}
}

func RegisterWebController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewWebRepo(db)
	controller := NewWebController(repo)

	// store := session.New()
	WebRouter := router.Group("/web")
	// WebRouter.Use(session.New(session.Config{
	// 	KeyLookup:  "cookie:sessionid",
	// 	Expiration: time.Hour * 24, // Session expiration time
	// }))

	WebRouter.Get("/", controller.GetPage)
	WebRouter.Post("/set", controller.PostPasswordAndCode)
}
