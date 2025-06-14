package mock

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type WorkoutInteractorMock struct {
	mock.Mock
}

func (w *WorkoutInteractorMock) CreateWorkout(data model.Workout) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) CreateExercise(data model.Exercice) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) CreateSeries(data model.Serie) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) GetOneWorkoutByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	args := w.Called(uid)
	return args.Get(0).(database.MigrateWorkout), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetUnityById(id string) (database.MigrateUnityOfWorkout, *core.Error) {
	args := w.Called(id)
	return args.Get(0).(database.MigrateUnityOfWorkout), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetExerciseById(id string) (database.MigrateExercice, *core.Error) {
	args := w.Called(id)
	return args.Get(0).(database.MigrateExercice), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetSeriesById(id string) (database.MigrateSerie, *core.Error) {
	args := w.Called(id)
	return args.Get(0).(database.MigrateSerie), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetAllWorkoutByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	args := w.Called(uid)
	return args.Get(0).(database.MigrateWorkouts), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) UpdateWorkout(data model.Workout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) UpdateUnityOfWorkout(data model.UnityOfWorkout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) UpdateExercise(data model.Exercice) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) UpdateSeries(data model.Serie) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) DeleteWorkoutByUserId(uid string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) IsExist(ctx *gin.Context) bool {
	//TODO implement me
	panic("implement me")
}
