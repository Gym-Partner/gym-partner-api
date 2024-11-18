package test

import (
	"fmt"
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
			name: core.TestINTCreateSuccess,
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
				userMock.On("Create", user).Return(model.User{}, core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser), error(nil)).Once()
			},
			expectedRes: model.User{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser, error(nil)),
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
			case core.TestINTCreateSuccess:
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
				//assert.Equal(t, err, value.expectedErr)
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
			name: core.TestINTGetAllSuccess,
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
			case core.TestINTGetAllSuccess:
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

func TestUserInteractor_GETONE(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name        string
		setupMock   func(userMock *mock.UserMock)
		expectedRes model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestINTGetOneSuccess,
			setupMock: func(userMock *mock.UserMock) {
				userMock.On("GetOneById", user.Id).Return(user, (*core.Error)(nil)).Once()
			},
			expectedRes: user,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestUserNotFound,
			setupMock: func(userMock *mock.UserMock) {
				userMock.On("GetOneById", user.Id).Return(model.User{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetOneUser)).Once()
			},
			expectedRes: model.User{},
			expectedErr: core.NewError(http.StatusInternalServerError, core.ErrDBGetOneUser),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)
			context := &gin.Context{}
			context.Set("uid", &user.Id)

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)
			value.setupMock(UserMock)

			result, err := ui.GetOne(context)

			switch value.name {
			case core.TestINTGetOneSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUserNotFound:
				assert.NotNil(t, err)
				assert.Empty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
			UserMock.AssertExpectations(t)
		})
	}
}

func TestUserInteractor_GETONEBYEMAIL(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name        string
		setupMock   func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context)
		expectedRes model.User
		expectedErr *core.Error
	}{
		{
			name: core.TestINTGetOneSuccess,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context) {
				utilsMock.On("InjectBodyInModel", context).Return(user, (*core.Error)(nil)).Once()
				userMock.On("GetOneByEmail", user.Email).Return(user, (*core.Error)(nil)).Once()
			},
			expectedRes: user,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.TestUserNotFound,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context) {
				utilsMock.On("InjectBodyInModel", context).Return(user, (*core.Error)(nil)).Once()
				userMock.On("GetOneByEmail", user.Email).Return(model.User{}, core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, user.Email)))
			},
			expectedRes: model.User{Password: user.Password},
			expectedErr: core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, user.Email)),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)
			buf, _ := utils.StructToReadCloser(user)

			context := &gin.Context{
				Request: &http.Request{
					Body: buf,
				},
			}

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)
			value.setupMock(UserMock, UtilsMock, context)

			result, err := ui.GetOneByEmail(context)

			switch value.name {
			case core.TestINTGetOneSuccess:
				assert.Nil(t, err)
				assert.NotEmpty(t, result)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.TestUserNotFound:
				assert.NotNil(t, err)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
			UserMock.AssertExpectations(t)
			UtilsMock.AssertExpectations(t)
		})
	}
}

func TestUserInteractor_UPDATE(t *testing.T) {
	var target, patch model.User
	target.GenerateTestStruct()
	patch.GenerateTestStruct(target.Id)

	setupTest := []struct {
		name        string
		setupMock   func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context)
		expectedRes *core.Error
	}{
		{
			name: core.TestINTUdateSuccess,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context) {
				utilsMock.On("InjectBodyInModel", context).Return(patch, (*core.Error)(nil)).Once()
				userMock.On("IsExist", patch.Id, "ID").Return(true).Once()
				userMock.On("GetOneById", patch.Id).Return(target, (*core.Error)(nil)).Once()
				utilsMock.On("Bind", &target, patch).Return((*core.Error)(nil)).Once()
				userMock.On("Update", target).Return((*core.Error)(nil)).Once()
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestUserNotExistFailed,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context) {
				utilsMock.On("InjectBodyInModel", context).Return(patch, (*core.Error)(nil)).Once()
				userMock.On("IsExist", patch.Id, "ID").Return(false).Once()
			},
			expectedRes: core.NewError(http.StatusBadRequest, fmt.Sprintf(core.ErrIntUserNotExist, target.Id)),
		},
		{
			name: core.TestUserNotFound,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context) {
				utilsMock.On("InjectBodyInModel", context).Return(patch, (*core.Error)(nil)).Once()
				userMock.On("IsExist", patch.Id, "ID").Return(true).Once()
				userMock.On("GetOneById", patch.Id).Return(model.User{}, core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, target.Email))).Once()
			},
			expectedRes: core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, target.Email)),
		},
		{
			name: core.TestInternalErrorFailed,
			setupMock: func(userMock *mock.UserMock, utilsMock *mock.UtilsMock[model.User], context *gin.Context) {
				utilsMock.On("InjectBodyInModel", context).Return(patch, (*core.Error)(nil)).Once()
				userMock.On("IsExist", patch.Id, "ID").Return(true).Once()
				userMock.On("GetOneById", patch.Id).Return(target, (*core.Error)(nil)).Once()
				utilsMock.On("Bind", &target, patch).Return((*core.Error)(nil)).Once()
				userMock.On("Update", target).Return(core.NewError(http.StatusInternalServerError, fmt.Sprintf(core.ErrDBUpdateUser, target.Id))).Once()
			},
			expectedRes: core.NewError(http.StatusInternalServerError, fmt.Sprintf(core.ErrDBUpdateUser, target.Id)),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)
			buf, _ := utils.StructToReadCloser(patch)

			context := &gin.Context{
				Request: &http.Request{
					Body: buf,
				},
			}
			context.Set("uid", &patch.Id)

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)

			value.setupMock(UserMock, UtilsMock, context)

			result := ui.Update(context)

			switch value.name {
			case core.TestINTGetOneSuccess:
				assert.Nil(t, result)
				assert.Equal(t, result, value.expectedRes)
			case core.TestUserNotExistFailed:
				assert.NotNil(t, result)
				assert.Equal(t, result, value.expectedRes)
			case core.TestUserNotFound:
				assert.NotNil(t, result)
				assert.Equal(t, result, value.expectedRes)
			case core.TestInternalErrorFailed:
				assert.NotNil(t, result)
				assert.Equal(t, result, value.expectedRes)
			}
			UserMock.AssertExpectations(t)
			UtilsMock.AssertExpectations(t)
		})
	}
}

func TestUserInteractor_DELETE(t *testing.T) {
	var token string
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name        string
		setupMock   func(userMock *mock.UserMock, cognitoMock *mock.CognitoMock)
		expectedRes *core.Error
	}{
		{
			name: core.TestINTDeleteSuccess,
			setupMock: func(userMock *mock.UserMock, cognitoMock *mock.CognitoMock) {
				userMock.On("IsExist", user.Id, "ID").Return(true).Once()
				userMock.On("Delete", user.Id).Return((*core.Error)(nil)).Once()
				cognitoMock.On("DeleteUser", token).Return((*core.Error)(nil)).Once()
			},
			expectedRes: (*core.Error)(nil),
		},
		{
			name: core.TestUserNotExistFailed,
			setupMock: func(userMock *mock.UserMock, cognitoMock *mock.CognitoMock) {
				userMock.On("IsExist", user.Id, "ID").Return(false).Once()
			},
			expectedRes: core.NewError(http.StatusBadRequest, fmt.Sprintf(core.ErrIntUserNotExist, user.Id)),
		},
		{
			name: core.TestInternalErrorFailed,
			setupMock: func(userMock *mock.UserMock, cognitoMock *mock.CognitoMock) {
				userMock.On("IsExist", user.Id, "ID").Return(true).Once()
				userMock.On("Delete", user.Id).Return(core.NewError(http.StatusInternalServerError, fmt.Sprintf(core.ErrDBDeleteUser, user.Id)))
			},
			expectedRes: core.NewError(http.StatusInternalServerError, fmt.Sprintf(core.ErrDBDeleteUser, user.Id)),
		},
		{
			name: core.TestUserNotDeletedCognito,
			setupMock: func(userMock *mock.UserMock, cognitoMock *mock.CognitoMock) {
				userMock.On("IsExist", user.Id, "ID").Return(true).Once()
				userMock.On("Delete", user.Id).Return((*core.Error)(nil)).Once()
				cognitoMock.On("DeleteUser", token).Return(core.NewError(http.StatusBadRequest, core.ErrAWSCognitoDeleteUser)).Once()
			},
			expectedRes: core.NewError(http.StatusBadRequest, core.ErrAWSCognitoDeleteUser),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)
			context := &gin.Context{}

			context.Set("uid", &user.Id)
			context.Set("token", token)

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)
			value.setupMock(UserMock, CognitoMock)

			result := ui.Delete(context)

			switch value.name {
			case core.TestINTDeleteSuccess:
				assert.Nil(t, result)
				assert.Equal(t, result, value.expectedRes)
			case core.TestUserNotExistFailed:
				assert.NotNil(t, result)
				assert.Equal(t, result, value.expectedRes)
			case core.TestInternalErrorFailed:
				assert.NotNil(t, result)
				assert.Equal(t, result, value.expectedRes)
			case core.TestUserNotDeletedCognito:
				assert.NotNil(t, result)
				assert.Equal(t, result, value.expectedRes)
			}
			UserMock.AssertExpectations(t)
			CognitoMock.AssertExpectations(t)
		})
	}
}

func TestUserInteractor_LOGIN(t *testing.T) {
	var token string
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name        string
		setupMock   func(cognitoMock *mock.CognitoMock)
		expectedRes string
		expectedErr *core.Error
	}{
		{
			name: core.TestINTLoginSuccess,
			setupMock: func(cognitoMock *mock.CognitoMock) {
				cognitoMock.On("SignIn", user).Return(token, (*core.Error)(nil)).Once()
			},
			expectedRes: token,
			expectedErr: (*core.Error)(nil),
		},
		{
			name: core.ErrAWSCognitoAuthUser,
			setupMock: func(cognitoMock *mock.CognitoMock) {
				cognitoMock.On("SignIn", user).Return("", core.NewError(http.StatusUnauthorized, core.ErrAWSCognitoAuthUser))
			},
			expectedRes: "",
			expectedErr: core.NewError(http.StatusUnauthorized, core.ErrAWSCognitoAuthUser),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserMock := new(mock.UserMock)
			UtilsMock := new(mock.UtilsMock[model.User])
			CognitoMock := new(mock.CognitoMock)

			ui := interactor.MockUserInteractor(UserMock, UtilsMock, CognitoMock)
			value.setupMock(CognitoMock)

			result, err := ui.Login(user)

			switch value.name {
			case core.TestINTLoginSuccess:
				assert.Nil(t, err)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			case core.ErrAWSCognitoAuthUser:
				assert.NotNil(t, err)
				assert.Equal(t, result, value.expectedRes)
				assert.Equal(t, err, value.expectedErr)
			}
			CognitoMock.AssertExpectations(t)
		})
	}
}
