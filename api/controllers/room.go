package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type RoomController struct {
	repo  *repos.RoomRepo
	store *session.Store
}

func (con *RoomController) GetRoom(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C30: "+err.Error())
	}
	data, err := con.repo.GetRoom(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C31: "+err.Error())
	}
	// return c.JSON(data)
	return c.JSON(&fiber.Map{
		"room": data,
	})
}
func (con *RoomController) GetRooms(c *fiber.Ctx) error {
	data, err := con.repo.GetRooms()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C31: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"room": data,
	})
}
func (con *RoomController) PostRoom(c *fiber.Ctx) error {
	var room models.Room
	if err := c.BodyParser(&room); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C33: "+err.Error())
	}
	data, err := con.repo.PostRoom(room)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C34: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"room": data,
	})
}
func (con *RoomController) PostRooms(c *fiber.Ctx) error {
	var rooms []models.Room
	if err := c.BodyParser(&rooms); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C39: "+err.Error())
	}
	data, err := con.repo.PostRooms(rooms)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C39: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"room": data,
	})
}
func (con *RoomController) PutRoom(c *fiber.Ctx) error {
	var room models.Room
	if err := c.BodyParser(&room); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C35: "+err.Error())
	}
	data, err := con.repo.PutRoom(room)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C36: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"room": data,
	})
}
func (con *RoomController) DelRoom(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C30: "+err.Error())
	}
	data, err := con.repo.DelRoom(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C31: "+err.Error())
	}
	// return c.JSON(data)
	return c.JSON(&fiber.Map{
		"room": data,
	})
}

func (con *RoomController) DelRooms(c *fiber.Ctx) error {
	var ids []models.Delete
	if err := c.BodyParser(&ids); err != nil {
		return c.JSON(ids)
		return fiber.NewError(fiber.StatusInternalServerError, "C30: "+err.Error())
	}
	data, err := con.repo.DelRooms(ids)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C31: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"sus": data,
	})
}

func NewRommController(repo *repos.RoomRepo, store *session.Store) *RoomController {
	return &RoomController{repo, store}
}

func RegisterRoomController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewRoomRepo(reg.Db)
	controller := NewRommController(repo, reg.Store)

	RoomRuter := reg.Router.Group("/room")

	RoomRuter.Use(security.TokenMiddleware(reg.Store))
	RoomRuter.Get("/:id", controller.GetRoom)
	RoomRuter.Get("/", controller.GetRooms)
	RoomRuter.Post("/", controller.PostRoom)
	RoomRuter.Post("/many", controller.PostRooms)
	RoomRuter.Put("/", controller.PutRoom)
	RoomRuter.Delete("/:id", controller.DelRoom)
	RoomRuter.Delete("/del/many", controller.DelRooms)
	// RoomRuter.Delete("/", controller.DelRoom)
}
