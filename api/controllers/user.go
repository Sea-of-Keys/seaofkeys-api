package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

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

func NewUsercontroller(repo *repos.UserRepo) *UserController {
	return &UserController{repo}
}

func RegisterUserController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewUserRepo(db)
	controller := NewUsercontroller(repo)

	UserRouter := router.Group("/user")

	UserRouter.Get("/:id", controller.GetUser)
}
