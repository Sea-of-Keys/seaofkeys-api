package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type TeamController struct {
	repo *repos.TeamRepo
}

func (con *TeamController) GetTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	UID := uint(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.GetTeam(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func (con *TeamController) GetTeams(c *fiber.Ctx) error {
	data, err := con.repo.GetTeams()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func (con *TeamController) PostTeam(c *fiber.Ctx) error {
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		return c.JSON(&fiber.Map{
			"team": team,
		})
	}
	data, err := con.repo.PostTeam(team)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func (con *TeamController) PutTeam(c *fiber.Ctx) error {
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.PutTeam(team)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
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
func (con *TeamController) DelTeams(c *fiber.Ctx) error {
	var teams []models.Delete
	if err := c.BodyParser(&teams); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.DelTeams(teams)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"sus": data,
	})
}
func (con *TeamController) GetAllUserNotOnTheTeam(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.GetAllUserNotOnTheTeam(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"users": data,
	})
}

func (con *TeamController) PostAddToTeam(c *fiber.Ctx) error {
	var team models.TeamUsers
	if err := c.BodyParser(&team); err != nil {
		return c.JSON(team)
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.AddToTeam(team)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C21: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}

// func (con *TeamController) PostRemoveFromTeam(c *fiber.Ctx) error {
// 	var team models.AddToTeam
// 	if err := c.BodyParser(&team); err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "C22: "+err.Error())
// 	}
// 	data, err := con.repo.RemoveFromTeam(team.TeamID, team.UserID)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
// 	}
// 	return c.JSON(&fiber.Map{
// 		"team": data,
// 	})
// }

func (con *TeamController) DeleteUsersRemoveFromTeam(c *fiber.Ctx) error {
	var team models.TeamUsers
	if err := c.BodyParser(&team); err != nil {
		return c.JSON(team)
		// return fiber.NewError(fiber.StatusInternalServerError, "C22: "+err.Error())
	}
	data, err := con.repo.RemoveUsersFromTeam(team)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func (con *TeamController) RemoveTeamsFromUser(c *fiber.Ctx) error {
	var TSU models.UserTeams
	if err := c.BodyParser(&TSU); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
	}
	data, err := con.repo.RemoveTeamsFromUser(TSU)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"User": data,
	})
}
func (con *TeamController) AddTeamsToUser(c *fiber.Ctx) error {
	var TSU models.UserTeams
	if err := c.BodyParser(&TSU); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
	}
	data, err := con.repo.AddTeamsToUser(TSU)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C23: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"User": data,
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
	TeamRouter.Delete("/remove", controller.DeleteUsersRemoveFromTeam)
	TeamRouter.Delete("/del/:id", controller.DelTeam)
	TeamRouter.Delete("/del", controller.DelTeams)
	TeamRouter.Post("/", controller.PostTeam)
	TeamRouter.Get("/:id", controller.GetTeam)
	TeamRouter.Get("/", controller.GetTeams)
	TeamRouter.Put("/", controller.PutTeam)
	TeamRouter.Get("/users/:id", controller.GetAllUserNotOnTheTeam)
	TeamRouter.Delete("/user", controller.RemoveTeamsFromUser)
	TeamRouter.Post("/user", controller.AddTeamsToUser)
	// TeamRouter.Delete("/remove/more", controller.PostRemoveUsersFromTeam)
}
