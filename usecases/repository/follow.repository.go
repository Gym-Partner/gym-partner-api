package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IFollowRepository interface {
	FollowerIsExistByFollowedId(followedId string) bool
	GetAllByUserId(followedId string) (model.Follows, *core.Error)
	AddFollower(data model.Follow) *core.Error
	RemoveFollower(data model.Follow) *core.Error
}
