package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

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
	return nil
}
func (con *HistoryController) PostHistory(c *fiber.Ctx) error {
	return nil
}
func (con *HistoryController) PutHistory(c *fiber.Ctx) error {
	return nil
}
func (con *HistoryController) DelHistory(c *fiber.Ctx) error {
	return nil
}

func NewHistoryController(repo *repos.HistoryRepo) *HistoryController {
	return &HistoryController{repo}
}

func RegisterHistoryController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewHistoryRepo(db)
	controller := NewHistoryController(repo)
	fmt.Print(controller)
}
