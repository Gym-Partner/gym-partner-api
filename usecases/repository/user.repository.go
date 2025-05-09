package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IUserRepository interface {
	IsExist(data, OPT string) bool

	GetAll() (model.Users, *core.Error)
	GetOneById(uid string) (model.User, *core.Error)
	GetOneByEmail(email string) (model.User, *core.Error)

	Create(data model.User) (model.User, *core.Error)
	Update(data model.User) *core.Error
	Delete(uid string) *core.Error

	Search(query string, limit, offset int) (model.Users, *core.Error)

	UserImageIsExist(uid string) bool
	GetImageByUserId(uid string) (model.UserImage, *core.Error)
	UploadImage(data model.UserImage) *core.Error
	DeleteUserImage(uid string) *core.Error
}
