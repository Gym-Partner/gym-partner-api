package test

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/domain/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"testing"
)

const (
	TestSuccess = "SUCCESS"
	TestFailed = "FAILED"
)

func TestUserInteractor_INSERT(t *testing.T) {
	var user model.User
	user.GenerateTestStruct()

	var context *gin.Context

	buf, _ := utils.StructToReadCloser(user)
	context.Request.Body = buf

	setupTest := []struct{
		name string
		setupMock func(mock *core.Mock)
		expectedRes *model.User
		expectedErr *core.Error
	}{
		{
			name: TestSuccess,
			setupMock: func(mock *core.Mock) {
				mock.On("IsExist", user.Email, "EMAIL").Return(false).Once()
				mock.On("Create", user).Return(user, nil).Once()
			},
		},
	}

	for _, value := range setupTest {
		
	}
}