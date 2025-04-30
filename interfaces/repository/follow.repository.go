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
		fr.Log.Error(core.ErrDBAddFollower, retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBAddFollower, data.FollowedId),
			retour.Error)
	}
	return nil
}

func (fr FollowRepository) RemoveFollower(data model.Follow) *core.Error {
	if retour := fr.DB.Table("follows").Where("follower_id = ? AND followed_id = ?", data.FollowerId, data.FollowedId).Delete(&data); retour.Error != nil {
		fr.Log.Error(core.ErrDBRemoveFollower, retour.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBRemoveFollower, data.FollowedId),
			retour.Error)
	}
	return nil
}

func (fr FollowRepository) GetAllByUserId(userId string) (model.UserFollows, *core.Error) {
	var followed []string
	var followers []string
	var userFollows model.UserFollows

	if followedReturn := fr.DB.Table("follows").Where("followed_id = ?", userId).Select("follower_id").Find(&followed); followedReturn.Error != nil {
		fr.Log.Error(core.ErrDBGetFollowers, followedReturn.Error.Error())

		return model.UserFollows{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBGetFollowers, userId),
			followedReturn.Error)
	}
	userFollows.Followings = followed

	if followerReturn := fr.DB.Table("follows").Where("follower_id = ?", userId).Select("followed_id").Find(&followers); followerReturn.Error != nil {
		fr.Log.Error(core.ErrDBGetFollowed, followerReturn.Error.Error())

		return model.UserFollows{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBGetFollowed, userId),
			followerReturn.Error)
	}
	userFollows.Followers = followers

	return userFollows, nil
}
