package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type AuthController struct {
	repo *repos.AuthRepo
}

func (con *AuthController) Login(c *fiber.Ctx) error {
	var user models.User
	// var err error
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}
	data, err := con.repo.PostLogin(user)
	if err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}
	// fmt.Println(data)
	return c.JSON(data)
}

func NewAuthController(repo *repos.AuthRepo) *AuthController {
	return &AuthController{repo}
}
func RegisterAuthController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewAuthRepo(db)
	controller := NewAuthController(repo)

	AuthRouter := router.Group("/auth")

	AuthRouter.Post("/login", controller.Login)
}
