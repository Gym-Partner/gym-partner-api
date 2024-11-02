package test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

func TestUserInteractor_INSERT(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	buf, _ := utils.StructToReadCloser(user)
	context := &gin.Context{
		Request: &http.Request{
			Body: buf,
		},
	}

	setupTest := []struct {
		name        string
		setupMock   func(mock *core.Mock[model.User])
		expectedRes *model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestCreateSuccess,
			setupMock: func(mock *core.Mock[model.User]) {
				mock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				mock.On("Create", user).Return(user, nil).Once()
			},
			expectedRes: &user,
			expectedErr: nil,
		},
		{
			name: core.TestUserExistFailed,
			setupMock: func(mock *core.Mock[model.User]) {
				mock.On("IsExist", user.Email, "EMAIL").Return(true).Once()
			},
			expectedRes: &model.User{},
			expectedErr: core.NewError(core.InternalErrCode, core.ErrDBUserExist),
		},
		{
			name: core.TestInternalErrorFailed,
			setupMock: func(mock *core.Mock[model.User]) {
				mock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				mock.On("Create", user).Return(nil, core.NewError(core.InternalErrCode, core.ErrDBCreateUser)).Once()
			},
			expectedRes: &model.User{},
			expectedErr: core.NewError(core.InternalErrCode, core.ErrDBCreateUser),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(core.Mock[model.User])
			ui := interactor.MockUserInteractor(UserMock)

			value.setupMock(UserMock)

			result, err := ui.Create(context)

			switch value.name {
			case core.TestCreateSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUserExistFailed:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestInternalErrorFailed:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
			UserMock.AssertExpectations(t)
		})
	}
}
