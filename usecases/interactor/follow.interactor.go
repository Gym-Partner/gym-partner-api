package interactor

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
)

type IFollowInteractor interface {
	AddFollower(ctx *gin.Context) *core.Error
	RemoveFollower(ctx *gin.Context) *core.Error
	GetFollowers(ctx *gin.Context) (model.Follows, *core.Error)
}

type FollowInteractor struct {
	IFollowRepository repository.IFollowRepository
	IUtils            utils.IUtils[model.Follow]
}

func (fi *FollowInteractor) AddFollower(ctx *gin.Context) *core.Error {
	// TODO implement me
	panic("implement me")
}

func (fi *FollowInteractor) RemoveFollower(ctx *gin.Context) *core.Error { panic("implement me") }

func (fi *FollowInteractor) GetFollowers(ctx *gin.Context) (model.Follows, *core.Error) {
	panic("implement me")
}
