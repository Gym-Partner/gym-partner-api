package mock

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type UserInteractorMock struct {
	mock.Mock
}

func (u *UserInteractorMock) IsExist(data, OPT string) bool {
	args := u.Called(data, OPT)
	return args.Bool(0)
}

func (u *UserInteractorMock) GetAll() (model.Users, *core.Error) {
	args := u.Called()
	return args.Get(0).(model.Users), args.Error(1).(*core.Error)
}

func (u *UserInteractorMock) GetOneById(uid string) (model.User, *core.Error) {
	args := u.Called(uid)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserInteractorMock) GetOneByEmail(email string) (model.User, *core.Error) {
	args := u.Called(email)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserInteractorMock) Create(data model.User) (model.User, *core.Error) {
	args := u.Called(data)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserInteractorMock) Update(data model.User) *core.Error {
	args := u.Called(data)
	return args.Error(0).(*core.Error)
}

func (u *UserInteractorMock) Delete(uid string) *core.Error {
	args := u.Called(uid)
	return args.Error(0).(*core.Error)
}

func (u *UserInteractorMock) GetImageByUserId(uid string) (model.UserImage, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserInteractorMock) DeleteUserImage(uid string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (u *UserInteractorMock) UserImageIsExist(uid string) bool {
	//TODO implement me
	panic("implement me")
}

func (u *UserInteractorMock) UploadImage(data model.UserImage) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (u *UserInteractorMock) Search(query string, limit, offset int) (model.Users, *core.Error) {
	//TODO implement me
	panic("implement me")
}
