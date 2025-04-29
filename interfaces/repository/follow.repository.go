package repository

import (
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type FollowRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (fr FollowRepository) IsExist(userId string) bool {
	var follow model.Follow

	if retour := fr.DB.Table("follow").Where("user_id = ?", userId).First(&follow); retour.Error != nil {
		fr.Log.Error(retour.Error.Error())
		return false
	}

	if follow.Id == "" {
		fr.Log.Error("error")
		return false
	} else {
		return true
	}
}

func (fr FollowRepository) Create(data model.Follow) (model.Follow, *core.Error) {
	// TODO implement me
	panic("implement me")
}

func (fr FollowRepository) GetAll() (model.Follows, *core.Error) {
	// TODO implement me
	panic("implement me")
}
