package delivery

import (
	"github.com/Gym-Partner/api-common/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func RegisterRoutes(router *gin.Engine, deps *router.Dependencies) {

	v1 := router.Group(viper.GetString("API_PREFIX"))
	{
		v1.GET("/ping", func(c *gin.Context) {})
	}
}
