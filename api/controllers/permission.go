package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type PermissionController struct {
	repo  *repos.PermissionRepo
	store *session.Store
}

func PostBody(c *fiber.Ctx) (*models.Permission, error) {
	var permission models.Permission
	if err := c.BodyParser(&permission); err != nil {
		return nil, errors.New("C850")
	}
	return &permission, nil
}
func (con *PermissionController) GetPermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C800: "+err.Error())
	}
	data, err := con.repo.GetPermission(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permission": data,
	})
}
func (con *PermissionController) GetPermissions(c *fiber.Ctx) error {
	data, err := con.repo.GetPermissions()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permissions": data,
	})
}
func (con *PermissionController) PostPermission(c *fiber.Ctx) error {
	var permission models.Permission
	if err := c.BodyParser(&permission); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	data, err := con.repo.PostPermission(permission)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permission": data,
	})
}
func (con *PermissionController) PutPermission(c *fiber.Ctx) error {
	var permission models.Permission
	if err := c.BodyParser(&permission); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	data, err := con.repo.PutPermission(permission)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permission": data,
	})
}
func (con *PermissionController) DelPermission(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C800: "+err.Error())
	}
	data, err := con.repo.DelPermission(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permission": data,
	})
}
func (con *PermissionController) DelPermissions(c *fiber.Ctx) error {
	var ids []models.Delete
	if err := c.BodyParser(&ids); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	data, err := con.repo.DelPermissions(ids)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"sus": data,
	})
}
func (con *PermissionController) GetFindUsersPermissions(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C800: "+err.Error())
	}
	data, err := con.repo.GetUsersPermissions(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C800: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permissions": data,
	})
}
func (con *PermissionController) GetFindTeamsPermissions(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C800: "+err.Error())
	}
	data, err := con.repo.GetTeamsPermissions(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C800: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permissions": data,
	})
}

func NewPermissionController(
	repo *repos.PermissionRepo,
	store *session.Store,
) *PermissionController {
	return &PermissionController{repo, store}
}

func RegisterPermissionController(reg models.RegisterController) {
	repo := repos.NewPermissionRepo(reg.Db)
	controller := NewPermissionController(repo, reg.Store)

	PermissionRouter := reg.Router.Group("/permission")
	PermissionRouter.Get("/:id", controller.GetPermission)
	PermissionRouter.Get("/", controller.GetPermissions)
	PermissionRouter.Post("/", controller.PostPermission)
	PermissionRouter.Put("/", controller.PutPermission)
	PermissionRouter.Delete("/del/:id", controller.DelPermission)
	PermissionRouter.Delete("/del", controller.DelPermissions)
	PermissionRouter.Get("/user/:id", controller.GetFindUsersPermissions)
	PermissionRouter.Get("/team/:id", controller.GetFindTeamsPermissions)
}
