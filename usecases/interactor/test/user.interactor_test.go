package test

import (
	"net/http"
	"testing"

	"gitlab.com/gym-partner1/api/gym-partner-api/mock"

	"gitlab.com/gym-partner1/api/gym-partner-api/utils"

	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

func TestUserInteractor_INSERT(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name        string
		setupMock   func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], cognitoMock *mock.CognitoMock, ctx *gin.Context)
		expectedRes model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestCreateSuccess,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], cognitoMock *mock.CognitoMock, ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(user, (*core.Error)(nil)).Once()
				userMock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				utilsMock.On("GenerateUUID").Return(user.Id).Once()
				utilsMock.On("HashPassword", user.Password).Return(user.Password, (*core.Error)(nil)).Once()
				userMock.On("Create", user).Return(user, (*core.Error)(nil)).Once()
				cognitoMock.On("SignUp", user).Return((*core.Error)(nil)).Once()
			},
			expectedRes: user,
			expectedErr: nil,
		},
		{
			name: core.TestUserExistFailed,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], cognitoMock *mock.CognitoMock, ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(user, (*core.Error)(nil)).Once()
				userMock.On("IsExist", user.Email, "EMAIL").Return(true).Once()
			},
			expectedRes: model.User{},
			expectedErr: core.NewError(http.StatusBadRequest, core.ErrIntUserExist),
		},
		{
			name: core.TestInternalErrorFailed,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], cognitoMock *mock.CognitoMock, ctx *gin.Context) {
				utilsMock.On("InjectBodyInModel", ctx).Return(user, (*core.Error)(nil)).Once()
				userMock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				utilsMock.On("GenerateUUID").Return(user.Id).Once()
				utilsMock.On("HashPassword", user.Password).Return(user.Password, (*core.Error)(nil))
				userMock.On("Create", user).Return(model.User{}, core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser)).Once()
			},
			expectedRes: model.User{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)

			buf, _ := utils.StructToReadCloser(user)
			context := &gin.Context{
				Request: &http.Request{
					Body: buf,
				},
			}

			value.setupMock(UserMock, UtilsMock, CognitoMock, context)

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
			UtilsMock.AssertExpectations(t)
			CognitoMock.AssertExpectations(t)
		})
	}
}

func TestUserInteractor_GETALL(t *testing.T) {
	var users model.Users
	users.GenerateTestStruct()

	setupTest := []struct {
		name        string
		setupMock   func(userMock *mock.UserMock)
		expectedRes model.Users
		expectedErr *core.Error
	}{
		{
			name: core.TestGetAllSuccess,
			setupMock: func(userMock *mock.UserMock) {
				userMock.On("GetAll").Return(users, (*core.Error)(nil)).Once()
			},
			expectedRes: users,
			expectedErr: nil,
		},
		{
			name: core.TestUsersNotFound,
			setupMock: func(userMock *mock.UserMock) {
				userMock.On("GetAll").Return(model.Users{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetAllUser)).Once()
			},
			expectedRes: model.Users{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetAllUser),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)

			value.setupMock(UserMock)

			result, err := ui.GetAll()

			switch value.name {
			case core.TestGetAllSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUsersNotFound:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
			UserMock.AssertExpectations(t)
		})
	}
}

func TestUserInteractor_GETONE(t *testing.T) {}

func TestUserInteractor_UPDATE(t *testing.T) {}

func TestUserInteractor_DELETE(t *testing.T) {}

func TestUserInteractor_LOGIN(t *testing.T) {}
