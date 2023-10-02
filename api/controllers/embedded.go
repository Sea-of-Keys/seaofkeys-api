package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type EmbeddedController struct {
	repo *repos.EmbeddedRepo
}

// Just for testing
type Login struct {
	ID     uint   `json:"id"`
	Code   string `json:"code"`
	RoomID uint   `json:"room_id"`
}

func (con *EmbeddedController) EmbeededLogin(c *fiber.Ctx) error {
	var login Login
	if err := c.BodyParser(&login); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	sus, err := con.repo.PostCode(login.Code, login.ID, login.RoomID)
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

func NewEmbeddedController(repo *repos.EmbeddedRepo) *EmbeddedController {
	return &EmbeddedController{repo}
}

func RegisterEmbeddedController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewEmbeddedRepo(db)
	controller := NewEmbeddedController(repo)

	EmbeddedRouter := router.Group("/em")

	EmbeddedRouter.Post("/login", controller.EmbeededLogin)
}
