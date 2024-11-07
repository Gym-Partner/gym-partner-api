package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IWorkoutRepository interface {
	CreateWorkout(data model.Workout) *core.Error
	CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error
	CreateExcercice(data model.Exercice) *core.Error
	CreateSerie(data model.Serie) *core.Error
}
