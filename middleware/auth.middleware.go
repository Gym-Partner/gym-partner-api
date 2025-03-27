package middleware

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

type IAuthMiddleware interface {
	// GenerateJWT create two type of token. First the real token and second the refreshToken
	// The distinction is in the secretKey and expiredTime
	GenerateJWT(uid, secretKey string, expiredTime time.Duration, endTime ...*time.Time) (string, error)
	VerifyRefreshToken(auth model.Auth, refreshToken string) *core.Error
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

func (a AuthMiddleware) VerifyRefreshToken(auth model.Auth, refreshToken string) *core.Error {
	claims := &model.CustomClaims{
		UserId: auth.UserId,
	}

	token, parseErr := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("REFRESH_TOKEN_SECRET")), nil
	})

	if parseErr != nil || !token.Valid {
		return core.NewError(http.StatusUnauthorized, "Invalid or expired refresh token", parseErr)
	}

	if claims.UserId != auth.UserId {
		return core.NewError(http.StatusUnauthorized, "Invalid refresh token for the user")
	}
	return nil
}
