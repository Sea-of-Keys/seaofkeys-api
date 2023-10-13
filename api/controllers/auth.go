package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/repos"
)

type AuthController struct {
	repo *repos.AuthRepo
}
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var jwtKey = []byte("my_secret_key")

func (con *AuthController) Login(c *fiber.Ctx) error {
	var user models.Login
	// var err error
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}
	data, err := con.repo.PostLogin(user)
	if err != nil {
		return fiber.NewError(fiber.StatusNoContent, err.Error())
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Email: *data.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SCRERT")))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	c.Set("Authorization", "Bearer "+tokenString)
	return c.JSON(&fiber.Map{
		"token": tokenString,
		"user":  data,
	})
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
	return nil
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

func NewAuthController(repo *repos.AuthRepo) *AuthController {
	return &AuthController{repo}
}
func RegisterAuthController(db *gorm.DB, router fiber.Router) {
	repo := repos.NewAuthRepo(db)
	controller := NewAuthController(repo)

	AuthRouter := router.Group("/auth")
	// AuthRouter.Static("/static", "./static")
	AuthRouter.Post("/login", controller.Login)
	AuthRouter.Get("/test", controller.Code)
	AuthRouter.Get("/", controller.SetPassword)
	AuthRouter.Get("/hello", controller.Hello)
}
