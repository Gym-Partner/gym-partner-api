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

func (fr FollowRepository) FollowerIsExistByFollowedId(followedId string) bool {
	// TODO implement me
	panic("implement me")
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

func (fr FollowRepository) GetAllByUserId(FollowedId string) (model.Follows, *core.Error) {
	// TODO implement me
	panic("implement me")
}
