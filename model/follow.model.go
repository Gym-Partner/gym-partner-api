package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Follow struct {
	FollowerId string    `json:"follower_id"`
	FollowedId string    `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type Follows []Follow

func (f *Follow) Response() gin.H {
	return gin.H{
		"data": gin.H{
			"follower_id": f.FollowerId,
			"followed_id": f.FollowedId,
			"created_at":  f.CreatedAt,
		},
	}
}

func (f *Follows) Response() gin.H {
	var result []gin.H

	for _, follow := range *f {
		result = append(result, gin.H{
			"follower_id": follow.FollowerId,
			"followed_id": follow.FollowedId,
			"created_at":  follow.CreatedAt,
		})
	}

	return gin.H{
		"data": result,
	}
}

type UserFollows struct {
	Followers  []string `json:"followers"`
	Followings []string `json:"followings"`
}
