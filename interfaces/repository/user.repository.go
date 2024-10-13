package repository

import (
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
    "gorm.io/gorm"
)

type UserRepository struct {
	Sql *gorm.DB
}

func (u UserRepository) Create(data model.User) (model.User, error) {
    //TODO implement me
    panic("implement me")
}

func (u UserRepository) IsExist(data, OPT string) bool {
    //TODO implement me
    panic("implement me")
}
