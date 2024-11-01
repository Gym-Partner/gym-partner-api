package core

import (
    "github.com/stretchr/testify/mock"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) IsExist(data, OPT string) bool {
    args := m.Called(data, OPT)
    return args.Bool(0)
}

func (m *Mock) GetAll() (model.Users, *Error) {
    args := m.Called()
    return args.Get(0).(model.Users), args.Error(1).(*Error)
}

func (m *Mock) GetOneById(uid string) (model.User, *Error) {
    args := m.Called(uid)
    return args.Get(0).(model.User), args.Error(1).(*Error)
}

func (m *Mock) GetOneByEmail(email string) (model.User, *Error) {
    args := m.Called(email)
    return args.Get(0).(model.User), args.Error(1).(*Error)
}

func (m *Mock) Create(data model.User) (model.User, *Error) {
    args := m.Called(data)
    return args.Get(0).(model.User), args.Error(1).(*Error)
}

func (m *Mock) Update(data model.User) *Error {
    args := m.Called(data)
    return args.Error(0).(*Error)
}

func (m *Mock) Delete(uid string) *Error {
    args := m.Called(uid)
    return args.Error(0).(*Error)
}
