package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type HistoryController struct {
	repo *repos.HistoryRepo
}

func (con *HistoryController) GetHistory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	UID := uint(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.GetHistory(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) GetHistorys(c *fiber.Ctx) error {
	data, err := con.repo.GetHistorys()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) PostHistory(c *fiber.Ctx) error {
	var history models.History
	if err := c.BodyParser(&history); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.PostHistory(history)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) PutHistory(c *fiber.Ctx) error {
	var history models.History
	if err := c.BodyParser(&history); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.PutHistory(history)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) DelHistory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	UID := uint(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.DelHistory(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}

func NewHistoryController(repo *repos.HistoryRepo) *HistoryController {
	return &HistoryController{repo}
}

func RegisterHistoryController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewHistoryRepo(db)
	controller := NewHistoryController(repo)

	HistoryRouter := router.Group("/history")

	HistoryRouter.Get("/", controller.GetHistorys)
	HistoryRouter.Get("/:id", controller.GetHistory)
	HistoryRouter.Post("/", controller.PostHistory)
	HistoryRouter.Put("/", controller.PutHistory)
	HistoryRouter.Delete("/:id", controller.DelHistory)
}
