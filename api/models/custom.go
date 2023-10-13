package models

type Delete struct {
	ID uint `json:"id"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type EmbeddedLogin struct {
	ID     uint   `json:"id"`
	RoomID uint   `json:"room_id"`
	Code   string `json:"code"`
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
