package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type UtilsMock[T model.User] struct {
	mock.Mock
}

func (u *UtilsMock[T]) SchemaToModel(workout database.MigrateWorkout, unitie database.MigrateUnitiesOfWorkout, exercices database.MigrateExercices, series database.MigrateSeries) model.Workout {
	//TODO implement me
	panic("implement me")
}

func (u *UtilsMock[T]) HashPassword(password string) (string, *core.Error) {
	args := u.Called(password)
	return args.Get(0).(string), args.Error(1).(*core.Error)
}

func (u *UtilsMock[T]) InjectBodyInModel(ctx *gin.Context) (T, *core.Error) {
	args := u.Called(ctx)
	return args.Get(0).(T), args.Error(1).(*core.Error)
}

func (u *UtilsMock[T]) Bind(target, patch interface{}) *core.Error {
	args := u.Called(&target, patch)
	return args.Error(0).(*core.Error)
}

func (u *UtilsMock[T]) GenerateUUID() string {
	args := u.Called()
	return args.Get(0).(string)
}
