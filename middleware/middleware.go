package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/gym-partner1/api/gym-partner-api/core"
)

func InitMiddleware(log *core.Log) gin.HandlerFunc {
	return func(context *gin.Context) {
		congito := core.NewCognito(log)
		context.Set("cognito", congito)
	}
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cognito, _ := ctx.Get("cognito")

		token := ctx.Request.Header["Authorization"]
		newToken := strings.Join(token, "")
		if len(newToken) <= 0 {
			ctx.JSON(500, gin.H{
				"message": "Token undefined in request's headers",
			})
			ctx.Abort()
			return
		}

		uid, err := cognito.(*core.Cognito).GetUserByToken(newToken)
		if err != nil {
			ctx.JSON(err.Code, err.Respons())
			ctx.Abort()
			return
		}

		ctx.Set("uid", uid)
		ctx.Set("token", newToken)
	}
}
