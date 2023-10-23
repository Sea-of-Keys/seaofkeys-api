package repos

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
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
func (r *EmbeddedRepo) UpdateSecrect(oldtoken string, newtoken string) (bool, error) {
	var em models.Embedded
	if err := r.db.Debug().Where("scret = ?", oldtoken).Find(&em).Error; err != nil {
		return false, err
	}
	if em.ID == 0 {
		return false, errors.New("not found")
	}
	em.Scret = newtoken
	if err := r.db.Debug().Model(&em).Updates(&em).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (r *EmbeddedRepo) PostEmbeddedSetup(emb models.EmbedSetup) (bool, error) {
	var em models.Embedded
	if err := r.db.Debug().First(&em, emb.EmbeddedID).Error; err != nil {
		return false, err
	}
	fmt.Printf("embeebedID: %v\n", em.ID)
	if em.ID == 0 {
		// fmt.
		return false, errors.New("not found")
	}
	fmt.Printf("ssshhh: %v\nScret: %v\n", emb.Ssshhh, em.Scret)
	if emb.Ssshhh != em.Scret {
		fmt.Printf("ssshhh: %v\nScret: %v\n", emb.Ssshhh, em.Scret)
		return false, errors.New("token not a match")
	}

	return true, nil
}

func (r *EmbeddedRepo) PostCodeLogin(code, userID string, roomID uint) (bool, error) {
	var user models.User
	var pem models.Permission
	userIdInt, _ := strconv.Atoi(userID)
	currentTime := time.Now()
	day := time.Now().Weekday()
	dayINT := int(day)
	formattedTime := currentTime.Format("15:04:05")
	// SQL SELECT * FROM permissions AS p WHERE p.user_id = 1 AND p.room_id = 3 OR p.room_id = 3 AND p.team_id IN (SELECT team_id FROM teams_users WHERE team_id = p.team_id AND user_id = 1);
	// if err := r.db.Debug().Raw("SELECT * FROM permissions AS p WHERE p.user_id = ? AND p.room_id = ? OR p.room_id = ? AND p.team_id IN (SELECT team_id FROM teams_users WHERE team_id = p.team_id AND user_id = ?)", userIdInt, roomID, roomID, userIdInt).Scan(&pem).Error; err != nil {
	// 	return models.Permission{}, err
	// }
	if err := r.db.Debug().Find(&user, userIdInt).Error; err != nil {
		return false, err
	}
	if err := r.db.Debug().Table("permissions").
		Preload("Team.Users").
		Preload("User").
		Preload("Weekdays").
		Where("user_id = ? AND room_id = ?", userIdInt, roomID).Or("room_id = ? AND team_id IN (SELECT team_id FROM teams_users WHERE team_id = permissions.team_id AND user_id = ?)", roomID, userIdInt).
		Find(&pem).Error; err != nil {
		return false, err
	}
	fmt.Println(pem.UserID)
	fmt.Println(pem.UserID)
	fmt.Println(pem.UserID)
	fmt.Println(pem.UserID)
	fmt.Println(pem.UserID)
	if pem.ID != 0 {
		if !security.CheckPasswordHash(code, *user.Code) {
			return false, errors.New("code or user not a match")
		}
		pemSTimeStr := pem.StartTime.String()
		pemETimeStr := pem.EndTime.String()

		pemSTime, _ := time.Parse("15:04:05", pemSTimeStr)
		pemETime, _ := time.Parse("15:04:05", pemETimeStr)

		pemSTimeFormatted := pemSTime.Format("15:04:05")
		pemETimeFormatted := pemETime.Format("15:04:05")
		if true {
			if pemSTimeFormatted < formattedTime && pemETimeFormatted > formattedTime {
				for _, v := range pem.Weekdays {
					fmt.Printf("Day: %v\nTime.Now Date: %v", v.Day, dayINT)
					fmt.Printf("Day: %v\nTime.Now Date: %v", v.Day, dayINT)
					fmt.Printf("Day: %v\nTime.Now Date: %v", v.Day, dayINT)
					fmt.Printf("Day: %v\nTime.Now Date: %v", v.Day, dayINT)
					if v.Day == dayINT {
						fmt.Printf("Day: %v\n", v.Day)
						var newLogin models.History
						newLogin.UserID = user.ID
						newLogin.PermissionID = pem.ID
						if ok, err := r.PostHistoryLogin(newLogin); ok && err == nil {
							return true, nil
						} else {
							return false, err
						}
					}
				}
			}
		}
	}

	return false, errors.New("not found")
}
func (r *EmbeddedRepo) PostHistoryLogin(newLogin models.History) (bool, error) {
	currentTime := time.Now()
	layout := "2006-01-02 15:04:05"
	formattedTime := currentTime.Format(layout)
	newLogin.At = formattedTime
	if err := r.db.Debug().Create(&newLogin).Error; err != nil {
		return false, err
	}
	return true, nil

}
func NewEmbeddedRepo(db *gorm.DB) EmbeddedRepoInterface {
	return &EmbeddedRepo{db}
}
