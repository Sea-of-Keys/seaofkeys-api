package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type WebController struct {
	repo  *repos.WebRepo
	store *session.Store
}

func (con *WebController) GetPage(c *fiber.Ctx) error {
	// var data = models.UserPC
	// var userPC models.UserPC
	var err error
	token := c.Params("token")
	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	// sess.Set("Cfailed", false)
	// sess.Set("SetKronborg", 1)
	// if err := sess.Save(); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	// sess, err = con.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
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
	var Kfail bool
	sess, err = con.store.Get(c)
	if err != nil {
		panic(err)
	}
	Cfaill := sess.Get("Cfailed")
	sess, err = con.store.Get(c)
	if err != nil {
		panic(err)
	}
	Kfaill := sess.Get("Cfailed")
	if val, ok := Cfaill.(bool); ok {
		Cfail = val
	} else {
		Cfail = false
	}
	if val, ok := Kfaill.(bool); ok {
		Kfail = val
	} else {
		Kfail = false
	}
	// Kfail = Kfaill.(bool)
	fmt.Println("1")
	fmt.Println(sess.Get("SetToken"))
	data := fiber.Map{
		"User":    userPC,
		"Cfailed": Cfail,
		"Kfailed": Kfail,
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
	sess.Set("Cfailed", false)
	// if err := sess.Save(); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	// sess, err = con.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
	sess.Set("Kfailed", false)
	// if err := sess.Save(); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	// sess, err = con.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := sess.Save(); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	// sess, err = con.store.Get(c)
	// if err != nil {
	// 	panic(err)
	// }
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
	if FormData.PasswordOne != FormData.PasswordTwo {
		fmt.Println("3")
		// sess, err = con.store.Get(c)
		// if err != nil {
		// 	panic(err)
		// }
		sess.Set("Cfailed", true)
		// sess.Save()
		if FormData.CodeOne != FormData.CodeTwo && FormData.CodeOne != "" {
			// sess, err = con.store.Get(c)
			// if err != nil {
			// 	panic(err)
			// }
			fmt.Println("4")
			fmt.Println("4")
			fmt.Println("4")
			fmt.Println("4")
			sess.Set("Kfailed", true)
			// sess.Save()
		}
		// sess, err = con.store.Get(c)
		// if err != nil {
		// 	panic(err)
		// }
		getToken := sess.Get("SetToken")
		CToken := getToken.(string)
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

	return c.Redirect("https://api.seaofkeys.com")
}
func (con *WebController) PostPasswordAndCode2(c *fiber.Ctx) error {
	var FormData models.SetPasswordAndCode

	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	// sess.Set("password", 1)
	// sess.Set("code", 2)
	// sess.Set("SetToken", token)
	// if err := sess.Save(); err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	fmt.Printf("sess: %v\n", sess)
	fmt.Printf("sess: %v\n", sess)
	fmt.Printf("sess: %v\n", sess)
	fmt.Printf("sess: %v\n", sess)
	fmt.Printf("sess: %v\n", sess)
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
func (con *WebController) PostPasswordAndCode3(c *fiber.Ctx) error {
	var FormData models.SetPasswordAndCode

	// Retrieve the session once
	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}

	// Set session values
	sess.Set("Cfailed", false)
	sess.Set("Kfailed", false)

	// Save the session once after setting values
	if err := sess.Save(); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if sess.Get("SetToken") == nil {
		fmt.Println("1")
		fmt.Printf("sess: %v\n", sess)
		return fiber.NewError(fiber.StatusNetworkAuthenticationRequired, "C105: no session started")
	}

	if err := c.BodyParser(&FormData); err != nil {
		fmt.Println("2")
		return fiber.NewError(fiber.StatusInternalServerError, "C10: "+err.Error())
	}

	if FormData.PasswordOne != FormData.PasswordTwo {
		fmt.Println("3")

		// Set session values and save
		sess.Set("Cfailed", true)
		sess.Set("Kfailed", true)
		sess.Save()

		// Retrieve SetToken value
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

func RegisterWebController(reg models.RegisterController, store ...*session.Store) {

	repo := repos.NewWebRepo(reg.Db)
	controller := NewWebController(repo, reg.Store)

	WebRouter := reg.Router.Group("/web")
	WebRouter.Get("/token/:token?", controller.GetPage)
	WebRouter.Post("/set", controller.PostPasswordAndCode)
	// WebRouter.Use(security.TokenMiddleware(store))
	WebRouter.Get("/test/One", controller.TestOne)
	WebRouter.Get("/test/Two", controller.TestTwo)
}
