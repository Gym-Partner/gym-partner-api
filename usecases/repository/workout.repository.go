package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IWorkoutRepository interface {
	CreateWorkout(data model.Workout) *core.Error
	CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error
	CreateExcercice(data model.Exercice) *core.Error
	CreateSerie(data model.Serie) *core.Error

	GetOneWorkoutByUserId(uid string) (database.MigrateWorkout, *core.Error)
	GetUntiesById(ids []string) (database.MigrateUnitiesOfWorkout, *core.Error)
	GetExercicesById(ids []string) (database.MigrateExercices, *core.Error)
	GetSeriesById(ids []string) (database.MigrateSeries, *core.Error)

	GetAllWorkoutByUserId(uid string) (database.MigrateWorkouts, *core.Error)
}
