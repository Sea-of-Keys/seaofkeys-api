package repos

import (
	"errors"
	"fmt"
	"strconv"

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
	if err := r.db.Debug().Preload("User").Preload("Team.Users").Where("room_id = ?", RoomID).Find(&per, ID).Error; err != nil {
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
func (r *EmbeddedRepo) PostCodeV2(code, UserID string, RoomID uint) (bool, error) {
	var pem []models.Permission
	// var user []models.User
	if err := r.db.Debug().Preload("User").Preload("Team.Users").Where("room_id = ?", RoomID).Find(&pem).Error; err != nil {
		return false, err
	}
	for _, v := range pem {
		if v.Team != nil {

			for _, g := range v.Team.Users {
				// UserCode := *g.Code
				if middleware.CheckPasswordHash(code, *g.Code) {
					return true, nil
				}
			}
		}
		// UserCode := *v.User.Code
		if middleware.CheckPasswordHash(code, *v.User.Code) {
			return true, nil
		}

	}
	return false, nil

}
func (r *EmbeddedRepo) PostCodeV3(code, UserID string, RoomID uint) (models.Permission, error) {
	// var user models.User
	// var Userpem models.Permission
	var pem models.Permission
	userIdInt, _ := strconv.Atoi(UserID)

	// SQL SELECT * FROM permissions AS p WHERE p.user_id = 1 AND p.room_id = 3 OR p.room_id = 3 AND p.team_id IN (SELECT team_id FROM teams_users WHERE team_id = p.team_id AND user_id = 1);

	// if err := r.db.Debug().
	// 	Joins("LEFT JOIN permissions AS p ON ? = p.user_id AND ? = p.room_id", userIdInt, RoomID).
	// 	Joins("LEFT JOIN permissions AS q ON q.room_id = ? AND q.team_id IN (SELECT user_id FROM teams_users WHERE team_id = ?) ", RoomID, userIdInt).
	// 	Preload("Team.Users").
	// 	Preload("User").
	// 	// Where("p.start_time < ? AND p.start_time > ?", time.Now(), time.Now()).
	// 	Find(&pem).Error; err != nil {
	// 	return models.Permission{}, nil
	// }
	// fmt.Println()
	return pem, nil

}

func NewEmbeddedRepo(db *gorm.DB) *EmbeddedRepo {
	return &EmbeddedRepo{db}
}
