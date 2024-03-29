package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type RoomRepo struct {
	db *gorm.DB
}

// Don't no if it needs to be *models.Room or models.Room
func (r *RoomRepo) GetRoom(id uint) (*models.Room, error) {
	var room models.Room
	if err := r.db.Debug().First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
func (r *RoomRepo) GetRooms() ([]models.Room, error) {
	var rooms []models.Room
	if err := r.db.Debug().Find(&rooms).Error; err != nil {
		return nil, err
	}
	return rooms, nil
}
func (r *RoomRepo) PostRoom(room models.Room) (*models.Room, error) {
	if err := r.db.Debug().Create(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
func (r *RoomRepo) PostRooms(room []models.Room) ([]models.Room, error) {
	if err := r.db.Debug().Create(&room).Error; err != nil {
		return nil, err
	}
	return room, nil
}
func (r *RoomRepo) PutRoom(room models.Room) (*models.Room, error) {
	if err := r.db.Debug().Model(&room).Updates(&room).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
func (r *RoomRepo) DelRoom(id uint) (bool, error) {
	var room models.Room
	if err := r.db.Debug().Debug().Delete(&room, id).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *RoomRepo) DelRooms(id []models.Delete) (bool, error) {
	var room models.Room
	var embedded models.Embedded
	var permission models.Permission

	for _, v := range id {
		room.ID = v.ID
		if err := r.db.Debug().Model(&permission).Where("room_id = ?", room.ID).Updates(map[string]interface{}{
			"RoomID": nil,
		}).Error; err != nil {
			return false, err
		}
		if err := r.db.Debug().Model(&embedded).Where("room_id = ?", room.ID).Updates(map[string]interface{}{
			"RoomID": nil,
		}).Error; err != nil {
			return false, err
		}
		if err := r.db.Debug().Debug().Delete(&room, v.ID).Error; err != nil {
			return false, err
		}

	}
	return true, nil
}

func NewRoomRepo(db *gorm.DB) *RoomRepo {
	return &RoomRepo{db}
}
