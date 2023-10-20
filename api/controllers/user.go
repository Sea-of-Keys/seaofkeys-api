package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type UserController struct {
	repo  *repos.UserRepo
	store *session.Store
}

func (con *UserController) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	UID := uint(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C10: "+err.Error())
	}
	data, err := con.repo.GetUser(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C11: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"user": data,
	})
}
func (con *UserController) Gets(c *fiber.Ctx) error {
	data, err := con.repo.GetUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C12: "+err.Error())
	}
	return c.JSON(data)
}
func (con *UserController) Post(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(&fiber.Map{
			"user": user,
		})
		// return fiber.NewError(fiber.StatusInternalServerError, "C13: "+err.Error())
	}
	data, err := con.repo.PostUser(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C14: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"user": data,
	})
}
func (con *UserController) Posts(c *fiber.Ctx) error {
	var user []models.User

	fmt.Println(user)
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(&fiber.Map{
			"user": user,
			// return fiber.NewError(fiber.StatusInternalServerError, "C15: "+err.Error())
		})
	}
	fmt.Println(user)
	data, err := con.repo.PostUsers(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C16: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"user": data,
	})
}
func (con *UserController) Put(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(&fiber.Map{
			"user": user,
		})
		// return fiber.NewError(fiber.StatusInternalServerError, "C17: "+err.Error())
	}
	data, err := con.repo.PutUser(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C18: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"user": data,
	})
}
func (con *UserController) Del(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C19: "+err.Error())
	}
	data, err := con.repo.DelUser(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"sus": data,
	})
}
func (con *UserController) Dels(c *fiber.Ctx) error {
	var users []models.Delete
	if err := c.BodyParser(&users); err != nil {
		return c.JSON(users)
	}
	data, err := con.repo.DelUsers(users)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(data)
}
func (con *UserController) GetTeamsUserIsNotOn(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.GetAllTeamsUserIsNotOn(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"teams": data,
	})
}
func NewUsercontroller(repo *repos.UserRepo, store *session.Store) models.UserInterfaceMethods {
	return &UserController{repo, store}
}

func RegisterUserController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewUserRepo(reg.Db)
	controller := NewUsercontroller(repo, reg.Store)

	UserRouter := reg.Router.Group("/user")

	UserRouter.Use(security.TokenMiddleware(reg.Store))
	UserRouter.Get("/:id", controller.Get)
	UserRouter.Get("/", controller.Gets)
	UserRouter.Post("/", controller.Post)
	UserRouter.Post("/more", controller.Posts)
	UserRouter.Put("/", controller.Put)
	UserRouter.Delete("/del/:id", controller.Del)
	UserRouter.Delete("/del", controller.Dels)
	UserRouter.Get("/teams/:id", controller.GetTeamsUserIsNotOn)
}
