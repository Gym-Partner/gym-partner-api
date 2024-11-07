package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
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
			IUtils: utils.Utils[model.Workout]{},
		},
		Log: db.Logger,
	}
}

func (wc *WorkoutController) Create(ctx *gin.Context) {
	workout, err := wc.WorkoutInteractor.Create(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusCreated, workout.Respons())
}
