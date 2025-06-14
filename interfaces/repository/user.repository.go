package repository

import (
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

const USERS_TABLE_NAME = "users"
const USERS_IMAGES_TABLE_NAME = "users_image"

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

	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Where(queryColumn+" = ?", data).
		Find(&user); raw.Error != nil {
		u.Log.Error(raw.Error.Error())
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
	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Create(&data); raw.Error != nil {
		u.Log.Error(core.ErrDBCreateUser, raw.Error.Error())

		return model.User{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBCreateUser, data.Email),
			raw.Error)
	}

	return data, nil
}

func (u UserRepository) GetAll() (model.Users, *core.Error) {
	var users model.Users

	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Select("id, first_name, last_name, username, email").
		First(&users); raw.Error != nil {
		u.Log.Error(core.ErrDBGetAllUser, raw.Error.Error())

		return model.Users{}, core.NewError(
			http.StatusInternalServerError,
			core.ErrAppDBGetAllUser,
			raw.Error)
	}

	return users, nil
}

func (u UserRepository) GetOneById(uid string) (model.User, *core.Error) {
	var user model.User

	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Where("id = ?", uid).
		First(&user); raw.Error != nil {
		u.Log.Error(core.ErrDBGetOneUser, uid, raw.Error.Error())

		return model.User{}, core.NewError(
			http.StatusNotFound,
			fmt.Sprintf(core.ErrAppDBGetOneUser, uid),
			raw.Error)
	}

	return user, nil
}

func (u UserRepository) GetOneByEmail(email string) (model.User, *core.Error) {
	var user model.User

	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Where("email = ?", email).
		Select("id").
		First(&user); raw.Error != nil {
		u.Log.Error(core.ErrDBGetOneUser, email, raw.Error.Error())

		return model.User{}, core.NewError(
			http.StatusNotFound,
			fmt.Sprintf(core.ErrAppDBGetOneUser, email),
			raw.Error)
	}

	return user, nil
}

func (u UserRepository) Update(data model.User) *core.Error {
	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Save(&data); raw.Error != nil {
		u.Log.Error(core.ErrDBUpdateUser, data.Email, raw.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBUpdateUser, data.Email),
			raw.Error)
	}

	return nil
}

func (u UserRepository) Delete(uid string) *core.Error {
	var user model.User

	if raw := u.DB.
		Table(USERS_TABLE_NAME).
		Where("id = ?", uid).
		Delete(&user); raw.Error != nil {
		u.Log.Error(core.ErrDBDeleteUser, uid, raw.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBDeleteUser, uid),
			raw.Error)
	}

	return nil
}

func (u UserRepository) Search(query string, limit, offset int) (model.Users, *core.Error) {
	var users model.Users

	if raw := u.DB.
		Table("user").
		Where("LOWER(first_name) LIKE ? OR LOWER(last_name) LIKE ? OR LOWER(username) LIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").
		Select("id", "first_name", "last_name", "username", "age").
		Limit(limit).
		Offset(offset).
		First(&users); raw.Error != nil {
		u.Log.Error(core.ErrDBSearchUsers, raw.Error.Error())

		return model.Users{}, core.NewError(
			http.StatusNotFound,
			core.ErrAppDBSearchUsers,
			raw.Error)
	}

	return users, nil
}

func (u UserRepository) UploadImage(data model.UsersImage) *core.Error {
	if raw := u.DB.
		Table(USERS_IMAGES_TABLE_NAME).
		Create(&data); raw.Error != nil {
		u.Log.Error(core.ErrDBCreateUserImage, raw.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBCreateUserImage),
			raw.Error)
	}

	return nil
}

func (u UserRepository) UserImageIsExist(uid string) bool {
	var userImage model.UsersImage

	if raw := u.DB.
		Table(USERS_IMAGES_TABLE_NAME).
		Where("user_id = ?", uid).
		First(&userImage); raw.Error != nil {
		u.Log.Error(raw.Error.Error())
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
	var userImage model.UsersImage

	if raw := u.DB.
		Table(USERS_IMAGES_TABLE_NAME).
		Where("user_id = ?", uid).
		Delete(&userImage); raw.Error != nil {
		u.Log.Error(core.ErrDBDeleteUserImage, raw.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBDeleteUserImage),
			raw.Error)
	}
	return nil
}

func (u UserRepository) GetImageByUserId(uid string) (model.UsersImage, *core.Error) {
	var userImage model.UsersImage

	if raw := u.DB.
		Table(USERS_IMAGES_TABLE_NAME).
		Where("user_id = ?", uid).
		First(&userImage); raw.Error != nil {
		if errors.Is(raw.Error, gorm.ErrRecordNotFound) {
			return model.UsersImage{}, nil
		}

		u.Log.Error(core.ErrDBGetUserImage, uid, raw.Error.Error())

		return model.UsersImage{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBGetUserImage, uid),
			raw.Error)
	}

	return userImage, nil
}
