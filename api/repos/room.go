package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type RoomRepo struct {
	db *gorm.DB
}

// Don't no if it needs to be *models.Room or models.Room
func (r *RoomRepo) GetRoom(id string) (models.Room, error) {
	return models.Room{}, nil
}
func (r *RoomRepo) GetRooms() ([]models.Room, error) {
	return nil, nil
}
func (r *RoomRepo) PostRoom(room models.Room) (*models.Room, error) {
	return nil, nil
}
func (r *RoomRepo) PostRooms(room models.Room) ([]models.Room, error) {
	return nil, nil
}
func (r *RoomRepo) PutRoom(room models.Room) (*models.Room, error) {
	return nil, nil
}
func (r *RoomRepo) DelRoom(id uint) (bool, error) {
	return true, nil
}
