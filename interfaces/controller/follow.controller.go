package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/interfaces/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/services/interactor"
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

// AddFollower godoc
// @Summary Add follower
// @Description Add one user's followers
// @Tags Follows
// @Accept json
// @Param Authorization header string true "User Token"
// @Success 202 {object} nil "Follower successfully added"
// @Failure 401 {object} core.Error "Follower already exist in database"
// @Failure 500 {object} core.Error "Internal server error"
// @Router /user/follower/add_follower [post]
func (fc *FollowController) AddFollower(ctx *gin.Context) {
	if err := fc.IFollowInteractor.AddFollower(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

// RemoveFollower godoc
// @Summary Remove follower
// @Description Remove one user's followers
// @Tags Follows
// @Accept json
// @Param Authorization header string true "User Token"
// @Success 200 {object} nil "Follower successfully removed"
// @Failure 401 {object} core.Error "Follower not exist in database"
// @Failure 500 {object} core.Error "Internal server error"
// @Router /user/follower/remove_follower [post]
func (fc *FollowController) RemoveFollower(ctx *gin.Context) {
	if err := fc.IFollowInteractor.RemoveFollower(ctx); err != nil {
		ctx.JSON(err.Code, err.Respons())
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
