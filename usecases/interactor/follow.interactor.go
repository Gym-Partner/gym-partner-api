package interactor

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IFollowInteractor interface {
	Create(ctx *gin.Context) (model.Follow, *core.Error)
}

type FollowInteractor struct {
	IFollowRepository repository.IFollowRepository
	IUtils            utils.IUtils[model.Follow]
}

func (fi *FollowInteractor) Create(ctx *gin.Context) (model.Follow, *core.Error) {
	// TODO implement me
	panic("implement me")
}
