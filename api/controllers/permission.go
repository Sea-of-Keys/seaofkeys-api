package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
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
func (con *PermissionController) Get(c *fiber.Ctx) error {
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
func (con *PermissionController) Gets(c *fiber.Ctx) error {
	data, err := con.repo.GetPermissions()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C801: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"permissions": data,
	})
}
func (con *PermissionController) Post(c *fiber.Ctx) error {
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
func (con *PermissionController) Posts(c *fiber.Ctx) error {
	return nil
}
func (con *PermissionController) Put(c *fiber.Ctx) error {
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
func (con *PermissionController) Del(c *fiber.Ctx) error {
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
func (con *PermissionController) Dels(c *fiber.Ctx) error {
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
) PermissionInterfaceMethods {
	return &PermissionController{repo, store}
}

func RegisterPermissionController(reg models.RegisterController) {
	repo := repos.NewPermissionRepo(reg.Db)
	controller := NewPermissionController(repo, reg.Store)

	PermissionRouter := reg.Router.Group("/permission")
	PermissionRouter.Use(security.TokenMiddleware(reg.Store))
	PermissionRouter.Get("/:id", controller.Get)
	PermissionRouter.Get("/", controller.Gets)
	PermissionRouter.Post("/", controller.Post)
	PermissionRouter.Put("/", controller.Put)
	PermissionRouter.Delete("/del/:id", controller.Del)
	PermissionRouter.Delete("/del", controller.Dels)
	PermissionRouter.Get("/user/:id", controller.GetFindUsersPermissions)
	PermissionRouter.Get("/team/:id", controller.GetFindTeamsPermissions)
}
