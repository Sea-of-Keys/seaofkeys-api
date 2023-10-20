package repos

import (
	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type HistoryRepo struct {
	db *gorm.DB
}

func (r *HistoryRepo) GetHistory(id uint) (*models.History, error) {
	var history models.History
	if err := r.db.Debug().F.Preload("User").Preload("Permission").First(&history, id).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (r *HistoryRepo) GetHistorys() ([]models.History, error) {
	var history []models.History
	if err := r.db.Debug().Preload("User").Preload("Permission").Find(&history).Error; err != nil {
		return nil, err
	}
	return history, nil
}
func (r *HistoryRepo) PostHistory(history models.History) (*models.History, error) {
	if err := r.db.Debug().Create(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (r *HistoryRepo) PutHistory(history models.History) (*models.History, error) {
	if err := r.db.Debug().Model(&history).Updates(&history).Error; err != nil {
		return nil, err
	}
	return &history, nil
}
func (r *HistoryRepo) DelHistory(id uint) (bool, error) {
	var history models.History
	if err := r.db.Debug().Delete(&history, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func NewHistoryRepo(db *gorm.DB) *HistoryRepo {
	return &HistoryRepo{db}
}
