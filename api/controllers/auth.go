package controllers

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
	"github.com/Sea-of-Keys/seaofkeys-api/pkg"
)

type AuthController struct {
	repo    repos.AuthRepoInterface
	store   *session.Store
	Logging pkg.CustomLogginAndErrorInterface
}

func (con *AuthController) Login(c *fiber.Ctx) error {
	var user models.Login
	sess, err := con.store.Get(c)
	if err != nil {
		con.Logging.Log("session", err)
		return fiber.NewError(fiber.StatusNoContent, "failed to start a session")
	}
	if err := c.BodyParser(&user); err != nil {
		con.Logging.Log("AC100", err)
		return fiber.NewError(fiber.StatusNoContent, "not the right body")
	}

	data, err := con.repo.PostLogin(user)
	if err != nil {
		con.Logging.Log("AC101", err)
		return fiber.NewError(fiber.StatusNoContent, "AC101: email or password did not match")
	}
	tokenString, err := security.NewToken(data.ID, *data.Email)
	if err != nil {
		con.Logging.Log("AC102", err)
		return fiber.NewError(fiber.StatusInternalServerError, "AC101: failed to create token")
	}
	sess.Set("ActiveToken", tokenString)
	sess.Save()
	return c.JSON(&fiber.Map{
		"token": tokenString,
		"user":  data,
	})
}

func (con *AuthController) Logout(c *fiber.Ctx) error {
	sess, err := con.store.Get(c)
	if err != nil {
		con.Logging.Log("session", err)
		return fiber.NewError(fiber.StatusNoContent, "failed to find the session")
	}
	sess.Destroy()

	return c.JSON(&fiber.Map{
		"logout": true,
	})
}
func (con *AuthController) RefreshToken(c *fiber.Ctx) error {

	sess, err := con.store.Get(c)
	if err != nil {
		con.Logging.Log("CC104", err)
		return fiber.NewError(fiber.StatusInternalServerError, "AC104 failed to find your session")
	}
	sToken := sess.Get("ActiveToken")
	tokenStr := sToken.(string)
	id, email, err := security.GetTokenData(tokenStr, os.Getenv("PSCRERT"))
	if err != nil {
		con.Logging.Log("AC105", err)
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"AC105: failed to get token data",
		)
	}
	newToken, err := con.repo.CheckTokenData(id, email)
	if err != nil {
		con.Logging.Log("AC106", err)
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"AC106: failed to create a new token",
		)
	}
	sess.Set("ActiveToken", newToken)
	sess.Save()
	return c.SendStatus(200)
}

func NewAuthController(repo repos.AuthRepoInterface, store *session.Store, log pkg.CustomLogginAndErrorInterface) AuthInterfaceMethods {
	return &AuthController{repo, store, log}
}
func RegisterAuthController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewAuthRepo(reg.Db)
	log := pkg.NewCustomLogginAndError()
	controller := NewAuthController(repo, reg.Store, log)

	AuthRouter := reg.Router.Group("/auth")
	AuthRouter.Post("/login", controller.Login)
	AuthRouter.Use(security.TokenMiddleware(reg.Store))
	AuthRouter.Get("/logout", controller.Logout)
	AuthRouter.Get("/", controller.RefreshToken)
}
