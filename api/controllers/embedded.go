package controllers

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
	"github.com/Sea-of-Keys/seaofkeys-api/pkg"
)

type EmbeddedController struct {
	repo  *repos.EmbeddedRepo
	store *session.Store
}

// Just for testing
type Login struct {
	// ID     uint   `json:"id"`
	Code   string `json:"code"`
	RoomID uint   `json:"room_id"`
	UserID uint   `json:"user_id"`
}

func (con *EmbeddedController) Setup(c *fiber.Ctx) error {
	// ####### TODO #######
	// make a table in database to check somfig
	// make a randowm token
	// make a session with that token
	// send the token back to the client
	// client/embedded use that token to encrypt the code
	var emb models.EmbedSetup
	sess, err := con.store.Get(c)
	if err != nil {
		return &pkg.CustomError{Code: "SES001", Message: "Failed to get session"}
	}
	if err := c.BodyParser(&emb); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if ok, err := con.repo.PostEmbeddedSetup(emb); !ok || err != nil {
		fmt.Println("Failed to setup")
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	base64, err := security.NewBase64Token()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	token, err := security.NewEmbeddedToken(base64)
	fmt.Printf("What is set as token %v\n", token)
	sess.Set("EmbeddedSession", token)
	if err := sess.Save(); err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	fmt.Println(token)

	return c.JSON(&fiber.Map{
		"session_is": true,
	})
}
func (con *EmbeddedController) Refrech(c *fiber.Ctx) error {
	return nil
}
func (con *EmbeddedController) Login(c *fiber.Ctx) error {
	var login Login
	if err := c.BodyParser(&login); err != nil {
		gg := errors.New("E22: " + err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, gg.Error())
	}
	return c.JSON(&fiber.Map{
		"message": "IT IS A LIVE",
		// "data":    sus,
	})
}
func (con *EmbeddedController) EmbeddedLoginLive(c *fiber.Ctx) error {
	var login models.EmbeddedLogin
	sess, err := con.store.Get(c)
	if err != nil {
		return nil
	}
	if sess.Get("EmbeddedSession") == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "E22: "+err.Error())
	}
	if err := c.BodyParser(&login); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "E22: "+err.Error())
	}
	fmt.Println("Kronborg")
	result := strings.Split(login.Code, "#")
	sus, err := con.repo.PostCodeLive(result[0], result[1], login.RoomID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(&fiber.Map{
		"message": "IT IS A LIVE",
		"data":    sus,
	})
}

func NewEmbeddedController(
	repo *repos.EmbeddedRepo,
	store *session.Store,
) EmbeddedInterfaceMethods {
	return &EmbeddedController{repo, store}
}

func RegisterEmbeddedController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewEmbeddedRepo(reg.Db)
	controller := NewEmbeddedController(repo, store[0])

	EmbeddedRouter := reg.Router.Group("/em")

	EmbeddedRouter.Post("/setup", controller.Setup)
	EmbeddedRouter.Use(security.TokenEmbeddedMiddleware(store[0]))
	EmbeddedRouter.Post("/login", controller.EmbeddedLoginLive)
	// EmbeddedRouter.Post("/log1", controller.EmbeddedLogin2)
	// EmbeddedRouter.Post("/log2", controller.EmbeddedLogin3)
	// EmbeddedRouter.Use(security.EmbeddedMiddleware(store[0]))
}
