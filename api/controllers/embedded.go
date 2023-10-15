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
	sess, err := con.store.Get(c)
	if err != nil {
		return &pkg.CustomError{Code: "SES001", Message: "Failed to get session"}
	}
	g, err := security.NewEmbeddedToken()
	sess.Set("EmbeddedSession", g)
	sess.Save()
	fmt.Println(g)
	if err != nil {
		return &pkg.CustomError{Code: "S001", Message: "Failed to token"}
	}
	return nil
}
func (con *EmbeddedController) EmbeededLogin(c *fiber.Ctx) error {
	var login Login
	if err := c.BodyParser(&login); err != nil {
		gg := errors.New("E22: " + err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, gg.Error())
	}
	sus, err := con.repo.PostCodeV2(login.Code, "g", login.RoomID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// if !sus {
	// 	c.JSON(&fiber.Map{
	// 		"message": "it's dead :-(",
	// 	})
	// }
	return c.JSON(&fiber.Map{
		"message": "IT IS A LIVE",
		"data":    sus,
	})
}

func (con *EmbeddedController) EmbeddedLogin2(c *fiber.Ctx) error {
	var login models.EmbeddedLogin
	if err := c.BodyParser(&login); err != nil {
		// gg := errors.New("E22: " + err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, "E22: "+err.Error())
	}
	result := strings.Split(login.Code, "#")
	sus, err := con.repo.PostCodeV3(result[0], result[1], login.RoomID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(&fiber.Map{
		"message": "IT IS A LIVE",
		"data":    sus,
	})
}
func NewEmbeddedController(repo *repos.EmbeddedRepo, store *session.Store) *EmbeddedController {
	return &EmbeddedController{repo, store}
}

func RegisterEmbeddedController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewEmbeddedRepo(reg.Db)
	controller := NewEmbeddedController(repo, store[0])

	EmbeddedRouter := reg.Router.Group("/em")

	EmbeddedRouter.Post("/login", controller.EmbeddedLogin2)
	// EmbeddedRouter.Use(security.EmbeddedMiddleware(store[0]))
}
