package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
	"gitlab.com/gym-partner1/api/gym-partner-api/model"
)

func InitMiddleware(log *core.Log) gin.HandlerFunc {
	return func(context *gin.Context) {}
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header["Authorization"]
		newToken := strings.Join(token, "")
		if len(newToken) <= 0 {
			ctx.JSON(500, gin.H{
				"message": "Token undefined in request's headers",
			})
			ctx.Abort()
			return
		}

		claims := &model.CustomClaims{}
		tokenAfterParse, err := jwt.ParseWithClaims(newToken, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("TOKEN_SECRET")), nil
		})

		if err != nil || !tokenAfterParse.Valid {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error to parse token",
			})
			ctx.Abort()
			return
		}

		if isTokenExpired(newToken) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is expired",
			})
			ctx.Abort()
			return
		}

		ctx.Set("uid", claims.UserId)
		ctx.Set("token", newToken)
	}
}

func isTokenExpired(tokenStr string) bool {
	claims := &model.CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("TOKEN_SECRET")), nil
	})
	if err != nil {
		return true
	}

	if !token.Valid {
		return true
	}

	if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
		return true
	}

	return false
}
