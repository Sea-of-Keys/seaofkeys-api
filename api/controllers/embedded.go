package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
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
	return nil
}
func (con *EmbeddedController) EmbeededLogin(c *fiber.Ctx) error {
	var login Login
	if err := c.BodyParser(&login); err != nil {
		gg := errors.New("E22: " + err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, gg.Error())
	}
	sus, err := con.repo.PostCodeV2(login.Code, login.RoomID, login.UserID)
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

func NewEmbeddedController(repo *repos.EmbeddedRepo, store *session.Store) *EmbeddedController {
	return &EmbeddedController{repo, store}
}

func RegisterEmbeddedController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewEmbeddedRepo(reg.Db)
	controller := NewEmbeddedController(repo, store[0])

	EmbeddedRouter := reg.Router.Group("/em")

	EmbeddedRouter.Post("/login", controller.EmbeededLogin)
	// EmbeddedRouter.Use(security.EmbeddedMiddleware(store[0]))
}
