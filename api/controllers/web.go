package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type WebController struct {
	repo     repos.WebRepoInterface
	userRepo repos.UserRepoInterface
	store    *session.Store
}

func (con *WebController) GetPage(c *fiber.Ctx) error {
	var err error
	token := c.Params("token")
	sess, err := con.store.Get(c)
	if err != nil {
		data := fiber.Map{
			"message": "failed to get session",
		}
		return c.Render("error/index", data)
	}
	fmt.Printf("token: %v\n", token)
	userPC, err := con.repo.GetCheckToken(token)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"user is not set to get a password or code",
		)
	}
	fmt.Println("1")
	getToken := sess.Get("WebToken")
	if getToken == nil {
		sess.Set("WebToken", token)
		if err := sess.Save(); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		sess, err = con.store.Get(c)
		if err != nil {
			panic(err)
		}
		fmt.Println("2")
		getToken = sess.Get("WebToken")
	}

	fmt.Println("1")
	CToken := getToken.(string)
	fmt.Printf("CToken: %v\n", CToken)

	fmt.Println("2")
	if CToken != userPC.Token {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"sessin token does not match provide token",
		)
	}
	fmt.Println("3")
	var Cfail bool
	sess, err = con.store.Get(c)
	if err != nil {
		panic(err)
	}
	fmt.Println("4")
	Cfaill := sess.Get("Cfailed")
	if val, ok := Cfaill.(bool); ok {
		Cfail = val
	} else {
		Cfail = false
	}
	fmt.Println("5")

	data := fiber.Map{
		"User":    userPC,
		"Cfailed": Cfail,
	}
	fmt.Printf("data: %v\n", data)
	if err := sess.Save(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Render("web/index", data)
}
func (con *WebController) PostNewCodes(c *fiber.Ctx) error {
	var FormData models.SetPasswordAndCode

	sess, err := con.store.Get(c)
	if err != nil {
		data := fiber.Map{
			"message": "failed to get session",
		}
		return c.Render("error/index", data)
	}
	sess.Set("Cfailed", false)

	if sess.Get("WebToken") == nil {
		fmt.Println("1")
		fmt.Printf("sess: %v\n", sess)
		if err := sess.Save(); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return fiber.NewError(fiber.StatusNetworkAuthenticationRequired, "C105: no session started")
	}
	if err := c.BodyParser(&FormData); err != nil {
		fmt.Println("2")
		if err := sess.Save(); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, "C10: "+err.Error())
	}

	getToken := sess.Get("WebToken")
	CToken := getToken.(string)
	if FormData.CodeOne != FormData.CodeTwo || FormData.CodeOne == "" {
		fmt.Println("3")

		sess.Set("Cfailed", true)
		if err := sess.Save(); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.Redirect(fmt.Sprintf("/web/token/%v", CToken))
	}

	fmt.Printf(
		"PasswordOne: %v\nPaswordTwo: %v\nCodeOne: %v\nCodeTwo: %v\n",
		FormData.PasswordOne,
		FormData.PasswordTwo,
		FormData.CodeOne,
		FormData.CodeTwo,
	)
	if err := sess.Save(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println("4")
	fmt.Println("4")
	fmt.Println("4")
	fmt.Println("4")
	if ok, err := con.userRepo.PutPassword(FormData.CodeOne, CToken); ok || err == nil {
		return c.Redirect("/web/home")
	}

	return c.Redirect(fmt.Sprintf("/web/token/%v", CToken))
}

func (con *WebController) TestOne(c *fiber.Ctx) error {
	sess, err := con.store.Get(c)
	if err != nil {
		data := fiber.Map{
			// "message": "failed to get session",
		}
		return c.Render("error/index", data)
	}
	fmt.Println(sess)
	data := fiber.Map{}
	return c.Render("home/index", data)
}
func (con *WebController) TestTwo(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"name": "name",
	})
}
func (con *WebController) Error(c *fiber.Ctx) error {
	data := fiber.Map{}

	return c.Render("error/index", data)
}
func NewWebController(
	repo repos.WebRepoInterface,
	userRepo repos.UserRepoInterface,
	store *session.Store,
) WebInterfaceMethods {
	return &WebController{repo, userRepo, store}
}

func RegisterWebController(reg models.RegisterController, store ...*session.Store) {

	repo := repos.NewWebRepo(reg.Db)
	userRepo := repos.NewUserRepo(reg.Db)
	controller := NewWebController(repo, userRepo, reg.Store)

	WebRouter := reg.Router.Group("/web")
	WebRouter.Get("/token/:token?", controller.GetPage)
	WebRouter.Use(security.WebsiteTokenMiddleware(reg.Store))
	WebRouter.Post("/set", controller.PostNewCodes)
	WebRouter.Get("/home", controller.TestOne)
	WebRouter.Get("/error", controller.Error)
}
