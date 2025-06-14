package repository

import (
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

const AUTH_TABLE_NAME = "auth"

type AuthRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (a AuthRepository) GetUserIDByEmail(email string) (string, *core.Error) {
	var user model.User

	if raw := a.DB.
		Table(USERS_TABLE_NAME).
		Select("id").
		Where("email = ?", email).
		First(&user); raw.Error != nil {
		a.Log.Error(raw.Error.Error())
		return "", core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, email, raw.Error.Error()), raw.Error)
	}
	return user.Id, nil
}

func (a AuthRepository) GetAuthByRefreshToken(refreshToken string) (model.Auth, *core.Error) {
	var auth model.Auth

	if raw := a.DB.
		Table(AUTH_TABLE_NAME).
		Where("refresh_token = ?", refreshToken).
		First(&auth); raw.Error != nil {
		a.Log.Error(raw.Error.Error())
		return model.Auth{}, core.NewError(http.StatusNotFound, "", raw.Error)
	}
	return auth, nil
}

func (a AuthRepository) Create(data model.Auth) *core.Error {
	if raw := a.DB.
		Table(AUTH_TABLE_NAME).
		Create(&data); raw.Error != nil {
		a.Log.Error(raw.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateAuth, raw.Error)
	}
	return nil
}

func (a AuthRepository) Delete(uid string) *core.Error {
	var auth model.Auth

	if raw := a.DB.
		Table(AUTH_TABLE_NAME).
		Where("user_id = ?", uid).
		Delete(&auth); raw.Error != nil {
		a.Log.Error(raw.Error.Error())
		return core.NewError(http.StatusInternalServerError, "", raw.Error)
	}
	return nil
}
