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
