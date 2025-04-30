package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IFollowRepository interface {
	GetAllByUserId(userId string) (model.Follows, *core.Error)
	Create(data model.Follow) *core.Error
}
