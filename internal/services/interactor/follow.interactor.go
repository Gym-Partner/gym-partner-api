package interactor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gitlab.com/gym-partner1/api/gym-partner-api/services/repository"
	"gitlab.com/gym-partner1/api/gym-partner-api/utils"
	"net/http"
)

type IFollowInteractor interface {
	AddFollower(ctx *gin.Context) *core.Error
	RemoveFollower(ctx *gin.Context) *core.Error
}

type FollowInteractor struct {
	IFollowRepository repository.IFollowRepository
	IUtils            utils.IUtils[model.Follow]
}

func (fi *FollowInteractor) AddFollower(ctx *gin.Context) *core.Error {
	uid, _ := ctx.Get("uid")
	data, err := fi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}

	data.FollowedId = uid.(string)

	exist := fi.IFollowRepository.FollowerIsExistByFollowedId(data)
	if exist {
		return core.NewError(
			http.StatusUnauthorized,
			fmt.Sprintf(core.ErrAppINTFollowerExist, data.FollowerId, data.FollowedId))
	}

	if err := fi.IFollowRepository.AddFollower(data); err != nil {
		return err
	}
	return nil
}

func (fi *FollowInteractor) RemoveFollower(ctx *gin.Context) *core.Error {
	uid, _ := ctx.Get("uid")
	data, err := fi.IUtils.InjectBodyInModel(ctx)
	if err != nil {
		return err
	}

	data.FollowedId = uid.(string)

	exist := fi.IFollowRepository.FollowerIsExistByFollowedId(data)
	if !exist {
		return core.NewError(
			http.StatusUnauthorized,
			fmt.Sprintf(core.ErrAppINTFollowerNotExist, data.FollowerId, data.FollowedId))
	}

	if err := fi.IFollowRepository.RemoveFollower(data); err != nil {
		return err
	}
	return nil
}
