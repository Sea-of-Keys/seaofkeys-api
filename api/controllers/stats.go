package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type StatsController struct {
	repo  *repos.StatsRepo
	store *session.Store
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
func NewStatsController(repo *repos.StatsRepo, store *session.Store) *StatsController {
	return &StatsController{repo, store}
}

func RegisterStatsController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewStatsRepo(reg.Db)
	controller := NewStatsController(repo, reg.Store)

	StatsRouter := reg.Router.Group("/stats")

	StatsRouter.Use(security.TokenMiddleware(reg.Store))
	StatsRouter.Get("/users", controller.GetUsersCount)
	StatsRouter.Get("/teams", controller.GetTeamsCount)
	StatsRouter.Get("/rooms", controller.GetRoomsCount)
	StatsRouter.Get("/logins", controller.GetLoginsCount)
}
