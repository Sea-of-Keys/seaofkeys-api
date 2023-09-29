package repos

import "gorm.io/gorm"

type HistoryRepo struct {
	db *gorm.DB
}
