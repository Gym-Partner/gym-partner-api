package repository

import (
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (u UserRepository) IsExist(data, OPT string) bool {
	var user model.User
	var queryColumn string

	switch OPT {
	case "ID":
		queryColumn = "id"
	case "EMAIL":
		queryColumn = "email"
	}

	if retour := u.DB.Table("users").Where(queryColumn+" = ?", data).First(&user); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return false
	}

	if user.Id == "" {
		u.Log.Error(core.ErrDBUserNotFound)
		return false
	} else {
		return true
	}
}

func (u UserRepository) Create(data model.User) (model.User, *core.Error) {
	if retour := u.DB.Table("users").Create(&data); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return model.User{}, core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser, retour.Error)
	}

	return data, nil
}

func (u UserRepository) GetAll() (model.Users, *core.Error) {
	var users model.Users

	if retour := u.DB.Table("users").Select("id, first_name, last_name, username, email").Find(&users); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return model.Users{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetAllUser, retour.Error)
	}

	return users, nil
}

func (u UserRepository) GetOneById(uid string) (model.User, *core.Error) {
	var user model.User

	if retour := u.DB.Table("users").Where("id = ?", uid).First(&user); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return model.User{}, core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, uid), retour.Error)
	}

	return user, nil
}

func (u UserRepository) GetOneByEmail(email string) (model.User, *core.Error) {
	var user model.User

	if retour := u.DB.Table("users").Where("email = ?", email).Select("id").First(&user); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return model.User{}, core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, email), retour.Error)
	}

	return user, nil
}

func (u UserRepository) Update(data model.User) *core.Error {
	if retour := u.DB.Table("users").Save(&data); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, fmt.Sprintf(core.ErrDBUpdateUser, data.Id), retour.Error)
	}

	return nil
}

func (u UserRepository) Delete(uid string) *core.Error {
	var user model.User

	if retour := u.DB.Table("users").Where("id = ?", uid).Delete(&user); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return core.NewError(http.StatusInternalServerError, fmt.Sprintf(core.ErrDBDeleteUser, uid), retour.Error)
	}

	return nil
}
