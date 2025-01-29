package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type AuthRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (a AuthRepository) GetUserIDByEmail(email string) (string, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) Create(data model.Auth) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (a AuthRepository) Delete(uid string) *core.Error {
	//TODO implement me
	panic("implement me")
}
