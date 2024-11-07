package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
)

type WorkoutController struct {
	WorkoutInteractor interactor.WorkoutInteractor
	Log               *core.Log
}

func NewWorkoutController(db *core.Database) *WorkoutController {
	return &WorkoutController{
		WorkoutInteractor: interactor.WorkoutInteractor{
			IWorkoutRepository: repository.WorkoutRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
		},
		Log: db.Logger,
	}
}

func (wc *WorkoutController) Create(ctx *gin.Context) {
	// Implement me
	panic("implement me")
}
