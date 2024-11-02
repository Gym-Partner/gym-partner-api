package core

import (
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/mock"
    "gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
)

type Mock[T model.User] struct {
    mock.Mock
}

func (m *Mock[T]) HashPassword(password string) (string, *Error) {
    args := m.Called(password)
    return args.Get(0).(string), args.Error(1).(*Error)
}

func (m *Mock[T]) InjectBodyInModel(ctx *gin.Context) (T, *Error) {
    args := m.Called(ctx)
    return args.Get(0).(T), args.Error(1).(*Error)
}

func (m *Mock[T]) Bind(target, patch interface{}) *Error {
    //TODO implement me
    panic("implement me")
}

func (m *Mock[T]) IsExist(data, OPT string) bool {
    args := m.Called(data, OPT)
    return args.Bool(0)
}

func (m *Mock[T]) GetAll() (model.Users, *Error) {
    args := m.Called()
    return args.Get(0).(model.Users), args.Error(1).(*Error)
}

func (m *Mock[T]) GetOneById(uid string) (model.User, *Error) {
    args := m.Called(uid)
    return args.Get(0).(model.User), args.Error(1).(*Error)
}

func (m *Mock[T]) GetOneByEmail(email string) (model.User, *Error) {
    args := m.Called(email)
    return args.Get(0).(model.User), args.Error(1).(*Error)
}

func (m *Mock[T]) Create(data model.User) (model.User, *Error) {
    args := m.Called(data)
    return args.Get(0).(model.User), args.Error(1).(*Error)
}

func (m *Mock[T]) Update(data model.User) *Error {
    args := m.Called(data)
    return args.Error(0).(*Error)
}

func (m *Mock[T]) Delete(uid string) *Error {
    args := m.Called(uid)
    return args.Error(0).(*Error)
}
