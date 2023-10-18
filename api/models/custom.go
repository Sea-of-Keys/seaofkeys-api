package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Delete struct {
	ID uint `json:"id"`
}
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type EmbeddedLogin struct {
	ID     uint   `json:"id"`
	Code   string `json:"code"`
	RoomID uint   `json:"room_id"`
	UserID uint   `json:"user_id"`
}

type TeamUsers struct {
	TeamID uint   `json:"team_id"`
	UserID []uint `json:"users"`
}
type UserTeams struct {
	UserID uint   `json:"user_id"`
	TeamID []uint `json:"teams"`
}

type AddToTeam struct {
	UserID uint `json:"user_id"`
	TeamID uint `json:"team_id"`
}

type SetPasswordAndCode struct {
	PasswordOne string `json:"password_one"`
	PasswordTwo string `json:"password_two"`
	CodeOne     string `json:"code_one"`
	CodeTwo     string `json:"code_two"`
}
type EmbedSetup struct {
	EmbeddedID uint   `json:"embedded_id"`
	Ssshhh     string `json:"ssshhh"`
}
type Te interface {
}
type RegisterController struct {
	Db     *gorm.DB
	Router fiber.Router
	Store  *session.Store
}

type Token struct {
	ID    uint
	Email string
	Token string
}
type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}
