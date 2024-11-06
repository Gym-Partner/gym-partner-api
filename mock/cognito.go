package mock

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type CognitoMock struct {
	mock.Mock
}

func (cm *CognitoMock) SignUp(user model.User) *core.Error {
	args := cm.Called(user)
	return args.Error(0).(*core.Error)
}

func (cm *CognitoMock) SignIn(user model.User) (string, *core.Error) {
	args := cm.Called(user)
	return args.String(0), args.Error(1).(*core.Error)
}

func (cm *CognitoMock) GetUserByToken(token string) (*string, *core.Error) {
	args := cm.Called(token)
	return args.Get(0).(*string), args.Error(1).(*core.Error)
}

func (cm *CognitoMock) DeleteUser(token string) *core.Error {
	args := cm.Called(token)
	return args.Error(0).(*core.Error)
}
