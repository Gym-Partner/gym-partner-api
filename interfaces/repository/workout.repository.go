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

func (wr WorkoutRepository) Create(data model.Workout) *core.Error {
	// Implement me
	panic("implement me")
}
