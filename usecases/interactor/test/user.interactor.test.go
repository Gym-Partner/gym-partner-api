package test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

func TestUserInteractor_INSERT(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	var context *gin.Context
	buf, _ := utils.StructToReadCloser(user)
	context.Request.Body = buf

	setupTest := []struct {
		name        string
		setupMock   func(mock *core.Mock)
		expectedRes *model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestCreateSuccess,
			setupMock: func(mock *core.Mock) {
				mock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				mock.On("Create", user).Return(user, nil).Once()
			},
			expectedRes: &user,
			expectedErr: nil,
		},
		{
			name: core.TestUserExistFailed,
			setupMock: func(mock *core.Mock) {
				mock.On("IsExist", user.Email, "EMAIL").Return(true).Once()
			},
			expectedRes: nil,
			expectedErr: core.NewError(core.InternalErrCode, core.ErrDBUserExist),
		},
		{
			name: core.TestInternalErrorFailed,
			setupMock: func(mock *core.Mock) {
				mock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				mock.On("Create", user).Return(nil, core.NewError(core.InternalErrCode, core.ErrDBCreateUser)).Once()
			},
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {

		})
	}
}
