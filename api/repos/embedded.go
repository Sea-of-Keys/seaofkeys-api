package repos

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/middleware"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
)

type EmbeddedRepo struct {
	db *gorm.DB
}

func (r *EmbeddedRepo) GetEmbedded(id uint) (*models.Embedded, error) {
	return nil, nil
}
func (r *EmbeddedRepo) GetEmbeddeds() ([]models.Embedded, error) {
	return nil, nil
}
func (r *EmbeddedRepo) PostEmbedded(embed models.Embedded) (*models.Embedded, error) {
	return nil, nil
}
func (r *EmbeddedRepo) PutEmbedded(embed models.Embedded) (*models.Embedded, error) {
	return nil, nil
}
func (r *EmbeddedRepo) DelEmbedded(id uint) (bool, error) {
	return true, nil
}

// The Embedded it self
func (r *EmbeddedRepo) GetSetup(id uint) error {
	return nil
}
func (r *EmbeddedRepo) PostSetup() error {
	return nil
}
func (r *EmbeddedRepo) PostCode(code string, ID, RoomID uint) (*models.Permission, error) {
	var per models.Permission
	if err := r.db.Debug().Preload("User").Preload("Team.Users").Find(&per, ID, RoomID).Error; err != nil {
		return nil, err
	}
	if per.Team != nil {
		// var team models.Team
		for _, v := range per.Team.Users {
			if middleware.CheckPasswordHash(code, v.Code) {
				fmt.Println(v.Email)
				fmt.Println("Coden Passer")
				return nil, errors.New("det virker")
			}
		}
		// check team frist
		// if err r.db.Debug.
		return &per, nil
	}
	return &per, nil
}

func NewEmbeddedRepo(db *gorm.DB) *EmbeddedRepo {
	return &EmbeddedRepo{db}
}
