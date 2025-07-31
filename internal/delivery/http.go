package delivery

import (
	"github.com/Gym-Partner/api-common/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gitlab.com/gym-partner1/api/gym-partner-api/internal/controllers/user"
)

func RegisterRoutes(router *gin.Engine, deps *router.Dependencies) {
	_ = user.New(deps.Database, deps.Catalog)

	v1 := router.Group(viper.GetString("API_PREFIX"))
	{
		v1.GET("/ping", func(c *gin.Context) {})
	}
}
