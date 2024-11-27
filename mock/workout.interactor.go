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

func (w WorkoutInteractorMock) CreateWorkout(data model.Workout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) CreateExcercice(data model.Exercice) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) CreateSerie(data model.Serie) *core.Error {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) GetOneWorkoutByUserId(uid string) (database.MigrateWorkout, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) GetUntyById(id string) (database.MigrateUnityOfWorkout, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) GetExerciceById(id string) (database.MigrateExercice, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) GetSerieById(id string) (database.MigrateSerie, *core.Error) {
	//TODO implement me
	panic("implement me")
}

func (w WorkoutInteractorMock) GetAllWorkoutByUserId(uid string) (database.MigrateWorkouts, *core.Error) {
	//TODO implement me
	panic("implement me")
}
