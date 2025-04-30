package repository

import (
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

type FollowRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (fr FollowRepository) FollowerIsExistByFollowedId(data model.Follow) bool {
	var newData model.Follow

	if retour := fr.DB.Table("follows").Where("follower_id = ? AND followed_id = ?", data.FollowerId, data.FollowedId).First(&newData); retour.Error != nil {
		fr.Log.Error(retour.Error.Error())
		return false
	}

	if newData.FollowedId == "" {
		fr.Log.Error("Follower not found in database")
		return false
	} else {
		return true
	}
}

func (fr FollowRepository) AddFollower(data model.Follow) *core.Error {
	if retour := fr.DB.Table("follows").Create(&data); retour.Error != nil {
		fr.Log.Error("Failed to add follower | originalErr: %s", retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to add follower for %s user", data.FollowedId),
			retour.Error)
	}
	return nil
}

func (fr FollowRepository) RemoveFollower(data model.Follow) *core.Error {
	if retour := fr.DB.Table("follows").Where("follower_id = ? AND followed_id = ?", data.FollowerId, data.FollowedId).Delete(&data); retour.Error != nil {
		fr.Log.Error("Failed to remove follower | originalErr: %s", retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to remove follower for %s user", data.FollowedId),
			retour.Error)
	}
	return nil
}

func (fr FollowRepository) GetAllByUserId(followedId string) (model.Follows, *core.Error) {
	var follows model.Follows

	if retour := fr.DB.Table("follows").Where("followed_id = ?", followedId).Find(&follows); retour.Error != nil {
		fr.Log.Error("Failed to get follows | originalErr: %s", retour.Error.Error())

		return model.Follows{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf("Failed to get follows for %s user", followedId),
			retour.Error)
	}
	return follows, nil
}
