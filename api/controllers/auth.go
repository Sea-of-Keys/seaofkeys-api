package controllers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type AuthController struct {
	repo  *repos.AuthRepo
	store *session.Store
}

var jwtKey = []byte("my_secret_key")

func (con *AuthController) Login(c *fiber.Ctx) error {
	// var token models.Token
	var user models.Login
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}
	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	fmt.Printf("LoginSess: %v\n", sess)
	fmt.Printf("LoginSess: %v\n", sess)
	fmt.Printf("LoginSess: %v\n", sess)
	data, err := con.repo.PostLogin(user)
	if err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}
	tokenString, err := security.NewToken(data.ID, *data.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "C003: "+err.Error())
	}
	sess.Set("ActiveToken", tokenString)
	sess.Save()
	// fmt.Println(tokenString)
	// fmt.Printf("sess: %v\n token: %v\n", sess, tokenString)
	c.Set("Authorization", "Bearer "+tokenString)
	// c.Set("Authorization", "Bearer "+tokenString)

	return c.JSON(&fiber.Map{
		"token": tokenString,
		"user":  data,
	})
}

// func (con *AuthController) LoginOrginal(c *fiber.Ctx) error {
// 	var user models.Login
// 	// var err error
// 	if err := c.BodyParser(&user); err != nil {
// 		return fiber.NewError(fiber.StatusNoContent, err.Error())
// 	}
// 	sess, err := con.store.Get(c)
// 	if err != nil {
// 		panic(err)
// 	}
// 	data, err := con.repo.PostLogin(user)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusNoContent, err.Error())
// 	}
// 	expirationTime := time.Now().Add(5 * time.Minute)
// 	claims := &Claims{
// 		Email: *data.Email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			// In JWT, the expiry time is expressed as unix milliseconds
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte(os.Getenv("SCRERT")))
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}
// 	sess.Set("ActiveToken", tokenString)
// 	sess.Save()
// 	c.Set("Authorization", "Bearer "+tokenString)

//		return c.JSON(&fiber.Map{
//			"token": tokenString,
//			"user":  data,
//		})
//	}
func (con *AuthController) Logout(c *fiber.Ctx) error {
	sess, err := con.store.Get(c)
	if err != nil {
		panic(err)
	}
	// sess.Set("ActiveToken", nil)
	sess.Delete("ActiveToken")
	sess.Save()

	return c.JSON(&fiber.Map{
		"logout": true,
	})
}
func (con *AuthController) RefreshToken(c *fiber.Ctx) error {
	// var token models.Token
	// var err
	sess, err := con.store.Get(c)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	sToken := sess.Get("ActiveToken")
	tokenStr := sToken.(string)
	// secretKey := os.Getenv("PSCRERT")
	id, email, err := security.RefreshToken(tokenStr, os.Getenv("PSCRERT"))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	newToken, err := con.repo.CheckTokenData(id, email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// NewToken, err := securhty.NewToken(newTokenData.ID, newTokenData.Email)
	// if err != nil {
	// 	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	// }
	sess.Set("ActiveToken", newToken)
	sess.Save()
	return c.SendStatus(200)
}

// ############# Func to show your password one time ###############
func (con *AuthController) Code(c *fiber.Ctx) error {
	data := fiber.Map{
		"user": "Joe",
	}
	return c.Render("auth/index", data)
}

// ########### Change Password ##############
func (con *AuthController) RestCode(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

// ############ Maby Make it so you get a email to set ure password #############
func (con *AuthController) SetPassword(c *fiber.Ctx) error {
	data := fiber.Map{}
	return c.Render("home/index", data)
}
func (con *AuthController) Hello(c *fiber.Ctx) error {
	fmt.Println("Kronborg er gud")
	return nil
}

func NewAuthController(repo *repos.AuthRepo, store *session.Store) *AuthController {
	return &AuthController{repo, store}
}
func RegisterAuthController(reg models.RegisterController, store ...*session.Store) {
	repo := repos.NewAuthRepo(reg.Db)
	controller := NewAuthController(repo, reg.Store)

	AuthRouter := reg.Router.Group("/auth")
	// AuthRouter.Static("/static", "./static")
	AuthRouter.Post("/login", controller.Login)
	AuthRouter.Get("/test", controller.Code)
	// AuthRouter.Get("/", controller.SetPassword)
	AuthRouter.Get("/hello", controller.Hello)
	AuthRouter.Use(security.TokenMiddleware(reg.Store))
	AuthRouter.Get("/logout", controller.Logout)
	AuthRouter.Get("/", controller.RefreshToken)
}
