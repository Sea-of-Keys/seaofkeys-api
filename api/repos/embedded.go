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
func (r *EmbeddedRepo) PostCode(code string, ID, RoomID uint) (bool, error) {
	var per models.Permission
	var user models.User
	if err := r.db.Debug().Preload("User").Preload("Team.Users").Find(&per, ID, RoomID).Error; err != nil {
		return false, err
	}
	if per.Team != nil {
		// var team models.Team
		for _, v := range per.Team.Users {
			UserCode := *user.Code
			if middleware.CheckPasswordHash(code, UserCode) {
				fmt.Println(v.Email)
				fmt.Println("Coden Passer")
				return true, errors.New("det virker")
			}
		}
		if err := r.db.Debug().Find(&user, &per.UserID).Error; err != nil {
			return false, errors.New("LORT PAA LORT")
		}
		UserCode := *user.Code
		if middleware.CheckPasswordHash(code, UserCode) {
			return true, nil
			// ret
		}

		// check team frist
		// if err r.db.Debug.
		return false, nil
	}
	return false, nil
}
func (r *EmbeddedRepo) PostCodeV2(code string, RoomID uint) (bool, error) {
	var pem []models.Permission
	// var user []models.User
	if err := r.db.Debug().Preload("User").Preload("Team.Users").Where("room_id = ?", RoomID).Find(&pem).Error; err != nil {
		return false, err
	}
	for _, v := range pem {
		if v.Team != nil {

			for _, g := range v.Team.Users {
				UserCode := *g.Code
				if middleware.CheckPasswordHash(code, UserCode) {
					return true, nil
				}
			}
		}
		UserCode := *v.User.Code
		if middleware.CheckPasswordHash(code, UserCode) {
			return true, nil
		}

	}
	return false, nil

}

func NewEmbeddedRepo(db *gorm.DB) *EmbeddedRepo {
	return &EmbeddedRepo{db}
}
