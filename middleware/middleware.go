package middleware

import (
    "github.com/gin-gonic/gin"
    "gitlab.com/gym-partner1/api/gym-partner-api/core"
    "strings"
)

func InitMiddleware(log *core.Log) gin.HandlerFunc {
	return func(context *gin.Context) {
		aws := core.NewAWS(log)

		context.Set("aws", aws)
		context.Set("logs", log)
    }
}

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		aws, _ := ctx.Get("aws")

        token := ctx.Request.Header["Authorization"]
		newToken := strings.Join(token, "")
        if len(token) <= 0 {
			ctx.JSON(500, gin.H{
				"message": "Token undefined in request's headers",
			})
			ctx.Abort()
        }

		cognito, err := aws.(*core.AWS).NewCognito()
		if err != nil {
			ctx.JSON(err.Code, err.Respons())
			ctx.Abort()
		}

		uid, err := cognito.GetUserByToken(newToken)
		if err != nil {
			ctx.JSON(err.Code, err.Respons())
			ctx.Abort()
		}

		ctx.Set("uid", uid)
		ctx.Set("token", newToken)
    }
}