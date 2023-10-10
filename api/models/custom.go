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
