package controller

import (
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/mock"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/services/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type WorkoutController struct {
	IWorkoutInteractor interactor.IWorkoutInteractor
	Log                *core.Log
}

func NewWorkoutController(db *core.Database) *WorkoutController {
	return &WorkoutController{
		IWorkoutInteractor: &interactor.WorkoutInteractor{
			IWorkoutRepository: repository.WorkoutRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
			IUtils: utils.Utils[model.Workout]{},
		},
		Log: db.Logger,
	}
}

func MockWorkoutController(WCMock *mock.WorkoutControllerMock) *WorkoutController {
	log := core.NewLog("/Users/oscar/Documents/gym-partner-env", true)
	log.ChargeLog()

	return &WorkoutController{
		IWorkoutInteractor: WCMock,
		Log:                log,
	}
}

// Create godoc
// @Summary Create user's workout
// @Schemes
// @Description Create new user's workout in database and return code created
// @Tags Workout
// @Accept json
// @Param Authorization header string true "User's token"
// @Param user_workout body model.Workout{} true "User's workout"
// @Success 201 {object} nil "User's workout created"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /workout/create [post]
func (wc *WorkoutController) Create(ctx *gin.Context) {
	if err := wc.IWorkoutInteractor.Create(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

// GetOneByUserId godoc
// @Summary Get one workout with user id
// @Schemes
// @Description Get one workout with user id from database
// @Tags Workout
// @Produce application/json
// @Param Authorization header string true "User's token"
// @Success 200 {object} model.Workout{} "User's workout"
// @Failure 500 {object} core.Error{} "Internal server error"
// @Router /user/workout/getOne [get]
func (wc *WorkoutController) GetOneByUserId(ctx *gin.Context) {
	workout, err := wc.IWorkoutInteractor.GetOneByUserId(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, workout.Respons())
}

func (wc *WorkoutController) GetAllByUserId(ctx *gin.Context) {
	workouts, err := wc.IWorkoutInteractor.GetAllByUserId(ctx)
	if err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, workouts.Respons())
}

func (wc *WorkoutController) Update(ctx *gin.Context) {
	if err := wc.IWorkoutInteractor.Update(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (wc *WorkoutController) Delete(ctx *gin.Context) {
	if err := wc.IWorkoutInteractor.Delete(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
