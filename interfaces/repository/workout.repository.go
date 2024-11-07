package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type WorkoutRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (wr WorkoutRepository) CreateWorkout(data model.Workout) *core.Error {
	// Implement me
	panic("implement me")
}

func (wr WorkoutRepository) CreateUnityOfWorkout(data model.UnityOfWorkout) *core.Error {
	// Implement me
	panic("implement me")
}

func (wr WorkoutRepository) CreateExcercice(data model.Exercice) *core.Error {
	// Implement me
	panic("implement me")
}

func (wr WorkoutRepository) CreateSerie(data model.Serie) *core.Error {
	// Implement me
	panic("implement me")
}
