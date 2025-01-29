package middleware

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
	"time"
)

type IAuthMiddleware interface {
	// GenerateJWT create two type of token. First the real token and second the refreshToken
	// The distinction is in the secretKey and expiredTime
	GenerateJWT(uid, secretKey string, expiredTime time.Duration, endTime ...*time.Time) (string, error)
}

type AuthMiddleware struct{}

func (a AuthMiddleware) GenerateJWT(uid, secretKey string, expiredTime time.Duration, endTime ...*time.Time) (string, error) {
	expirationTime := time.Now().Add(expiredTime)
	claims := &model.CustomClaims{
		UserId: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	if len(endTime) > 0 {
		*endTime[0] = expirationTime
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString(secretKey)))
}
