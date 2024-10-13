package repository

import (
    "github.com/google/uuid"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
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

func (u UserRepository) Create(data model.User) (model.User, error) {
    data.Id = uuid.New().String()

    if retour := u.DB.Table("users").Create(&data); retour.Error != nil {
        u.Log.Error(retour.Error.Error())
        return model.User{}, retour.Error
    }

    return data, nil
}