package interactor

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
)

type WorkoutInteractor struct {
	IWorkoutRepository repository.IWorkoutRepository
}

func (wi *WorkoutInteractor) Create(ctx *gin.Context) *core.Error {
	// Implement me
	panic("implement me")
}
