package repository

import (
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

const TABLE_NAME = "user"

type UserRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func MockUserRepository(db *gorm.DB) *UserRepository {
	log := core.NewLog("/Users/oscar/Documents/gym-partner-env", true)
	log.ChargeLog()

	return &UserRepository{
		DB:  db,
		Log: log,
	}
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

	if retour := u.DB.Table(TABLE_NAME).Where(queryColumn+" = ?", data).Find(&user); retour.Error != nil {
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
	if retour := u.DB.Table(TABLE_NAME).Create(&data); retour.Error != nil {
		u.Log.Error(core.ErrDBCreateUser, retour.Error.Error())

		return model.User{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBCreateUser, data.Email),
			retour.Error)
	}

	return data, nil
}

func (u UserRepository) GetAll() (model.Users, *core.Error) {
	var users model.Users

	if retour := u.DB.Table(TABLE_NAME).Select("id, first_name, last_name, username, email").Find(&users); retour.Error != nil {
		u.Log.Error(core.ErrDBGetAllUser, retour.Error.Error())

		return model.Users{}, core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBGetAllUser,
			retour.Error)
	}

	return users, nil
}

func (u UserRepository) GetOneById(uid string) (model.User, *core.Error) {
	var user model.User

	if retour := u.DB.Table(TABLE_NAME).Where("id = ?", uid).Find(&user); retour.Error != nil {
		u.Log.Error(core.ErrDBGetOneUser, uid, retour.Error.Error())

		return model.User{}, core.NewError(
			http.StatusNotFound,
			fmt.Sprintf(core.ErrAppDBGetOneUser, uid),
			retour.Error)
	}

	return user, nil
}

func (u UserRepository) GetOneByEmail(email string) (model.User, *core.Error) {
	var user model.User

	if retour := u.DB.Table(TABLE_NAME).Where("email = ?", email).Select("id").Find(&user); retour.Error != nil {
		u.Log.Error(core.ErrDBGetOneUser, email, retour.Error.Error())

		return model.User{}, core.NewError(
			http.StatusNotFound,
			fmt.Sprintf(core.ErrAppDBGetOneUser, email),
			retour.Error)
	}

	return user, nil
}

func (u UserRepository) Update(data model.User) *core.Error {
	if retour := u.DB.Table(TABLE_NAME).Save(&data); retour.Error != nil {
		u.Log.Error(core.ErrDBUpdateUser, data.Email, retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBUpdateUser, data.Email),
			retour.Error)
	}

	return nil
}

func (u UserRepository) Delete(uid string) *core.Error {
	var user model.User

	if retour := u.DB.Table(TABLE_NAME).Where("id = ?", uid).Delete(&user); retour.Error != nil {
		u.Log.Error(core.ErrDBDeleteUser, uid, retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBDeleteUser, uid),
			retour.Error)
	}

	return nil
}

func (u UserRepository) Search(query string, limit, offset int) (model.Users, *core.Error) {
	var users model.Users

	if retour := u.DB.Table("user").
		Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(username) LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Select("id", "first_name", "last_name", "username", "age").
		Limit(limit).
		Offset(offset).Find(&users); retour.Error != nil {
		u.Log.Error(core.ErrDBSearchUsers, retour.Error.Error())

		return model.Users{}, core.NewError(
			http.StatusNotFound,
			core.ErrAppDBSearchUsers,
			retour.Error)
	}

	return users, nil
}

func (u UserRepository) UploadImage(data model.UserImage) *core.Error {
	if retour := u.DB.Table("user_image").Create(&data); retour.Error != nil {
		u.Log.Error(core.ErrDBCreateUserImage, retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBCreateUserImage),
			retour.Error)
	}

	return nil
}

func (u UserRepository) UserImageIsExist(uid string) bool {
	var userImage model.UserImage

	if retour := u.DB.Table("user_image").Where("user_id = ?", uid).First(&userImage); retour.Error != nil {
		u.Log.Error(retour.Error.Error())
		return false
	}

	if userImage.Id == "" {
		u.Log.Error(core.ErrDBUserImageNotFound)
		return false
	} else {
		return true
	}
}

func (u UserRepository) DeleteUserImage(uid string) *core.Error {
	var userImage model.UserImage

	if retour := u.DB.Table("user_image").Where("user_id = ?", uid).Delete(&userImage); retour.Error != nil {
		u.Log.Error(core.ErrDBDeleteUserImage, retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBDeleteUserImage),
			retour.Error)
	}
	return nil
}
