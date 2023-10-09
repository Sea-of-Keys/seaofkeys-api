package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type UserController struct {
	repo *repos.UserRepo
}

func (con *UserController) GetUser(c *fiber.Ctx) error {
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
func (con *UserController) GetUsers(c *fiber.Ctx) error {
	data, err := con.repo.GetUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C12: "+err.Error())
	}
	return c.JSON(data)
}
func (con *UserController) PostUser(c *fiber.Ctx) error {
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
func (con *UserController) PostUsers(c *fiber.Ctx) error {
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
func (con *UserController) PutUser(c *fiber.Ctx) error {
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
func (con *UserController) DelUser(c *fiber.Ctx) error {
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

func NewUsercontroller(repo *repos.UserRepo) *UserController {
	return &UserController{repo}
}

func RegisterUserController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewUsercontroller(repo)

	UserRouter := router.Group("/user")

	UserRouter.Get("/:id", controller.GetUser)
	UserRouter.Get("/", controller.GetUsers)
	UserRouter.Post("/", controller.PostUser)
	UserRouter.Post("/more", controller.PostUsers)
	UserRouter.Put("/", controller.PutUser)
	UserRouter.Delete("/:id", controller.DelUser)
}
