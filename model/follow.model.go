package model

import (
	"time"
)

type Follow struct {
	FollowerId string    `json:"follower_id"`
	FollowedId string    `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type Follows []Follow
