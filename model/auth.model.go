package model

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Auth struct {
	Id           string    `json:"id"`
	UserId       string    `json:"user_id"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (t *Auth) Response() gin.H {
	return gin.H{
		"data": gin.H{
			"user_id":       t.UserId,
			"token":         t.Token,
			"refresh_token": t.RefreshToken,
			"expires_at":    t.ExpiresAt,
		},
	}
}
