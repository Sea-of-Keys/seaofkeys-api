package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type StatsController struct {
	repo *repos.StatsRepo
}

func (con *StatsController) GetUsersCount(c *fiber.Ctx) error {
	data, err := con.repo.GetUsersCount()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(&fiber.Map{
		"user_count": data,
	})
}
func (con *StatsController) GetTeamsCount(c *fiber.Ctx) error {
	data, err := con.repo.GetTeamsCount()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(&fiber.Map{
		"user_count": data,
	})
}
func (con *StatsController) GetRoomsCount(c *fiber.Ctx) error {
	data, err := con.repo.GetRoomsCount()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(&fiber.Map{
		"user_count": data,
	})
}
func (con *StatsController) GetLoginsCount(c *fiber.Ctx) error {
	data, err := con.repo.GetLoginsCount()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(&fiber.Map{
		"user_count": data,
	})
}
func NewStatsController(repo *repos.StatsRepo) *StatsController {
	return &StatsController{repo}
}

func RegisterStatsController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewStatsRepo(db)
	controller := NewStatsController(repo)

	StatsRouter := router.Group("/stats")

	StatsRouter.Get("/users", controller.GetUsersCount)
	StatsRouter.Get("/teams", controller.GetTeamsCount)
	StatsRouter.Get("/rooms", controller.GetRoomsCount)
	StatsRouter.Get("/logins", controller.GetLoginsCount)
}
