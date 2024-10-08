package repository

import "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"

type IUserRepository interface {
    IsExist(data, OPT string) bool

	GetAll() (model.Users, error)
    GetOneById(uid string) (model.User, error)

    Create(data model.User) (model.User, error)
    Update(data model.User) error
    Delete(uid string) error
}