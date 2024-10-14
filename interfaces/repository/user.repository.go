package repository

import (
    "fmt"
    "github.com/google/uuid"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gitlab.com/gym-partner1/api/gym-partner-api/utils"
    "gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
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

    if retour := u.DB.Table("users").Where(queryColumn + " = ?", data).First(&user); retour.Error != nil {
        u.Log.Error(retour.Error.Error())
        return false
    }

    if user.Id == "" {
        u.Log.Error("User not found")
        return false
    } else {
        return true
    }
}

func (u UserRepository) Create(data model.User) (model.User, *core.Error) {
    var user model.User
    var err *core.Error

    data.Id = uuid.New().String()
    data.Password, err = utils.HashPassword(data.Password)
    if err != nil {
        return model.User{}, err
    }

    if retour := u.DB.Table("users").Create(&data); retour.Error != nil {
        u.Log.Error(retour.Error.Error())
        return model.User{}, core.NewError(500, "Failed to create user in the database", retour.Error)
    }

    user.Id = data.Id
    user.FirstName = data.FirstName
    user.LastName = data.LastName
    user.UserName = data.UserName
    user.Email = data.Email

    return user, nil
}

func (u UserRepository) GetAll() (model.Users, *core.Error) {
    var users model.Users

    if retour := u.DB.Table("users").Select("id, first_name, last_name, username, email").Find(&users); retour.Error != nil {
        u.Log.Error(retour.Error.Error())
        return model.Users{}, core.NewError(500, "Failed to recover all of users", retour.Error)
    }

    return users, nil
}

func (u UserRepository) GetOneById(uid string) (model.User, *core.Error) {
    var user model.User

    if retour := u.DB.Table("users").Where("id = ?", uid).Select("id, first_name, last_name, username, email").First(&user); retour.Error != nil {
        u.Log.Error(retour.Error.Error())
        return model.User{}, core.NewError(500, fmt.Sprintf("Failed to recover the user with this ID : %s", uid), retour.Error)
    }

    return user, nil
}