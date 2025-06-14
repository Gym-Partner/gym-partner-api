package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IWorkoutRepository interface {
	IsExist(uid string) bool

	CreateWorkouts(data model.Workout) *core.Error
	CreateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error
	CreateExercise(data model.Exercise) *core.Error
	CreateSeries(data model.Serie) *core.Error

	GetOneWorkoutsByUserId(uid string) (database.MigrateWorkout, *core.Error)
	GetUnitiesById(id string) (database.MigrateUnityOfWorkout, *core.Error)
	GetExerciseById(id string) (database.MigrateExercise, *core.Error)
	GetSeriesById(id string) (database.MigrateSerie, *core.Error)

	GetAllWorkoutsByUserId(uid string) (database.MigrateWorkouts, *core.Error)

	UpdateWorkouts(data model.Workout) *core.Error
	UpdateUnitiesOfWorkout(data model.UnityOfWorkout) *core.Error
	UpdateExercise(data model.Exercise) *core.Error
	UpdateSeries(data model.Serie) *core.Error

	DeleteWorkoutsByUserId(uid string) *core.Error
}
