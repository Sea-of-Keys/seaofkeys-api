package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type TeamController struct {
	repo  *repos.TeamRepo
	store *session.Store
}

func (con *TeamController) Get(c *fiber.Ctx) error {
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
func (con *TeamController) Gets(c *fiber.Ctx) error {
	data, err := con.repo.GetTeams()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"team": data,
	})
}
func (con *TeamController) Post(c *fiber.Ctx) error {
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
func (con *TeamController) Posts(c *fiber.Ctx) error {
	return nil
}
func (con *TeamController) Put(c *fiber.Ctx) error {
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

func (con *TeamController) Del(c *fiber.Ctx) error {
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
func (con *TeamController) Dels(c *fiber.Ctx) error {
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

func NewTeamController(repo *repos.TeamRepo, store *session.Store) models.TeamInterfaceMethods {
	return &TeamController{repo, store}
}

func RegisterTeamController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewTeamRepo(reg.Db)
	controller := NewTeamController(repo, reg.Store)

	TeamRouter := reg.Router.Group("/team")

	TeamRouter.Use(security.TokenMiddleware(reg.Store))
	TeamRouter.Post("/add", controller.PostAddToTeam)
	TeamRouter.Delete("/remove", controller.DeleteUsersRemoveFromTeam)
	TeamRouter.Delete("/del/:id", controller.Del)
	TeamRouter.Delete("/del", controller.Dels)
	TeamRouter.Post("/", controller.Post)
	TeamRouter.Get("/:id", controller.Get)
	TeamRouter.Get("/", controller.Gets)
	TeamRouter.Put("/", controller.Put)
	TeamRouter.Get("/users/:id", controller.GetAllUserNotOnTheTeam)
	TeamRouter.Delete("/user", controller.RemoveTeamsFromUser)
	TeamRouter.Post("/user", controller.AddTeamsToUser)
	// TeamRouter.Delete("/remove/more", controller.PostRemoveUsersFromTeam)
}
