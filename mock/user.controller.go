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

func (u *UserControllerMock) Create(ctx *gin.Context) (model.User, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) GetAll() (model.Users, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) GetOne(c *gin.Context) (model.User, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) GetOneByEmail(ctx *gin.Context) (model.User, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) Update(ctx *gin.Context) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) Delete(ctx *gin.Context) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerMock) Login(user model.User) (string, *core.Error) {
	//TODO implement me
	panic("implement me")
}
