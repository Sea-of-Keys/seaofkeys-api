package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type TeamController struct {
	repo *repos.TeamRepo
}

type AddToTeam struct {
	UserID uint `json:"user_id"`
	TeamID uint `json:"team_id"`
}

func (con *TeamController) DelTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	UID := uint(id)
	data, err := con.repo.DelTeam(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"sus": data,
	})
}

func (con *TeamController) PostAddToTeam(c *fiber.Ctx) error {
	var team AddToTeam
	if err := c.BodyParser(&team); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.AddToTeam(team.TeamID, team.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C21: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func (con *TeamController) PostRemoveFromTeam(c *fiber.Ctx) error {
	var team AddToTeam
	if err := c.BodyParser(&team); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C22: "+err.Error())
	}
	data, err := con.repo.RemoveFromTeam(team.TeamID, team.UserID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func NewTeamController(repo *repos.TeamRepo) *TeamController {
	return &TeamController{repo}
}

func RegisterTeamController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewTeamRepo(db)
	controller := NewTeamController(repo)
	TeamRouter := router.Group("/team")

	TeamRouter.Post("/add", controller.PostAddToTeam)
	TeamRouter.Post("/remove", controller.PostRemoveFromTeam)
	TeamRouter.Delete("/:id", controller.DelTeam)
}
