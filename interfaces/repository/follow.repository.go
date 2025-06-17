package repository

import (
	"fmt"
	"net/http"

	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"gorm.io/gorm"
)

const FOLLOWS_TABLE_NAME = "follows"

type FollowRepository struct {
	DB  *gorm.DB
	Log *core.Log
}

func (fr FollowRepository) FollowerIsExistByFollowedId(data model.Follow) bool {
	var newData model.Follow

	if raw := fr.DB.
		Table(FOLLOWS_TABLE_NAME).
		Where("follower_id = ? AND followed_id = ?", data.FollowerId, data.FollowedId).
		First(&newData); raw.Error != nil {
		fr.Log.Error(raw.Error.Error())
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
	if raw := fr.DB.
		Table(FOLLOWS_TABLE_NAME).
		Create(&data); raw.Error != nil {
		fr.Log.Error(core.ErrDBAddFollower, raw.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBAddFollower, data.FollowedId),
			raw.Error)
	}
	return nil
}

func (fr FollowRepository) RemoveFollower(data model.Follow) *core.Error {
	if raw := fr.DB.
		Table(FOLLOWS_TABLE_NAME).
		Where("follower_id = ? AND followed_id = ?", data.FollowerId, data.FollowedId).
		Delete(&data); raw.Error != nil {
		fr.Log.Error(core.ErrDBRemoveFollower, raw.Error.Error())

		return core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBRemoveFollower, data.FollowedId),
			raw.Error)
	}
	return nil
}

func (fr FollowRepository) GetAllByUserId(userId string) (model.UserFollows, *core.Error) {
	type followResult struct {
		ID string `gorm:"column:follower_id"`
	}
	type followerResult struct {
		ID string `gorm:"column:followed_id"`
	}

	var followResults []followResult
	var followerResults []followerResult
	var userFollows model.UserFollows

	if res := fr.DB.
		Table(FOLLOWS_TABLE_NAME).
		Where("followed_id = ?", userId).
		Select("follower_id").
		Find(&followResults); res.Error != nil {
		fr.Log.Error(core.ErrDBGetFollowers, res.Error.Error())
		return model.UserFollows{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBGetFollowers, userId),
			res.Error,
		)
	}

	for _, r := range followResults {
		userFollows.Followings = append(userFollows.Followings, r.ID)
	}

	if res := fr.DB.
		Table(FOLLOWS_TABLE_NAME).
		Where("follower_id = ?", userId).
		Select("followed_id").
		Find(&followerResults); res.Error != nil {
		fr.Log.Error(core.ErrDBGetFollowed, res.Error.Error())
		return model.UserFollows{}, core.NewError(
			http.StatusInternalServerError,
			fmt.Sprintf(core.ErrAppDBGetFollowed, userId),
			res.Error,
		)
	}

	for _, r := range followerResults {
		userFollows.Followers = append(userFollows.Followers, r.ID)
	}

	return userFollows, nil
}
