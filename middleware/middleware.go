package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
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

		// Make change for recover user_id in token without AWS Cognito
		uid := ""

		ctx.Set("uid", uid)
		ctx.Set("token", newToken)
	}
}
