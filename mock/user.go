package mock

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type UserMock struct {
	mock.Mock
}

func (u *UserMock) IsExist(data, OPT string) bool {
	args := u.Called(data, OPT)
	return args.Bool(0)
}

func (u *UserMock) GetAll() (model.Users, *core.Error) {
	args := u.Called()
	return args.Get(0).(model.Users), args.Error(1).(*core.Error)
}

func (u *UserMock) GetOneById(uid string) (model.User, *core.Error) {
	args := u.Called(uid)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserMock) GetOneByEmail(email string) (model.User, *core.Error) {
	args := u.Called(email)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserMock) Create(data model.User) (model.User, *core.Error) {
	args := u.Called(data)
	return args.Get(0).(model.User), args.Error(1).(*core.Error)
}

func (u *UserMock) Update(data model.User) *core.Error {
	args := u.Called(data)
	return args.Error(0).(*core.Error)
}

func (u *UserMock) Delete(uid string) *core.Error {
	args := u.Called(uid)
	return args.Error(0).(*core.Error)
}
