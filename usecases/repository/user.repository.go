package repository

import (
)

type IUserRepository interface {
    // IsExist(data, OPT string) bool

	// GetAll() (model.Users, error)
    // GetOneById(uid string) (model.User, error)

    // Create(data model.User) (model.User, error)
    // Update(data model.User) error
    // Delete(uid string) error

    PING() error
}