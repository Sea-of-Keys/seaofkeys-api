package models

import "github.com/gofiber/fiber/v2"

// type Model interface {
// }
type RoomInterfaceMethods interface {
	Posts(*fiber.Ctx) error

	CRUDInterface
}

type TeamInterfaceMethods interface {
	PostAddToTeam(*fiber.Ctx) error
	DeleteUsersRemoveFromTeam(*fiber.Ctx) error
	GetAllUserNotOnTheTeam(*fiber.Ctx) error
	RemoveTeamsFromUser(*fiber.Ctx) error
	AddTeamsToUser(*fiber.Ctx) error

	CRUDInterface
}
type CRUDInterface interface {
	Get(*fiber.Ctx) error
	Gets(*fiber.Ctx) error
	Post(*fiber.Ctx) error
	Put(*fiber.Ctx) error
	Del(*fiber.Ctx) error
	Dels(*fiber.Ctx) error
	Posts(*fiber.Ctx) error
}
type UserInterfaceMethods interface {
	GetTeamsUserIsNotOn(*fiber.Ctx) error

	CRUDInterface
}
type WebInterfaceMethods interface {
	GetPage(*fiber.Ctx) error
	PostNewCodes(*fiber.Ctx) error
	TestOne(*fiber.Ctx) error
}
type AuthInterfaceMethods interface {
	Login(*fiber.Ctx) error
	Logout(*fiber.Ctx) error
	RefreshToken(*fiber.Ctx) error
}
type HistoryInterfaceMethods interface {
	CRUDInterface
}

type StatsInterfaceMethods interface {
	GetUsersCount(*fiber.Ctx) error
	GetTeamsCount(*fiber.Ctx) error
	GetRoomsCount(*fiber.Ctx) error
	GetLoginsCount(*fiber.Ctx) error
}
type PermissionInterfaceMethods interface {
	GetFindUsersPermissions(*fiber.Ctx) error
	GetFindTeamsPermissions(*fiber.Ctx) error
	CRUDInterface
}
type EmbeddedInterfaceMethods interface {
	Setup(*fiber.Ctx) error
	Login(*fiber.Ctx) error
	EmbeddedLoginLive(*fiber.Ctx) error

	// CRUDController
}
