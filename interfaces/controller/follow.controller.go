package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/usecases/interactor"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
)

type FollowController struct {
	IFollowInteractor interactor.IFollowInteractor
	Log               *core.Log
}

func NewFollowController(db *core.Database) *FollowController {
	return &FollowController{
		IFollowInteractor: &interactor.FollowInteractor{
			IFollowRepository: repository.FollowRepository{
				DB:  db.Handler,
				Log: db.Logger,
			},
			IUtils: utils.Utils[model.Follow]{},
		},
		Log: db.Logger,
	}
}

func (fc *FollowController) AddFollower(ctx *gin.Context) {
	if err := fc.IFollowInteractor.AddFollower(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (fc *FollowController) RemoveFollower(ctx *gin.Context) {}

func (fc *FollowController) GetFollowers(ctx *gin.Context) {}
