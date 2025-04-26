package repository

import (
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (a AuthRepository) GetUserIDByEmail(email string) (string, *core.Error) {
	var user model.User

	if retour := a.DB.Table("user").Select("id").Where("email = ?", email).First(&user); retour.Error != nil {
		a.Log.Error(retour.Error.Error())
		return "", core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, email), retour.Error)
	}
	return user.Id, nil
}

func (a AuthRepository) GetAuthByRefreshToken(refreshToken string) (model.Auth, *core.Error) {
	var auth model.Auth

	if retour := a.DB.Table("auth").Where("refresh_token = ?", refreshToken).First(&auth); retour.Error != nil {
		a.Log.Error(retour.Error.Error())
		return model.Auth{}, core.NewError(http.StatusNotFound, "", retour.Error)
	}
	return auth, nil
}

func (a AuthRepository) Create(data model.Auth) *core.Error {
	if retour := a.DB.Table("auth").Create(&data); retour.Error != nil {
		a.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateAuth, retour.Error)
	}
	return nil
}

func (a AuthRepository) Delete(uid string) *core.Error {
	var auth model.Auth

	if retour := a.DB.Table("auth").Where("user_id = ?", uid).Delete(&auth); retour.Error != nil {
		a.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, "", retour.Error)
	}
	return nil
}
