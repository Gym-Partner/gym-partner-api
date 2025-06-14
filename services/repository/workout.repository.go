package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IWorkoutRepository interface {
	CreateWorkout(data model.Workout) *core.Error
	CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error
	CreateExercise(data model.Exercice) *core.Error
	CreateSeries(data model.Serie) *core.Error

	GetOneWorkoutByUserId(uid string) (database.MigrateWorkout, *core.Error)
	GetUnityById(id string) (database.MigrateUnityOfWorkout, *core.Error)
	GetExerciseById(id string) (database.MigrateExercice, *core.Error)
	GetSeriesById(id string) (database.MigrateSerie, *core.Error)

	GetAllWorkoutByUserId(uid string) (database.MigrateWorkouts, *core.Error)

	UpdateWorkout(data model.Workout) *core.Error
	UpdateUnityOfWorkout(data model.UnityOfWorkout) *core.Error
	UpdateExercise(data model.Exercice) *core.Error
	UpdateSeries(data model.Serie) *core.Error

	DeleteWorkoutByUserId(uid string) *core.Error
}
