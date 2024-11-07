package interactor

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type WorkoutInteractor struct {
	IWorkoutRepository repository.IWorkoutRepository
	IUtils             utils.IUtils[model.Workout]
}

func (wi *WorkoutInteractor) Create(ctx *gin.Context) (model.Workout, *core.Error) {
	uid, _ := ctx.Get("uid")
	data, err := wi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return model.Workout{}, err
	}
	data.ChargeData(*uid.(*string))

	return data, nil
}
