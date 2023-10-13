package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type WebController struct {
	repo  *repos.WebRepo
	store *session.Store
}

func (con *WebController) GetPage(c *fiber.Ctx) error {
	// session, err := c.Locals("session").(*session.Session)
	token := c.Params("+")
	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	// Check Token HERE
	// if token is legiget set token in session
	sess.Set("token", token)
	sess.Save()

	//this is just to test if i get a token
	// Read and output the session variable
	sess, _ = con.store.Get(c)
	name := sess.Get("token")
	fmt.Printf("Name from session: %v\n", name)

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
func (con *WebController) TestOne(c *fiber.Ctx) error {

	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	// Check Token HERE
	// if token is legiget set token in session

	//this is just to test if i get a token
	// Read and output the session variable
	name := sess.Get("token")
	fmt.Printf("Name from session: %v\n", name)

	return c.JSON(&fiber.Map{
		"name": name,
	})
}
func (con *WebController) TestTwo(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"name": "name",
	})
}

func NewWebController(repo *repos.WebRepo, store *session.Store) *WebController {
	return &WebController{repo, store}
}

func RegisterWebController(db *gorm.DB, router fiber.Router, store *session.Store) {
	// store := session.New(session.Config{
	// 	KeyLookup:  "cookie:sessionid",
	// 	Expiration: time.Hour * 24, // Session expiration time
	// })
	repo := repos.NewWebRepo(db)
	controller := NewWebController(repo, store)

	WebRouter := router.Group("/web")
	WebRouter.Get("/set/+", controller.GetPage)
	WebRouter.Post("/set", controller.PostPasswordAndCode)
	WebRouter.Get("/test/One", controller.TestOne)
	WebRouter.Get("/test/Two", controller.TestTwo)
}
