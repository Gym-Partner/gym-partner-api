package repository

import (
	"fmt"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
	"net/http"
)

type AuthRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (a AuthRepository) GetUserIDByEmail(email string) (string, *core.Error) {
	var userId string

	if retour := a.DB.Table("user").Where("email = ?", email).Find(&userId); retour.Error != nil {
		a.Log.Error(retour.Error.Error())
		return userId, core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, email), retour.Error)
	}
	return userId, nil
}

func (a AuthRepository) Create(data model.Auth) *core.Error {
	if retour := a.DB.Table("auth").Create(&data); retour.Error != nil {
		a.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, core.ErrDBCreateAuth, retour.Error)
	}
	return nil
}

func (a AuthRepository) Delete(uid string) *core.Error {
	//TODO implement me
	panic("implement me")
}
