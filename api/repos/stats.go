package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type StatsRepo struct {
	db *gorm.DB
}

func (r *StatsRepo) GetUsersCount() (int, error) {
	var user models.User
	var count int64
	if err := r.db.Debug().Model(&user).Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}
func (r *StatsRepo) GetTeamsCount() (int, error) {
	var team models.Team
	var count int64
	if err := r.db.Debug().Model(&team).Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}
func (r *StatsRepo) GetRoomsCount() (int, error) {
	var room models.Room
	var count int64
	if err := r.db.Debug().Model(&room).Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}
func (r *StatsRepo) GetLoginsCount() (int, error) {
	var login models.User
	var count int64
	if err := r.db.Debug().Model(&login).Count(&count).Error; err != nil {
		return 0, nil
	}
	return int(count), nil
}

func NewStatsRepo(db *gorm.DB) *StatsRepo {
	return &StatsRepo{db}
}
