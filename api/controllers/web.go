package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type WebController struct {
	repo *repos.WebRepo
}

func (con *WebController) GetPage(c *fiber.Ctx) error {
	data := fiber.Map{
		"password": true,
	}
	return c.Render("web/index", data)
}

func NewWebController(repo *repos.WebRepo) *WebController {
	return &WebController{repo}
}

func RegisterWebController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewWebRepo(db)
	controller := NewWebController(repo)

	WebRouter := router.Group("/web")

	WebRouter.Get("/", controller.GetPage)
}
