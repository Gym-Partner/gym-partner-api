package model

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

type CustomClaims struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}

type UserToLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
