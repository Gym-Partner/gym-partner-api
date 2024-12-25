package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type WorkoutControllerMock struct {
	mock.Mock
}

func (w *WorkoutControllerMock) Create(ctx *gin.Context) *core.Error {
	args := w.Called(ctx)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutControllerMock) GetOneByUserId(ctx *gin.Context) (model.Workout, *core.Error) {
	args := w.Called(ctx)
	return args.Get(0).(model.Workout), args.Error(1).(*core.Error)
}
