package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type WebController struct {
	repo     *repos.WebRepo
	userRepo *repos.UserRepo
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
	fmt.Printf("sess: %v\n", sess)
	userPC, err := con.repo.GetCheckToken(token)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"user is not set to get a password or code",
		)
	}
	fmt.Println("1")
	getToken := sess.Get("SetToken")
	if getToken == nil {
		sess.Set("SetToken", token)
		if err := sess.Save(); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		sess, err = con.store.Get(c)
		if err != nil {
			panic(err)
		}
		getToken = sess.Get("SetToken")
	}

	fmt.Println("1")
	CToken := getToken.(string)
	if CToken != userPC.Token {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"sessin token does not match provide token",
		)
	}
	var Cfail bool
	sess, err = con.store.Get(c)
	if err != nil {
		panic(err)
	}
	Cfaill := sess.Get("Cfailed")
	if val, ok := Cfaill.(bool); ok {
		Cfail = val
	} else {
		Cfail = false
	}
	fmt.Println("1")
	fmt.Println(sess.Get("SetToken"))
	data := fiber.Map{
		"User":    userPC,
		"Cfailed": Cfail,
	}
	fmt.Printf("data: %v\n", data)
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

	if sess.Get("SetToken") == nil {
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

	getToken := sess.Get("SetToken")
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
	repo *repos.WebRepo,
	userRepo *repos.UserRepo,
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
	WebRouter.Post("/set", controller.PostNewCodes)
	// WebRouter.Use(security.TokenMiddleware(store))
	WebRouter.Get("/home", controller.TestOne)
	WebRouter.Get("/error", controller.Error)
}
