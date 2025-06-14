package mock

import (
	"github.com/stretchr/testify/mock"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type WorkoutInteractorMock struct {
	mock.Mock
}

func (w *WorkoutInteractorMock) CreateWorkouts(data model.Workout) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) CreateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) CreateExercise(data model.Exercise) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) CreateSeries(data model.Serie) *core.Error {
	args := w.Called(data)
	return args.Error(0).(*core.Error)
}

func (w *WorkoutInteractorMock) GetOneWorkoutsByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	args := w.Called(uid)
	return args.Get(0).(database.MigrateWorkout), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetUnitiesById(id string) (database.MigrateUnityOfWorkout, *core.Error) {
	args := w.Called(id)
	return args.Get(0).(database.MigrateUnityOfWorkout), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetExerciseById(id string) (database.MigrateExercise, *core.Error) {
	args := w.Called(id)
	return args.Get(0).(database.MigrateExercise), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetSeriesById(id string) (database.MigrateSerie, *core.Error) {
	args := w.Called(id)
	return args.Get(0).(database.MigrateSerie), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) GetAllWorkoutsByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	args := w.Called(uid)
	return args.Get(0).(database.MigrateWorkouts), args.Error(1).(*core.Error)
}

func (w *WorkoutInteractorMock) UpdateWorkouts(data model.Workout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) UpdateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) UpdateExercise(data model.Exercise) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) UpdateSeries(data model.Serie) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) DeleteWorkouts(id string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) DeleteUnities(id string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) DeleteExercises(id string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) DeleteSeries(id string) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkoutInteractorMock) IsExist(id string) bool {
	//TODO implement me
	panic("implement me")
}
