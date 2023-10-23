package controllers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type EmbeddedController struct {
	repo  repos.EmbeddedRepoInterface
	store *session.Store
}

func (con *EmbeddedController) Setup(c *fiber.Ctx) error {
	var emb models.EmbedSetup

	sess, err := con.store.Get(c)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to start a new session")
	}
	if err := c.BodyParser(&emb); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to proces the body")
	}
	if ok, err := con.repo.PostEmbeddedSetup(emb); !ok || err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "error in the setup")
	}

	base64, err := security.NewBase64Token()
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "error when getting a new token")
	}
	token, err := security.NewEmbeddedToken(base64)
	sess.Set("EmbeddedSession", token)
	if err := sess.Save(); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to save the token")
	}
	if ok, err := con.repo.UpdateSecrect(emb.Ssshhh, token); ok && err == nil {
		return c.JSON(&fiber.Map{
			"session_is": true,
		})
	}
	log.Println(err)
	return fiber.NewError(fiber.StatusInternalServerError, "failed to make a setup")

}

func (con *EmbeddedController) Login(c *fiber.Ctx) error {
	var login models.EmbeddedLogin
	sess, err := con.store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if sess.Get("EmbeddedSession") == nil {
		return fiber.NewError(fiber.StatusInternalServerError, "E21: "+err.Error())
	}
	if err := c.BodyParser(&login); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "E22: "+err.Error())
	}
	fmt.Println(login.Code)
	result := strings.Split(login.Code, "#")
	sus, err := con.repo.PostCodeLogin(result[0], result[1], login.RoomID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "E23: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"message": "IT IS A LIVE",
		"data":    sus,
	})
}
func (con *EmbeddedController) Refresh(c *fiber.Ctx) error {
	sess, err := con.store.Get(c)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to find session")
	}
	if sess.Get("EmbeddedSession") == nil {

		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "Session token not set")
	}
	base64, err := security.NewBase64Token()
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to make a new base64 token")
	}
	newtoken, err := security.NewEmbeddedToken(base64)
	if err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to make a new jwt token")
	}
	oldtoken := sess.Get("EmbeddedSession").(string)
	sess.Set("EmbeddedSession", newtoken)
	if err := sess.Save(); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to save token in session")
	}
	if ok, err := con.repo.UpdateSecrect(oldtoken, newtoken); ok && err == nil {
		return nil
	}
	log.Println(err)
	return fiber.NewError(fiber.StatusInternalServerError, "failed in makeing a refresh off token")
}
func NewEmbeddedController(
	repo repos.EmbeddedRepoInterface,
	store *session.Store,
) EmbeddedInterfaceMethods {
	return &EmbeddedController{repo, store}
}

func RegisterEmbeddedController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewEmbeddedRepo(reg.Db)
	controller := NewEmbeddedController(repo, store[0])

	EmbeddedRouter := reg.Router.Group("/em")

	EmbeddedRouter.Post("/setup", controller.Setup)
	EmbeddedRouter.Use(security.TokenEmbeddedMiddleware(store[0]))
	EmbeddedRouter.Get("/refresh", controller.Refresh)
	EmbeddedRouter.Post("/login", controller.Login)

}
