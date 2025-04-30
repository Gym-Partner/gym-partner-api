package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/database"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IFollowRepository interface {
	IsExist(userId string) bool
	GetByUserId(userId string) (database.MigrateFollow, *core.Error)
	Create(data model.Follow) (model.Follow, *core.Error)
}
