package repos

import (
	"errors"

	"gorm.io/gorm"

	"github.com/Sea-of-Keys/seaofkeys-api/api/middleware"
	"github.com/Sea-of-Keys/seaofkeys-api/api/models"
	"github.com/Sea-of-Keys/seaofkeys-api/api/security"
)

type AuthRepo struct {
	db *gorm.DB
}

func (repo *AuthRepo) PostLogin(user models.Login) (*models.User, error) {
	var checkUser models.User

	if err := repo.db.Debug().First(&checkUser, "email = ?", user.Email).Error; err != nil {
		return nil, errors.New("CAN'T FIND YOU NIGGA")
	}

	if !middleware.CheckPasswordHash(user.Password, *checkUser.Password) {
		return nil, errors.New("PLZ BE THIS")
	}
	checkUser.Password = nil
	checkUser.Code = nil

	return &checkUser, nil
}
func (repo *AuthRepo) PostLogout() (bool, error) {
	return true, nil
}

// func (repo *AuthRepo)

// Change Embedde Password ### CHANGE CODE TO int ?????
func (repo *AuthRepo) PutPassword(id uint, code string) (*models.User, error) {
	return nil, nil
}
func (repo *AuthRepo) CheckTokenData(id uint, email string) (string, error) {
	var user models.User
	if err := repo.db.Debug().First(&user, id).Error; err != nil {
		return "", err
	}
	if user.ID != id || *user.Email != email {
		return "", errors.New("user id or email does not match")
	}
	// tokenData.ID = user.ID
	// tokenData.Token = user.Email
	token, _ := security.NewToken(id, email)
	return token, nil
}

func NewAuthRepo(db *gorm.DB) AuthRepoInterface {
	return &AuthRepo{db}
}
