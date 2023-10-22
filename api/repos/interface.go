package repos

import "github.com/Sea-of-Keys/seaofkeys-api/api/models"

type AuthRepoInterface interface {
	PostLogin(models.Login) (*models.User, error)
	PutPassword(uint, string) (*models.User, error)
	CheckTokenData(uint, string) (string, error)
}

type WebRepoInterface interface {
	GetCheckToken(string) (*models.UserPC, error)
}
type UserRepoInterface interface {
	GetUser(uint) (*models.User, error)
	GetUsers() ([]models.User, error)
	PostUser(models.User) (*models.User, error)
	PostUsers([]models.User) ([]models.User, error)
	PutUser(models.User) (*models.User, error)
	DelUser(uint) (bool, error)
	DelUsers([]models.Delete) (bool, error)
	GetAllTeamsUserIsNotOn(uint) ([]models.Team, error)
	PutPassword(string, string, ...string) (bool, error)
}
type TeamRepoInterface interface {
	GetTeam(uint) (*models.Team, error)
	GetTeams() ([]models.Team, error)
	PostTeam(models.Team) (*models.Team, error)
	PutTeam(models.Team) (*models.Team, error)
	DelTeam(uint) (bool, error)
	DelTeams([]models.Delete) (bool, error)
	GetAllUserNotOnTheTeam(uint) ([]models.User, error)
	AddToTeam(models.TeamUsers) (*models.Team, error)
	AddUsersToTeam(uint, uint) (*models.Team, error)
	AddTeamsToUser(models.UserTeams) (*models.User, error)
	RemoveFromTeam(uint, uint) (*models.Team, error)
	RemoveUsersFromTeam(models.TeamUsers) (*models.Team, error)
	RemoveTeamsFromUser(models.UserTeams) (*models.User, error)
}
type HistoryRepoInterface interface {
	GetHistory(uint) (*models.History, error)
	GetHistorys() ([]models.History, error)
	PostHistory(models.History) (*models.History, error)
	PutHistory(models.History) (*models.History, error)
	DelHistory(uint) (bool, error)
}
type EmbeddedRepoInterface interface {
	GetEmbedded(uint) (*models.Embedded, error)
	GetEmbeddeds() ([]models.Embedded, error)
	PostEmbedded(models.Embedded) (*models.Embedded, error)
	PutEmbedded(models.Embedded) (*models.Embedded, error)
	DelEmbedded(uint) (bool, error)
	GetSetup(uint) error
	PostSetup() error
	UpdateSecrect(string, string) (bool, error)
	PostEmbeddedSetup(models.EmbedSetup) (bool, error)
	PostCodeLogin(string, string, uint) (bool, error)
	PostHistoryLogin(models.History) (bool, error)
}
type PermissionRepoInterface interface {
	GetPermission(uint) (*models.Permission, error)
	GetPermissions() ([]models.Permission, error)
	PostPermission(models.Permission) (*models.Permission, error)
	PutPermission(models.Permission) (*models.Permission, error)
	DelPermission(uint) (bool, error)
	DelPermissions([]models.Delete) (bool, error)
	GetUsersPermissions(uint) ([]models.Permission, error)
	GetTeamsPermissions(uint) ([]models.Permission, error)
	CleanPermission() error
}
type StatsRepoInterface interface {
	GetUsersCount() (int, error)
	GetTeamsCount() (int, error)
	GetRoomsCount() (int, error)
	GetLoginsCount() (int, error)
}
type RoomRepoInterface interface {
	GetRoom(uint) (*models.Room, error)
	GetRooms() ([]models.Room, error)
	PostRoom(models.Room) (*models.Room, error)
	PostRooms([]models.Room) ([]models.Room, error)
	PutRoom(models.Room) (*models.Room, error)
	DelRoom(uint) (bool, error)
	DelRooms([]models.Delete) (bool, error)
}
