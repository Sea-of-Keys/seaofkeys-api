package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type WebController struct {
	repo  *repos.WebRepo
	store *session.Store
}

func (con *WebController) GetPage(c *fiber.Ctx) error {
	// var data = models.UserPC
	// var userPC models.UserPC
	// var err error
	token := c.Params("token")
	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	userPC, err := con.repo.GetCheckToken(token)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"user is not set to get a password or code",
		)
	}
	getToken := sess.Get("SetToken")
	if getToken == nil {
		sess.Set("SetToken", userPC.Token)
		sess.Save()
		sess, err = con.store.Get(c)
		if err != nil {
			panic(err)
		}
		getToken = sess.Get("SetToken")
	}
	CToken := getToken.(string)
	if CToken != userPC.Token {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"sessin token does not match provide token",
		)
	}
	fmt.Println(sess.Get("SetToken"))
	data := fiber.Map{
		"User": userPC,
	}
	fmt.Printf("data: %v\n", data)
	return c.Render("web/index", data)
}
func (con *WebController) PostPasswordAndCode(c *fiber.Ctx) error {
	var FormData models.SetPasswordAndCode

	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	if sess.Get("SetToken") == nil {
		fmt.Println("1")
		return fiber.NewError(fiber.StatusNetworkAuthenticationRequired, "C105: no session started")
	}
	if err := c.BodyParser(&FormData); err != nil {
		fmt.Println("2")
		return fiber.NewError(fiber.StatusInternalServerError, "C10: "+err.Error())
	}
	if FormData.PasswordOne != FormData.PasswordTwo || FormData.CodeOne != FormData.CodeTwo {
		fmt.Println("3")

		getToken := sess.Get("SetToken")
		CToken := getToken.(string)
		return c.Redirect(fmt.Sprintf("/web/token/%v", CToken))
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
	fmt.Printf("sess: %v\n", sess)
	// Check Token HERE
	// if token is legiget set token in session

	//this is just to test if i get a token
	// Read and output the session variable
	name := sess.Get("ActiveToken")
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
	WebRouter.Get("/token/:token?", controller.GetPage)
	WebRouter.Post("/set", controller.PostPasswordAndCode)
	WebRouter.Use(security.TokenMiddleware(store))
	WebRouter.Get("/test/One", controller.TestOne)
	WebRouter.Get("/test/Two", controller.TestTwo)
}
