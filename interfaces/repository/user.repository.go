package repository

import (
    "gitlab.com/Titouan-Esc/api_common/logger"
    mongodb "gitlab.com/Titouan-Esc/api_common/mongo"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
)

type UserRepository struct {
	Sql *mongodb.Mongo
	Logger *logger.Log
}

func (u UserRepository) IsExist(data, OPT string) bool {
    //TODO implement me
    panic("implement me")
}

func (u UserRepository) GetAll() (model.Users, error) {
    //TODO implement me
    panic("implement me")
}

func (u UserRepository) GetOneById(uid string) (model.User, error) {
    //TODO implement me
    panic("implement me")
}

func (u UserRepository) Create(data model.User) error {
    //TODO implement me
    panic("implement me")
}

func (u UserRepository) Update(data model.User) error {
    //TODO implement me
    panic("implement me")
}

func (u UserRepository) Delete(uid string) error {
    //TODO implement me
    panic("implement me")
}
