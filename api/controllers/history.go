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
	return nil
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
