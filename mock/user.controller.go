package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type UserControllerMock struct {
	mock.Mock
}

func (u *UserControllerMock) UploadImage(ctx *gin.Context) (model.UserImage, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) Search(query string, limit, offset int) (model.Users, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) Create(ctx *gin.Context) (model.User, *core.Error) {
	args := u.Called(ctx)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserControllerMock) GetAll() (model.Users, *core.Error) {
	args := u.Called()
	return args.Get(0).(model.Users), args.Error(1).(*core.Error)
}

func (u *UserControllerMock) GetOne(ctx *gin.Context) (model.User, *core.Error) {
	args := u.Called(ctx)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserControllerMock) GetOneByEmail(ctx *gin.Context) (model.User, *core.Error) {
	args := u.Called(ctx)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserControllerMock) Update(ctx *gin.Context) *core.Error {
	args := u.Called(ctx)
	return args.Error(0).(*core.Error)
}

func (u *UserControllerMock) Delete(ctx *gin.Context) *core.Error {
	args := u.Called(ctx)
	return args.Error(0).(*core.Error)
}

func (u *UserControllerMock) Login(user model.User) (string, *core.Error) {
	args := u.Called(user)
	return args.String(0), args.Error(1).(*core.Error)
}
