package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type HistoryController struct {
	repo  *repos.HistoryRepo
	store *session.Store
}

func (con *HistoryController) Get(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	UID := uint(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.GetHistory(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) Gets(c *fiber.Ctx) error {
	data, err := con.repo.GetHistorys()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) Post(c *fiber.Ctx) error {
	var history models.History
	if err := c.BodyParser(&history); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.PostHistory(history)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) Put(c *fiber.Ctx) error {
	var history models.History
	if err := c.BodyParser(&history); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.PutHistory(history)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) Del(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	UID := uint(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	data, err := con.repo.DelHistory(UID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C20: "+err.Error())
	}
	return c.JSON(&fiber.Map{
		"history": data,
	})
}
func (con *HistoryController) Posts(c *fiber.Ctx) error {
	return nil
}
func (con *HistoryController) Dels(c *fiber.Ctx) error {
	return nil
}

//func (con *HistoryController) TestOne(c *fiber.Ctx) error {

//	sess, err := con.store.Get(c)
//	if err != nil {
//		panic(err)
//	}
//	// Check Token HERE
//	// if token is legiget set token in session

//	//this is just to test if i get a token
//	// Read and output the session variable
//	name := sess.Get("token")
//	fmt.Printf("Name from session: %v\n", name)

//	return c.JSON(&fiber.Map{
//		"name": name,
//	})
//}

func NewHistoryController(
	repo *repos.HistoryRepo,
	store *session.Store,
) models.HistoryInterfaceMethods {
	return &HistoryController{repo, store}
}

func RegisterHistoryController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewHistoryRepo(reg.Db)
	controller := NewHistoryController(repo, reg.Store)

	HistoryRouter := reg.Router.Group("/history")

	HistoryRouter.Use(security.TokenMiddleware(reg.Store))
	// HistoryRouter.Get("/test", controller.Gets)
	HistoryRouter.Get("/:id", controller.Get)
	HistoryRouter.Get("/", controller.Gets)
	HistoryRouter.Post("/", controller.Post)
	HistoryRouter.Put("/", controller.Put)
	HistoryRouter.Delete("/:id", controller.Del)
}
