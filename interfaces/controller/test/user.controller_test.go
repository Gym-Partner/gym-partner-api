package test

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/controller"
	"gitlab.com/gym-partner1/api/gym-partner-api/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"
)

const TOKEN = "test@gmail.com"

//func transformResponse(response gin.H) {
//	if code, ok := response["code"].(float64); ok {
//		response["code"] = int(code)
//	}
//
//	if dataArray, ok := response["data"].([]interface{}); ok {
//		var transformedArray []gin.H
//		for _, item := range dataArray {
//			// Convertir chaque élément en gin.H si possible
//			if itemMap, ok := item.(map[string]interface{}); ok {
//				transformedArray = append(transformedArray, gin.H(itemMap))
//			} else if itemGinH, ok := item.(gin.H); ok {
//				transformedArray = append(transformedArray, itemGinH)
//			}
//		}
//		response["data"] = transformedArray
//	}
//
//	if data, ok := response["data"].(map[string]interface{}); ok {
//		response["data"] = gin.H(data)
//	}
//}

func transformResponse(response gin.H) gin.H {
	// Convertir "code" en int si c'est un float64
	if code, ok := response["code"].(float64); ok {
		response["code"] = int(code)
	}

	// Convertir "data" en gin.H si c'est un map
	if data, ok := response["data"].(map[string]interface{}); ok {
		response["data"] = transformMap(data)
	}

	// Convertir "data" en slice de gin.H si c'est un tableau
	if dataArray, ok := response["data"].([]interface{}); ok {
		response["data"] = transformArray(dataArray)
	}

	return response
}

// Transforme une map[string]interface{} en gin.H avec traitement des champs imbriqués
func transformMap(input map[string]interface{}) gin.H {
	result := gin.H{}
	for key, value := range input {
		switch v := value.(type) {
		case float64:
			// Conversion des floats en int si applicable
			if float64(int(v)) == v {
				result[key] = int(v)
			} else {
				result[key] = v
			}
		case map[string]interface{}:
			// Conversion récursive des maps imbriquées
			result[key] = transformMap(v)
		case []interface{}:
			// Conversion des slices
			result[key] = transformArray(v)
		default:
			result[key] = v
		}
	}
	return result
}

// Transforme un slice []interface{} en []gin.H avec traitement des éléments
func transformArray(input []interface{}) []gin.H {
	var result []gin.H
	for _, item := range input {
		if itemMap, ok := item.(map[string]interface{}); ok {
			result = append(result, transformMap(itemMap))
		} else if itemGinH, ok := item.(gin.H); ok {
			result = append(result, itemGinH)
		}
	}
	return result
}

func TestUserController_CREATE(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name         string
		setupMock    func(uc *mock.UserControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONUserCreateSuccess,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("Create", ctx).Return(user, (*core.Error)(nil)).Once()
			},
			expectedCode: http.StatusCreated,
			expectedBody: user.Respons(),
		},
		{
			name: core.TestUserCreateFailed,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("Create", ctx).Return(model.User{}, core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser)).Once()
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(http.StatusInternalServerError, core.ErrDBCreateUser).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserControllerMock := new(mock.UserControllerMock)

			uc := controller.MockUserController(UserControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			buf, _ := utils.StructToReadCloser(user)
			ctx.Request = &http.Request{
				Body: buf,
			}

			value.setupMock(UserControllerMock, ctx)

			uc.Create(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			UserControllerMock.AssertExpectations(t)
		})
	}
}

func TestUserController_GETALL(t *testing.T) {
	var users model.Users
	users.GenerateTestStruct()

	setupTest := []struct {
		name         string
		setupMock    func(uc *mock.UserControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONUserGetAllSuccess,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetAll").Return(users, (*core.Error)(nil)).Once()
			},
			expectedCode: http.StatusOK,
			expectedBody: users.Respons(),
		},
		{
			name: core.TestUsersNotFound,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetAll").Return(model.Users{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetAllUser)).Once()
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(http.StatusInternalServerError, core.ErrDBGetAllUser).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserControllerMock := new(mock.UserControllerMock)

			uc := controller.MockUserController(UserControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			ctx.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{TOKEN},
				},
			}

			value.setupMock(UserControllerMock, ctx)

			uc.GetAll(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			UserControllerMock.AssertExpectations(t)
		})
	}
}

func TestUserController_GETONE(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name         string
		setupMock    func(uc *mock.UserControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONUserGetOneSuccess,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetOne", ctx).Return(user, (*core.Error)(nil)).Once()
			},
			expectedCode: http.StatusOK,
			expectedBody: user.Respons(),
		},
		{
			name: core.TestUserNotFound,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetOne", ctx).Return(model.User{}, core.NewError(http.StatusInternalServerError, core.ErrDBGetOneUser)).Once()
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(http.StatusInternalServerError, core.ErrDBGetOneUser).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserControllerMock := new(mock.UserControllerMock)
			uc := controller.MockUserController(UserControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			ctx.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{TOKEN},
				},
			}

			value.setupMock(UserControllerMock, ctx)

			uc.GetOne(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			UserControllerMock.AssertExpectations(t)
		})
	}
}

func TestUserController_UPDATE(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name         string
		setupMock    func(uc *mock.UserControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONUserUpdateSuccess,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("Update", ctx).Return((*core.Error)(nil)).Once()
			},
			expectedCode: http.StatusOK,
			expectedBody: nil,
		},
		{
			name: core.TestUserUpdateFailed,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("Update", ctx).Return(core.NewError(http.StatusInternalServerError, core.ErrDBUpdateUser)).Once()
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(http.StatusInternalServerError, core.ErrDBUpdateUser).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserControllerMock := new(mock.UserControllerMock)
			uc := controller.MockUserController(UserControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			ctx.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{TOKEN},
				},
			}

			value.setupMock(UserControllerMock, ctx)

			uc.Update(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			UserControllerMock.AssertExpectations(t)
		})
	}
}

func TestUserController_DELETE(t *testing.T) {
	setupTest := []struct {
		name         string
		setupMock    func(uc *mock.UserControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONUserDeleteSuccess,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("Delete", ctx).Return((*core.Error)(nil)).Once()
			},
			expectedCode: http.StatusOK,
			expectedBody: nil,
		},
		{
			name: core.TestUserDeleteFailed,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("Delete", ctx).Return(core.NewError(http.StatusInternalServerError, core.ErrDBDeleteUser)).Once()
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: core.NewError(http.StatusInternalServerError, core.ErrDBDeleteUser).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserControllerMock := new(mock.UserControllerMock)
			uc := controller.MockUserController(UserControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)

			ctx.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{TOKEN},
				},
			}

			value.setupMock(UserControllerMock, ctx)

			uc.Delete(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			UserControllerMock.AssertExpectations(t)
		})
	}
}

func TestUserController_LOGIN(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	setupTest := []struct {
		name         string
		setupMock    func(uc *mock.UserControllerMock, ctx *gin.Context)
		expectedCode int
		expectedBody gin.H
	}{
		{
			name: core.TestCONUserLoginSuccess,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetOneByEmail", ctx).Return(user, (*core.Error)(nil)).Once()
				uc.On("Login", user).Return(TOKEN, (*core.Error)(nil)).Once()
			},
			expectedCode: http.StatusOK,
			expectedBody: gin.H{
				"token": TOKEN,
			},
		},
		{
			name: core.TestUserNotFound,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetOneByEmail", ctx).Return(model.User{}, core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, user.Email))).Once()
			},
			expectedCode: http.StatusNotFound,
			expectedBody: core.NewError(http.StatusNotFound, fmt.Sprintf(core.ErrDBGetOneUser, user.Email)).Respons(),
		},
		{
			name: core.TestUserLoginFailed,
			setupMock: func(uc *mock.UserControllerMock, ctx *gin.Context) {
				uc.On("GetOneByEmail", ctx).Return(user, (*core.Error)(nil)).Once()
				uc.On("Login", user).Return("", core.NewError(http.StatusUnauthorized, core.ErrAWSCognitoAuthUser))
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: core.NewError(http.StatusUnauthorized, core.ErrAWSCognitoAuthUser).Respons(),
		},
	}

	for _, value := range setupTest {
		t.Run(value.name, func(t *testing.T) {
			UserControllerMock := new(mock.UserControllerMock)
			uc := controller.MockUserController(UserControllerMock)

			gin.SetMode(gin.TestMode)
			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			ctx.Request = &http.Request{
				Header: http.Header{
					"Authorization": []string{TOKEN},
				},
			}

			value.setupMock(UserControllerMock, ctx)

			uc.Login(ctx)
			assert.Equal(t, value.expectedCode, res.Code)

			var response gin.H
			err := json.Unmarshal(res.Body.Bytes(), &response)
			transformResponse(response)

			assert.NoError(t, err)
			assert.Equal(t, value.expectedBody, response)

			UserControllerMock.AssertExpectations(t)
		})
	}
}
